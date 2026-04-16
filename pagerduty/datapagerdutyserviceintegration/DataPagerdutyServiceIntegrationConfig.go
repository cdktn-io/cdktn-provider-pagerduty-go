// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datapagerdutyserviceintegration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataPagerdutyServiceIntegrationConfig struct {
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
	// examples "Amazon CloudWatch", "New Relic".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.2/docs/data-sources/service_integration#integration_summary DataPagerdutyServiceIntegration#integration_summary}
	IntegrationSummary *string `field:"required" json:"integrationSummary" yaml:"integrationSummary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.2/docs/data-sources/service_integration#service_name DataPagerdutyServiceIntegration#service_name}.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
}

