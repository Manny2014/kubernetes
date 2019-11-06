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
	corev1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	configMapPrinter "printer-controller/api/v1"
	printerv1 "printer-controller/api/v1"
	"reflect"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// ConfigMapPrinterReconciler reconciles a ConfigMapPrinter object
type ConfigMapPrinterReconciler struct {
	client.Client
	Log logr.Logger
}

func (r *ConfigMapPrinterReconciler) cmHandler(ctx context.Context, currentObject configMapPrinter.ConfigMapPrinter, req ctrl.Request, log logr.Logger) {

	var cm corev1.ConfigMap
	exists := true

	log.Info("checking if cm exists for", "name", req.NamespacedName.String())

	if err := r.Get(ctx, req.NamespacedName, &cm); err != nil {
		if apierrs.IsNotFound(err) {
			exists = false
		}
	}


	if !exists {
		log.Info("creating config map because it does not exists...", "data", currentObject.Spec.Information)

		cm := &corev1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name:      currentObject.Name,
				Namespace: currentObject.Namespace,
				OwnerReferences: []metav1.OwnerReference{
					*metav1.NewControllerRef(&currentObject, configMapPrinter.SchemeBuilder.GroupVersion.WithKind("ConfigMapPrinter")),
				},
			},
			Data: currentObject.Spec.Information,
		}

		scheme, _ := configMapPrinter.SchemeBuilder.Build()

		_ = controllerutil.SetControllerReference(&currentObject, cm, scheme)

		if err := r.Create(ctx, cm); err != nil {
			if !apierrs.IsAlreadyExists(err) {
				log.Error(err, "unable to create config map")
				return
			}
		}

		return
	}

	// Object was found
	log.Info("configmap already exists", "name", req.NamespacedName)

	if !reflect.DeepEqual(cm.Data, currentObject.Spec.Information) {
		if currentObject.Spec.Managed{
			// Owner is the source of truth
			// Set back to desired state.
			log.Info("Updating", "data", cm.Data)
			cm.Data = currentObject.Spec.Information
			if err := r.Update(ctx, &cm); err != nil {
				if !apierrs.IsAlreadyExists(err) {
					log.Error(err, "unable to update config map")
					return
				}
			}
		}else{
			currentObject.Status.Revision = currentObject.Status.Revision

			if err := r.Status().Update(ctx, &currentObject); err != nil {
				if !apierrs.IsAlreadyExists(err) {
					log.Error(err, "unable to update object revision")
					return
				}
			}
		}
	}

	return
}

func (r *ConfigMapPrinterReconciler) finalizer(currentObject configMapPrinter.ConfigMapPrinter, ctx context.Context, log logr.Logger) (ctrl.Result, error) {
	finalizer := "configmap.finalizer.printer.manny87.com"

	if currentObject.ObjectMeta.DeletionTimestamp.IsZero() {
		if !containsString(currentObject.ObjectMeta.Finalizers, finalizer) {
			currentObject.ObjectMeta.Finalizers = append(currentObject.ObjectMeta.Finalizers, finalizer)
			if err := r.Update(ctx, &currentObject); err != nil {
				log.Error(err, "unable to update object finalizer")
				return ctrl.Result{}, err // TODO: re-queue??
			}
		}
	} else {
		// Object is being delete
		// TODO: Delete external resources & remove finalizer...

		log.Info("deleting external resources")
		currentObject.ObjectMeta.Finalizers = removeString(currentObject.ObjectMeta.Finalizers, finalizer)
		if err := r.Update(ctx, &currentObject); err != nil {
			log.Error(err, "unable to delete finalizer")
			return ctrl.Result{}, err // TODO: re-queue??
		}

	}

	return ctrl.Result{}, nil
}

// +kubebuilder:rbac:groups=printer.manny87.com,resources=configmapprinters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=printer.manny87.com,resources=configmapprinters/status,verbs=get;update;patch
// +kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch;create;update;patch;delete

func (r *ConfigMapPrinterReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("configmapprinter", req.NamespacedName)

	var currentObject configMapPrinter.ConfigMapPrinter

	if err := r.Get(ctx, req.NamespacedName, &currentObject); err != nil {
		log.V(2).Info("unable to fetch object", "error", err.Error())
		return ctrl.Result{}, ignoreNotFound(err)
	}

	// your logic here
	log.Info("received object", "objectKind", currentObject)

	// Setup Finalizer - External object deletes
	_, _ = r.finalizer(currentObject, ctx, log)

	r.cmHandler(ctx, currentObject, req, log)

	return ctrl.Result{}, nil
}

func (r *ConfigMapPrinterReconciler) SetupWithManager(mgr ctrl.Manager) error {

	mgr.GetFieldIndexer().IndexField(&corev1.ConfigMap{}, ".metadata.controller", func(rawObj runtime.Object) []string {

		cm := rawObj.(*corev1.ConfigMap)
		owner := metav1.GetControllerOf(cm)

		log := r.Log.WithValues("cm-index-field", cm.Name)

		if owner == nil {
			return nil
		}

		if owner.APIVersion != configMapPrinter.GroupVersion.String() || owner.Kind != "ConfigMapPrinter" {
			return nil
		}

		log.Info("cm belongs to ConfigMapPrinter", "name", cm.Name)
		return []string{owner.Name}
	})

	return ctrl.NewControllerManagedBy(mgr).
		For(&printerv1.ConfigMapPrinter{}).
		Owns(&corev1.ConfigMap{}).
		Complete(r)
}
