// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package servicedependency

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceDependencyDependencyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ServiceDependencyDependencyList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceDependencyDependencyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceDependencyDependencyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceDependencyDependencyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceDependencyDependencyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceDependencyDependencyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceDependencyDependencyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

