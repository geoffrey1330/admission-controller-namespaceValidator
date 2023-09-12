/*
Copyright 2023.

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
	"net/http"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	admissionv1 "namespaceselector.geoffrey.io/namespaceselector/api/v1"
)

// NamespaceSelectorReconciler reconciles a NamespaceSelector object
type NamespaceSelectorReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=admission.geoffrey.io,resources=namespaceselectors,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=admission.geoffrey.io,resources=namespaceselectors/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=admission.geoffrey.io,resources=namespaceselectors/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the NamespaceSelector object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.4/pkg/reconcile
func (r *NamespaceSelectorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *NamespaceSelectorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&admissionv1.NamespaceSelector{}).
		Complete(r)
}

// NamespaceSelectorValidator validates resources in specific namespaces based on label requirements
type NamespaceSelectorValidator struct {
	decoder *admission.Decoder
}

// Handle validates the resource admission request
func (v *NamespaceSelectorValidator) Handle(ctx context.Context, req admission.Request) admission.Response {
	// Fetch the NamespaceSelector object
	namespaceSelector := &admissionv1.NamespaceSelector{}
	if err := v.decoder.Decode(req, namespaceSelector); err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	// Fetch the resource being admitted

	resource := &unstructured.Unstructured{}
	if err := v.decoder.DecodeRaw(req.AdmissionRequest.Object, resource); err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	// Check if the resource's namespace matches the NamespaceSelector's namespace
	if namespaceSelector.Spec.Namespace != resource.GetNamespace() {
		return admission.Allowed(fmt.Sprintf("Resource is not in namespace %s", namespaceSelector.Spec.Namespace))
	}

	// Check if the resource has the required label
	labels := resource.GetLabels()
	requiredLabel := namespaceSelector.Spec.RequiredLabel
	if labels[requiredLabel] == "" {
		return admission.Denied(fmt.Sprintf("Resource does not have the required label: %s", requiredLabel))
	}

	return admission.Allowed("Resource validated successfully")
}

// +kubebuilder:webhook:path=/validate-admission-example-com-v1-namespaceselector,mutating=false,failurePolicy=fail,groups=admission.example.com,resources=namespaceselectors,verbs=create;update,versions=v1,name=vnamespaceselector.kb.io,admissionReviewVersions=v1

// NamespaceSelectorValidator implements admission.DecoderInjector.
// A decoder will be automatically injected.
func (v *NamespaceSelectorValidator) InjectDecoder(d *admission.Decoder) error {
	v.decoder = d
	return nil
}
