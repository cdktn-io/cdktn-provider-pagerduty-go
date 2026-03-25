// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package schedulev2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Schedulev2Config struct {
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
	// The name of the schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.1/docs/resources/schedulev2#name Schedulev2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The time zone of the schedule (IANA format, e.g. 'America/New_York').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.1/docs/resources/schedulev2#time_zone Schedulev2#time_zone}
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
	// A description of the schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.1/docs/resources/schedulev2#description Schedulev2#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// rotation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.1/docs/resources/schedulev2#rotation Schedulev2#rotation}
	Rotation interface{} `field:"optional" json:"rotation" yaml:"rotation"`
	// List of team IDs to associate with this schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.1/docs/resources/schedulev2#teams Schedulev2#teams}
	Teams *[]*string `field:"optional" json:"teams" yaml:"teams"`
}

