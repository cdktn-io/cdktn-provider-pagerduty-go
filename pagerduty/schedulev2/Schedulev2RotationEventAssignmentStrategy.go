// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package schedulev2


type Schedulev2RotationEventAssignmentStrategy struct {
	// The assignment strategy type. Use 'rotating_member_assignment_strategy' for user-based rotation, or 'every_member_assignment_strategy' for all-hands coverage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.2/docs/resources/schedulev2#type Schedulev2#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// member block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.2/docs/resources/schedulev2#member Schedulev2#member}
	Member interface{} `field:"optional" json:"member" yaml:"member"`
	// Number of shifts per member per recurrence cycle. Defaults to 1 for 'rotating_member_assignment_strategy' when not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.2/docs/resources/schedulev2#shifts_per_member Schedulev2#shifts_per_member}
	ShiftsPerMember *float64 `field:"optional" json:"shiftsPerMember" yaml:"shiftsPerMember"`
}

