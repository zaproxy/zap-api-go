// Zed Attack Proxy (ZAP) and its related class files.
//
// ZAP is an HTTP/HTTPS proxy for assessing web application security.
//
// Copyright 2017 the ZAP development team
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// *** This file was automatically generated. ***
//

package zap

type RuleConfig struct {
	c *Client
}

// Show the specified rule configuration
func (r RuleConfig) RuleConfigValue(key string) (map[string]interface{}, error) {
	m := map[string]string{
		"key": key,
	}
	return r.c.Request("ruleConfig/view/ruleConfigValue/", m)
}

// Show all of the rule configurations
func (r RuleConfig) AllRuleConfigs() (map[string]interface{}, error) {
	return r.c.Request("ruleConfig/view/allRuleConfigs/", nil)
}

// Reset the specified rule configuration, which must already exist
func (r RuleConfig) ResetRuleConfigValue(key string) (map[string]interface{}, error) {
	m := map[string]string{
		"key": key,
	}
	return r.c.Request("ruleConfig/action/resetRuleConfigValue/", m)
}

// Reset all of the rule configurations
func (r RuleConfig) ResetAllRuleConfigValues() (map[string]interface{}, error) {
	return r.c.Request("ruleConfig/action/resetAllRuleConfigValues/", nil)
}

// Set the specified rule configuration, which must already exist
func (r RuleConfig) SetRuleConfigValue(key string, value string) (map[string]interface{}, error) {
	m := map[string]string{
		"key": key,
		"value": value,
	}
	return r.c.Request("ruleConfig/action/setRuleConfigValue/", m)
}

