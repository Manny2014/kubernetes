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
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"strconv"
	"time"

	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	objectPrinter "printer-controller/api/v1"
	printerv1 "printer-controller/api/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ObjectPrinterReconciler reconciles a ObjectPrinter object
type ObjectPrinterReconciler struct {
	client.Client
	Log logr.Logger
}

func ignoreNotFound(err error) error {
	if apierrs.IsNotFound(err) {
		return nil
	}
	return err
}

func (r *ObjectPrinterReconciler) printer(req ctrl.Request, log logr.Logger) {
	ctx := context.Background()
	var currentObject objectPrinter.ObjectPrinter

	// get object
	if err := r.Get(ctx, req.NamespacedName, &currentObject); err != nil {
		log.Error(err, "unable to fetch object from printer")
		return
	}

	log.Info("running printer", "object", req.NamespacedName.String())
	time.Sleep(1000 * time.Millisecond)
	log.Info("object message", "message", currentObject.Spec.Message)
	log.Info("done running printer", "object", req.NamespacedName.String())

	count := currentObject.Status.PrintCount + int64(1)
	currentObject.Status.PrintCount = count

	log.Info("current status print count", strconv.FormatInt(currentObject.Status.PrintCount, 10), "")

	if err := r.Status().Update(ctx, &currentObject); err != nil {
		log.Error(err, "unable to update object status count")
		return
	}

	log.Info("done updating", "object", req.NamespacedName.String())
}

func containsString(slice []string, s string) bool {

	for _, item := range slice {
		if item == s {
			return true
		}
	}

	return false
}

func removeString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}

// +kubebuilder:rbac:groups=printer.manny87.com,resources=objectprinters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=printer.manny87.com,resources=objectprinters/status,verbs=get;update;patch

func (r *ObjectPrinterReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("objectprinter", req.NamespacedName)
	var currentObject objectPrinter.ObjectPrinter

	log.V(1).Info("attempting to fetch object with name", req.Name, "")

	// get object
	if err := r.Get(ctx, req.NamespacedName, &currentObject); err != nil {
		//log.Error(err, "unable to fetch object")
		return ctrl.Result{}, ignoreNotFound(err)
	}

	// Add finalizer

	finalizer := "objectPrinter.finalizer.printer.manny87.com"

	if currentObject.ObjectMeta.DeletionTimestamp.IsZero() {
		if !containsString(currentObject.ObjectMeta.Finalizers, finalizer) {
			currentObject.ObjectMeta.Finalizers = append(currentObject.ObjectMeta.Finalizers, finalizer)
			if err := r.Update(ctx, &currentObject); err != nil {
				log.Error(err, "unable to update object finalizer")
				return ctrl.Result{}, nil
			}
		}
	} else {
		// Object is being delete
		// TODO: Delete external resources & remove finalizer...

		log.Info("deleting external resources")
		currentObject.ObjectMeta.Finalizers = removeString(currentObject.ObjectMeta.Finalizers, finalizer)
		if err := r.Update(ctx, &currentObject); err != nil {
			log.Error(err, "unable to delete finalizer")
			return ctrl.Result{}, nil // TODO: re-queue??
		}

		return ctrl.Result{}, nil
	}

	log.Info("Current status counter vs current spec", strconv.FormatInt(currentObject.Status.PrintCount, 10), strconv.FormatInt(*currentObject.Spec.PrintCount, 10))
	if currentObject.Status.PrintCount < *currentObject.Spec.PrintCount {
		log.Info("running printer")
		go r.printer(req, log)
	}

	return ctrl.Result{}, nil
}

func (r *ObjectPrinterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&printerv1.ObjectPrinter{}).
		Complete(r)
}
