/*
Copyright 2019 The Knative Authors.

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
	"encoding/json"
	"fmt"
	sourcesv1 "sample-source/api/v1"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// MannySourceReconciler reconciles a MannySource object
type MannySourceReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=sources.manny.dev,resources=mannysources,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sources.manny.dev,resources=mannysources/status,verbs=get;update;patch

func (r *MannySourceReconciler) Reconcile(request ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("mannysource", request.NamespacedName)

	instance := &sourcesv1.MannySource{}
	err := r.Get(ctx, request.NamespacedName, instance)

	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// your logic here
	original := instance.DeepCopy()

	reconcileErr := r.reconcile(ctx, instance)

	// Update object Status if necessary. This happens even if the reconcile
	// returned an error.
	if !equality.Semantic.DeepEqual(original.Status, instance.Status) {
		log.Info("Updating Status", "request", request.NamespacedName)
		// An error may occur here if the object was updated since the last Get.
		// Return the error so the request can be retried later.
		// This call uses the /status subresource to ensure that the object's spec
		// is never updated by the controller.
		if err := r.Status().Update(ctx, instance); err != nil {
			return reconcile.Result{}, err
		}
	}

	return ctrl.Result{}, reconcileErr
}

func (r *MannySourceReconciler) reconcile(ctx context.Context, instance *sourcesv1.MannySource) error {
	// Resolve the Sink URI based on the sink reference.
	sinkURI, err := r.resolveSinkRef(ctx, instance.Spec.Sink)
	if err != nil {
		return fmt.Errorf("Failed to get sink URI: %v", err)
	}

	// Set the SinkURI field on the SampleSource Status.
	instance.Status.SinkURI = sinkURI

	//TODO(user): Add additional behavior.
	return err
}

type addressableType struct {
	Status struct {
		Address *struct {
			Hostname string
		}
	}
}

// TODO(user): A version of this function is also available in the
// github.com/knative/eventing-contrib/pkg/controller/sinks package.
func (r *MannySourceReconciler) resolveSinkRef(ctx context.Context, sinkRef *corev1.ObjectReference) (string, error) {
	// Make sure the reference is not nil.

	if sinkRef == nil {
		return "", fmt.Errorf("sink reference is nil")
	}

	//TODO(user): Add support for corev1.Service.

	// Get the referenced Sink as an Unstructured object.
	sink := &unstructured.Unstructured{}
	sink.SetGroupVersionKind(sinkRef.GroupVersionKind())

	fmt.Println("OBJECT:", sink)

	if err := r.Get(ctx, client.ObjectKey{Namespace: sinkRef.Namespace, Name: sinkRef.Name}, sink); err != nil {
		return "", fmt.Errorf("Failed to get sink object: %v", err)
	}

	// Marshal the Sink into an Addressable struct to more easily extract its
	// hostname.
	addressable := &addressableType{}
	raw, err := sink.MarshalJSON()

	if err != nil {
		return "", fmt.Errorf("Failed to marshal sink: %v", err)
	}

	if err := json.Unmarshal(raw, addressable); err != nil {
		return "", fmt.Errorf("Failed to marshal sink into Addressable: %v", err)
	}

	// Check that the Addressable fields are present.
	if addressable.Status.Address == nil {
		return "", fmt.Errorf("Failed to resolve sink URI: sink does not contain address")
	}
	if addressable.Status.Address.Hostname == "" {
		return "", fmt.Errorf("Failed to resolve sink URI: address hostname is empty")
	}
	// Translate the Hostname into a URI.
	return fmt.Sprintf("http://%s/", addressable.Status.Address.Hostname), nil
}

func (r *MannySourceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&sourcesv1.MannySource{}).
		Complete(r)
}
