// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package jiracloudaccountmappingrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type JiraCloudAccountMappingRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.31.0/docs/resources/jira_cloud_account_mapping_rule#account_mapping JiraCloudAccountMappingRule#account_mapping}.
	AccountMapping *string `field:"required" json:"accountMapping" yaml:"accountMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.31.0/docs/resources/jira_cloud_account_mapping_rule#name JiraCloudAccountMappingRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.31.0/docs/resources/jira_cloud_account_mapping_rule#config JiraCloudAccountMappingRule#config}
	Config *JiraCloudAccountMappingRuleConfigA `field:"optional" json:"config" yaml:"config"`
	// Indicates if the rule is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.31.0/docs/resources/jira_cloud_account_mapping_rule#enabled JiraCloudAccountMappingRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

