// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package schedulev2


type Schedulev2Rotation struct {
	// event block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.1/docs/resources/schedulev2#event Schedulev2#event}
	Event interface{} `field:"optional" json:"event" yaml:"event"`
}

