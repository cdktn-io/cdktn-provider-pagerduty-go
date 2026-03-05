// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicecustomfieldvalue

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ServiceCustomFieldValueConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The custom field values to set for the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.31.3/docs/resources/service_custom_field_value#custom_fields ServiceCustomFieldValue#custom_fields}
	CustomFields interface{} `field:"required" json:"customFields" yaml:"customFields"`
	// The ID of the service to set custom field values for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.31.3/docs/resources/service_custom_field_value#service_id ServiceCustomFieldValue#service_id}
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
}

