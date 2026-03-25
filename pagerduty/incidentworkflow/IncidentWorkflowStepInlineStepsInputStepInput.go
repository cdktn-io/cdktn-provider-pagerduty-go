// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package incidentworkflow


type IncidentWorkflowStepInlineStepsInputStepInput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.1/docs/resources/incident_workflow#name IncidentWorkflow#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.1/docs/resources/incident_workflow#value IncidentWorkflow#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

