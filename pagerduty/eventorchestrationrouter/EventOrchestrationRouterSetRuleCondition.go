// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventorchestrationrouter


type EventOrchestrationRouterSetRuleCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/pagerduty/pagerduty/3.32.1/docs/resources/event_orchestration_router#expression EventOrchestrationRouter#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

