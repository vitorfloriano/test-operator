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

package v1

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	examplecomv1 "github/camilamacedo86/test-operator/api/v1"
)

// nolint:unused
// log is for logging in this package.
var wordpresslog = logf.Log.WithName("wordpress-resource")

// SetupWordpressWebhookWithManager registers the webhook for Wordpress in the manager.
func SetupWordpressWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&examplecomv1.Wordpress{}).
		WithValidator(&WordpressCustomValidator{}).
		WithDefaulter(&WordpressCustomDefaulter{}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-example-com-my-domain-v1-wordpress,mutating=true,failurePolicy=fail,sideEffects=None,groups=example.com.my.domain,resources=wordpresses,verbs=create;update,versions=v1,name=mwordpress-v1.kb.io,admissionReviewVersions=v1

// WordpressCustomDefaulter struct is responsible for setting default values on the custom resource of the
// Kind Wordpress when those are created or updated.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as it is used only for temporary operations and does not need to be deeply copied.
type WordpressCustomDefaulter struct {
	// TODO(user): Add more fields as needed for defaulting
}

var _ webhook.CustomDefaulter = &WordpressCustomDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind Wordpress.
func (d *WordpressCustomDefaulter) Default(_ context.Context, obj runtime.Object) error {
	wordpress, ok := obj.(*examplecomv1.Wordpress)

	if !ok {
		return fmt.Errorf("expected an Wordpress object but got %T", obj)
	}
	wordpresslog.Info("Defaulting for Wordpress", "name", wordpress.GetName())

	// TODO(user): fill in your defaulting logic.

	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// +kubebuilder:webhook:path=/validate-example-com-my-domain-v1-wordpress,mutating=false,failurePolicy=fail,sideEffects=None,groups=example.com.my.domain,resources=wordpresses,verbs=create;update,versions=v1,name=vwordpress-v1.kb.io,admissionReviewVersions=v1

// WordpressCustomValidator struct is responsible for validating the Wordpress resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type WordpressCustomValidator struct {
	// TODO(user): Add more fields as needed for validation
}

var _ webhook.CustomValidator = &WordpressCustomValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type Wordpress.
func (v *WordpressCustomValidator) ValidateCreate(_ context.Context, obj runtime.Object) (admission.Warnings, error) {
	wordpress, ok := obj.(*examplecomv1.Wordpress)
	if !ok {
		return nil, fmt.Errorf("expected a Wordpress object but got %T", obj)
	}
	wordpresslog.Info("Validation for Wordpress upon creation", "name", wordpress.GetName())

	// TODO(user): fill in your validation logic upon object creation.

	return nil, nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type Wordpress.
func (v *WordpressCustomValidator) ValidateUpdate(_ context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	wordpress, ok := newObj.(*examplecomv1.Wordpress)
	if !ok {
		return nil, fmt.Errorf("expected a Wordpress object for the newObj but got %T", newObj)
	}
	wordpresslog.Info("Validation for Wordpress upon update", "name", wordpress.GetName())

	// TODO(user): fill in your validation logic upon object update.

	return nil, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type Wordpress.
func (v *WordpressCustomValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	wordpress, ok := obj.(*examplecomv1.Wordpress)
	if !ok {
		return nil, fmt.Errorf("expected a Wordpress object but got %T", obj)
	}
	wordpresslog.Info("Validation for Wordpress upon deletion", "name", wordpress.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}
