/*
Copyright 2025.

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

package controller

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	crdv1 "devops.tools/controller/api/v1"
)

// ChowkiReconciler reconciles a Chowki object
type ChowkiReconciler struct {
	client.Client
	Cache  cache.Cache
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=crd.devops.tools,resources=chowkis,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=crd.devops.tools,resources=chowkis/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=crd.devops.tools,resources=chowkis/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Chowki object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.20.2/pkg/reconcile
func (r *ChowkiReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.Info("controller triggered")
	logger.Info("Reconciling Chowki resource", "Name", req.Name, "Namespace", req.Namespace)

	// TODO(user): your logic here
	var chowkiList crdv1.ChowkiList

	if err := r.List(ctx, &chowkiList); err != nil {
		logger.Error(err, "unable to fetch watcher list")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if len(chowkiList.Items) == 0 {
		logger.V(1).Info("no watchers configured")
		return ctrl.Result{}, nil
	} else {
		// List events in the namespace of the custom resource
		eventList := &corev1.EventList{}

		// Define the list options with the namespace filter
		listOptions := &client.ListOptions{
			Namespace: req.Namespace,
			Raw: &v1.ListOptions{
				FieldSelector: fmt.Sprintf("involvedObject.name=%s", req.Name), // Filter by resource name
			},
		}

		if err := r.List(ctx, eventList, listOptions); err != nil {
			logger.Error(err, "Failed to list events in namespace", "Namespace", req.Namespace)
			return ctrl.Result{}, err
		}

		// Print events
		for _, event := range eventList.Items {
			fmt.Printf("Event: %s | Reason: %s | Message: %s | Type: %s\n",
				event.InvolvedObject.Name, event.Reason, event.Message, event.Type)
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ChowkiReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&crdv1.Chowki{}).
		Named("chowki").
		Complete(r)
}
