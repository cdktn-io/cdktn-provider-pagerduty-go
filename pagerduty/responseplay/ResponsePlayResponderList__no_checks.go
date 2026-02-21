// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package responseplay

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ResponsePlayResponderList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_ResponsePlayResponderList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_ResponsePlayResponderList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ResponsePlayResponderList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ResponsePlayResponderList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ResponsePlayResponderList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ResponsePlayResponderList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewResponsePlayResponderListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

