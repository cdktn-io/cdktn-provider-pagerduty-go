// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package schedulev2


type Schedulev2RotationEvent struct {
	// When this event configuration starts producing shifts (ISO-8601 UTC).
	//
	// Must be a future time; the API will adjust past times to the current time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.3/docs/resources/schedulev2#effective_since Schedulev2#effective_since}
	EffectiveSince *string `field:"required" json:"effectiveSince" yaml:"effectiveSince"`
	// The shift end time with timezone offset (ISO-8601 format, e.g. '2024-01-01T17:00:00-05:00').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.3/docs/resources/schedulev2#end_time Schedulev2#end_time}
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// The name of the event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.3/docs/resources/schedulev2#name Schedulev2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of RRULE strings defining the recurrence pattern (RFC 5545, e.g. 'RRULE:FREQ=WEEKLY;BYDAY=MO,TU,WE,TH,FR').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.3/docs/resources/schedulev2#recurrence Schedulev2#recurrence}
	Recurrence *[]*string `field:"required" json:"recurrence" yaml:"recurrence"`
	// The shift start time with timezone offset (ISO-8601 format, e.g. '2024-01-01T09:00:00-05:00').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.3/docs/resources/schedulev2#start_time Schedulev2#start_time}
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// assignment_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.3/docs/resources/schedulev2#assignment_strategy Schedulev2#assignment_strategy}
	AssignmentStrategy interface{} `field:"optional" json:"assignmentStrategy" yaml:"assignmentStrategy"`
	// When this event configuration stops producing shifts (ISO-8601 UTC). Null or omitted means indefinite.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.3/docs/resources/schedulev2#effective_until Schedulev2#effective_until}
	EffectiveUntil *string `field:"optional" json:"effectiveUntil" yaml:"effectiveUntil"`
}

