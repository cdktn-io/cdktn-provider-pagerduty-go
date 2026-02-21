// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package incidentworkflow

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IncidentWorkflowStepInputList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IncidentWorkflowStepInputList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IncidentWorkflowStepInputList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IncidentWorkflowStepInputList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IncidentWorkflowStepInputList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IncidentWorkflowStepInputList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IncidentWorkflowStepInputList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIncidentWorkflowStepInputListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

