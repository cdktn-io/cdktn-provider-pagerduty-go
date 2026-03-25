// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationglobalcachevariable

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EventOrchestrationGlobalCacheVariableConfig struct {
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
	// configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.1/docs/resources/event_orchestration_global_cache_variable#configuration EventOrchestrationGlobalCacheVariable#configuration}
	Configuration *EventOrchestrationGlobalCacheVariableConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.1/docs/resources/event_orchestration_global_cache_variable#event_orchestration EventOrchestrationGlobalCacheVariable#event_orchestration}.
	EventOrchestration *string `field:"required" json:"eventOrchestration" yaml:"eventOrchestration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.1/docs/resources/event_orchestration_global_cache_variable#name EventOrchestrationGlobalCacheVariable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.1/docs/resources/event_orchestration_global_cache_variable#condition EventOrchestrationGlobalCacheVariable#condition}
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.1/docs/resources/event_orchestration_global_cache_variable#disabled EventOrchestrationGlobalCacheVariable#disabled}.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}

