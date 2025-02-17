// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validate

import (
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
)

var ValidateWebApplicationFirewallPolicyRuleGroupName = validation.StringInSlice([]string{
	"BadBots",
	"crs_20_protocol_violations",
	"crs_21_protocol_anomalies",
	"crs_23_request_limits",
	"crs_30_http_policy",
	"crs_35_bad_robots",
	"crs_40_generic_attacks",
	"crs_41_sql_injection_attacks",
	"crs_41_xss_attacks",
	"crs_42_tight_security",
	"crs_45_trojans",
	"General",
	"GoodBots",
	"Known-CVEs",
	"REQUEST-911-METHOD-ENFORCEMENT",
	"REQUEST-913-SCANNER-DETECTION",
	"REQUEST-920-PROTOCOL-ENFORCEMENT",
	"REQUEST-921-PROTOCOL-ATTACK",
	"REQUEST-930-APPLICATION-ATTACK-LFI",
	"REQUEST-931-APPLICATION-ATTACK-RFI",
	"REQUEST-932-APPLICATION-ATTACK-RCE",
	"REQUEST-933-APPLICATION-ATTACK-PHP",
	"REQUEST-941-APPLICATION-ATTACK-XSS",
	"REQUEST-942-APPLICATION-ATTACK-SQLI",
	"REQUEST-943-APPLICATION-ATTACK-SESSION-FIXATION",
	"REQUEST-944-APPLICATION-ATTACK-JAVA",
	"UnknownBots",
	"METHOD-ENFORCEMENT",
	"PROTOCOL-ENFORCEMENT",
	"PROTOCOL-ATTACK",
	"APPLICATION-ATTACK-LFI",
	"APPLICATION-ATTACK-RFI",
	"APPLICATION-ATTACK-RCE",
	"APPLICATION-ATTACK-PHP",
	"APPLICATION-ATTACK-NodeJS",
	"APPLICATION-ATTACK-XSS",
	"APPLICATION-ATTACK-SQLI",
	"APPLICATION-ATTACK-SESSION-FIXATION",
	"APPLICATION-ATTACK-SESSION-JAVA",
	"MS-ThreatIntel-WebShells",
	"MS-ThreatIntel-AppSec",
	"MS-ThreatIntel-SQLI",
	"MS-ThreatIntel-CVEs",
}, false)

var ValidateWebApplicationFirewallPolicyRuleSetVersion = validation.StringInSlice([]string{
	"0.1",
	"1.0",
	"2.1",
	"2.2.9",
	"3.0",
	"3.1",
	"3.2",
}, false)

var ValidateWebApplicationFirewallPolicyRuleSetType = validation.StringInSlice([]string{
	"OWASP",
	"Microsoft_BotManagerRuleSet",
	"Microsoft_DefaultRuleSet",
}, false)

var ValidateWebApplicationFirewallPolicyExclusionRuleSetVersion = validation.StringInSlice([]string{
	"2.1",
	"3.2",
}, false)

var ValidateWebApplicationFirewallPolicyExclusionRuleSetType = validation.StringInSlice([]string{
	"OWASP",
	"Microsoft_DefaultRuleSet",
}, false)
