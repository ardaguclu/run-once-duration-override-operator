package handlers

import (
	appsv1 "github.com/openshift/run-once-duration-override-operator/pkg/apis/apps/v1"
	"github.com/openshift/run-once-duration-override-operator/pkg/runoncedurationoverride/internal/condition"
	controllerreconciler "sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func NewValidationHandler(o *Options) *validationHandler {
	return &validationHandler{}
}

type validationHandler struct {
}

func (c *validationHandler) Handle(context *ReconcileRequestContext, original *appsv1.RunOnceDurationOverride) (current *appsv1.RunOnceDurationOverride, result controllerreconciler.Result, handleErr error) {
	current = original

	validationErr := original.Spec.Validate()
	if validationErr != nil {
		handleErr = condition.NewInstallReadinessError(appsv1.InvalidParameters, validationErr)
	}

	return
}
