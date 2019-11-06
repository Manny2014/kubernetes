/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	projectv1 "k8s-project/api/v1"
	sources "k8s-project/pkg/v1"
	corev1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"reflect"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"time"
)

// ConfigSourceReconciler reconciles a ConfigSource object
type ConfigSourceReconciler struct {
	client.Client
	Log logr.Logger
}

func (r *ConfigSourceReconciler) objectRetryOrNot(co projectv1.ConfigSource, ctx context.Context, log logr.Logger) (ctrl.Result, error) {
	defaultResult := ctrl.Result{RequeueAfter: 5000 * time.Millisecond}

	// Bump retry count
	co.Status.FailureCount = co.Status.FailureCount + 1

	if co.Status.FailureCount >= projectv1.MaxFailures {
		// No reque if maxed reached
		defaultResult = ctrl.Result{}

		// TODO: Mark parent object as DELETE for clean up....
	}

	if err := r.Update(ctx, &co); err != nil {
		log.Error(err, "unable update object")
		// We don't know why we can't get it...no retry...
		return ctrl.Result{}, err
	}

	return defaultResult, nil
}

// +kubebuilder:rbac:groups="",resources=secrets,verbs=get

func (r *ConfigSourceReconciler) cmReconcile(co projectv1.ConfigSource, req ctrl.Request, ctx context.Context) (ctrl.Result, error) {
	var s sources.Source
	log := r.Log.WithValues("cm-reconcile", fmt.Sprintf("%s/%s", co.Namespace, co.Name))

	log.Info("creating source", "name", req.NamespacedName)

	s, err := sources.NewSource(co.Spec.SourceType, co.Spec.SourceConfig)

	if err != nil {
		// Bump retry count
		return r.objectRetryOrNot(co, ctx, log)
	}

	log.Info("source construction was sucessful for", "name", req.NamespacedName)

	// Retrieve data from ConfigSource
	log.Info("retrieving data from source", "name", req.NamespacedName)
	data := make(map[string]string)
	data, err = s.GetData()

	if err != nil {
		return r.objectRetryOrNot(co, ctx, log)
	}

	log.Info("succesfully retrieved data from source", "name", req.NamespacedName)

	// TODO: Split to its own function

	// Create confimap template
	cm := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      co.Name,
			Namespace: co.Namespace,
		},
		Data: data,
	}

	scheme, _ := projectv1.SchemeBuilder.Build()

	_ = controllerutil.SetControllerReference(&co, cm, scheme)

	cmExists := true

	if err := r.Get(ctx, req.NamespacedName, cm); err != nil {
		if apierrs.IsNotFound(err) {
			cmExists = false
		}
	}

	if !cmExists {
		log.Info("creating configmap", "name", req.NamespacedName, "values", data)
		if err := r.Create(ctx, cm); err != nil {
			if !apierrs.IsAlreadyExists(err) {
				log.Error(err, "unable to create config map")
				return ctrl.Result{}, err
			}
		}

		return ctrl.Result{}, err
	} else {
		// CM Exits
		log.Info("configmap already exists", "name", req.NamespacedName)
		if !reflect.DeepEqual(cm.Data, data) {
			log.Info("Updating", "name", cm.Name)
			cm.Data = data
			if err := r.Update(ctx, cm); err != nil {
				log.Error(err, "unable to update config map")
				return ctrl.Result{}, err
			}
		}
	}

	return ctrl.Result{}, nil
}

// +kubebuilder:rbac:groups=project.manny87.com,resources=configsources,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=project.manny87.com,resources=configsources/status,verbs=get;update;patch
// +kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch;create;update;patch;delete

func (r *ConfigSourceReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("configsource", req.NamespacedName)

	var co projectv1.ConfigSource

	// STEP # 1:  Get the current object

	if err := r.Get(ctx, req.NamespacedName, &co); err != nil {
		log.V(2).Info("unable to retrieve object", "name", req.NamespacedName)
		return ctrl.Result{}, ignoreNotFound(err)
	}

	// STEP # 2: Reconcile ConfigMap
	return r.cmReconcile(co, req, ctx)
}

func (r *ConfigSourceReconciler) SetupWithManager(mgr ctrl.Manager) error {

	mgr.GetFieldIndexer().IndexField(&corev1.ConfigMap{}, ".metadata.controller", func(rawObj runtime.Object) []string {

		cm := rawObj.(*corev1.ConfigMap)
		owner := metav1.GetControllerOf(cm)

		log := r.Log.WithValues("cm-indexer", cm.Name)

		if owner == nil {
			return nil
		}

		if owner.APIVersion != projectv1.GroupVersion.String() || owner.Kind != "ConfigSource" {
			return nil
		}

		log.Info("cm belongs to ConfigSource", "name", cm.Name)
		return []string{owner.Name}
	})

	return ctrl.NewControllerManagedBy(mgr).
		For(&projectv1.ConfigSource{}).
		Owns(&corev1.ConfigMap{}).
		Complete(r)
}
