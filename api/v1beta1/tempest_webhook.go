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

//
// Generated by:
//
// operator-sdk create webhook --group test --version v1beta1 --kind Tempest --programmatic-validation --defaulting
//

package v1beta1

import (
        "errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
        "sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// TempestDefaults -
type TempestDefaults struct {
	ContainerImageURL string
}

var tempestDefaults TempestDefaults

// log is for logging in this package.
var tempestlog = logf.Log.WithName("tempest-resource")

func (r *Tempest) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-test-openstack-org-v1beta1-tempest,mutating=true,failurePolicy=fail,sideEffects=None,groups=test.openstack.org,resources=tempests,verbs=create;update,versions=v1beta1,name=mtempest.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Tempest{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Tempest) Default() {
	tempestlog.Info("default", "name", r.Name)

        r.Spec.Default()
}

// Default - set defaults for this Tempest spec.
func (spec *TempestSpec) Default() {
        if spec.ContainerImage == "" {
		spec.ContainerImage = tempestDefaults.ContainerImageURL
	}

	if spec.TempestconfRun == (TempestconfRunSpec{}) {
		spec.TempestconfRun.Create = true
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-test-openstack-org-v1beta1-tempest,mutating=false,failurePolicy=fail,sideEffects=None,groups=test.openstack.org,resources=tempests,verbs=create;update,versions=v1beta1,name=vtempest.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Tempest{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Tempest) ValidateCreate() (admission.Warnings, error) {
	tempestlog.Info("validate create", "name", r.Name)

        if len(r.Spec.Workflow) > 0 && r.Spec.Debug {
            return nil, errors.New("Workflow variable must be empty to run debug mode!")
        }
	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Tempest) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	tempestlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Tempest) ValidateDelete() (admission.Warnings, error) {
	tempestlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil, nil
}
