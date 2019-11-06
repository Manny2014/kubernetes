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
	"github.com/go-logr/logr"
	projectv1 "k8s-project/api/v1"
	corev1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// ProjectReconciler reconciles a Project object
type ProjectReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=project.manny87.com,resources=projects,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=project.manny87.com,resources=projects/status,verbs=get;update;patch
// +kubebuilder:rbac:groups="",resources=namespaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=project.manny87.com,resources=configsources,verbs=get;list;watch;create;update;patch;delete

func (r *ProjectReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("project", req.NamespacedName)

	var co projectv1.Project
	updateStatus := false

	log.Info("received object", "name", req.NamespacedName)
	// your logic here
	if err := r.Get(ctx, req.NamespacedName, &co); err != nil {
		log.V(2).Info("unable to retrieve object", "name", req.NamespacedName)
		return ctrl.Result{}, ignoreNotFound(err)
	}

	scheme, _ := projectv1.SchemeBuilder.Build()

	// Create ConfigSource
	if !co.Status.NSCreated {

		// Create Namespace
		nsExists := true

		ns := &corev1.Namespace{}

		ns.Name = co.Name
		_ = controllerutil.SetControllerReference(&co, ns, scheme)
		log.Info("checking if namespace exists", "name", req.NamespacedName)
		if err := r.Get(ctx, req.NamespacedName, ns); err != nil {
			if apierrs.IsNotFound(err) {
				log.Info("namespace does not exists", "name", req.NamespacedName)
				nsExists = false
			} else {
				return ctrl.Result{}, err
			}
		}

		if !nsExists {
			log.Info("creating namespace", "name", req.NamespacedName)
			if err := r.Create(ctx, ns); err != nil {
				log.V(2).Info("unable to create ns object", "name", req.NamespacedName)
				return ctrl.Result{}, err
			}

			co.Status.NSCreated = true
			updateStatus = true
		}
	}

	if !co.Status.CSCreated {

		cs := &projectv1.ConfigSource{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: co.Name,
				Name:      co.Name,
			},
			Spec: *co.Spec.ConfigSource,
		}

		_ = controllerutil.SetControllerReference(&co, cs, scheme)

		csExists := true

		if err := r.Get(ctx, req.NamespacedName, cs); err != nil {
			if apierrs.IsNotFound(err) {
				csExists = false
			}
		}

		if !csExists {
			if err := r.Create(ctx, cs); err != nil {
				log.V(2).Info("unable to create cs object", "name", req.NamespacedName)
				return ctrl.Result{}, err
			}

			co.Status.CSCreated = true
			updateStatus = true
		}
	}

	if updateStatus {
		if err := r.Status().Update(ctx, &co); err != nil {
			log.V(2).Info("unable to update status", "name", req.NamespacedName)
			return ctrl.Result{}, ignoreNotFound(err)
		}
	}

	return ctrl.Result{}, nil
}

func (r *ProjectReconciler) SetupWithManager(mgr ctrl.Manager) error {

	mgr.GetFieldIndexer().IndexField(&projectv1.ConfigSource{}, ".metadata.controller", func(rawObj runtime.Object) []string {

		cm := rawObj.(*projectv1.ConfigSource)
		owner := metav1.GetControllerOf(cm)

		log := r.Log.WithValues("cs-index-field", cm.Name)

		if owner == nil {
			return nil
		}

		if owner.APIVersion != projectv1.GroupVersion.String() || owner.Kind != "Project" {
			return nil
		}

		log.Info("cs belongs to ConfigMapPrinter", "name", cm.Name)
		return []string{owner.Name}
	})

	mgr.GetFieldIndexer().IndexField(&corev1.Namespace{}, ".metadata.controller", func(rawObj runtime.Object) []string {

		cm := rawObj.(*corev1.Namespace)
		owner := metav1.GetControllerOf(cm)

		log := r.Log.WithValues("ns-index-field", cm.Name)

		if owner == nil {
			return nil
		}

		if owner.APIVersion != projectv1.GroupVersion.String() || owner.Kind != "Project" {
			return nil
		}

		log.Info("ns belongs to Project", "name", cm.Name)
		return []string{owner.Name}
	})

	return ctrl.NewControllerManagedBy(mgr).
		For(&projectv1.Project{}).
		Owns(&projectv1.ConfigSource{}).
		Owns(&corev1.Namespace{}).
		Complete(r)
}
