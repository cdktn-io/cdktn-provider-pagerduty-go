// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package slackconnection

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SlackConnectionConfigAList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SlackConnectionConfigAList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SlackConnectionConfigAList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SlackConnectionConfigAList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SlackConnectionConfigAList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SlackConnectionConfigAList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SlackConnectionConfigAList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSlackConnectionConfigAListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

