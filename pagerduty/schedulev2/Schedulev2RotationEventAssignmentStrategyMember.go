// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package schedulev2


type Schedulev2RotationEventAssignmentStrategyMember struct {
	// The member type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.1/docs/resources/schedulev2#type Schedulev2#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The obfuscated user ID. Required when type is 'user_member'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.1/docs/resources/schedulev2#user_id Schedulev2#user_id}
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}

