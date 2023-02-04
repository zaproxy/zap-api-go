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

type AlertFilter struct {
	c *Client
}

// Lists the alert filters of the context with the given ID.
//
// This component is optional and therefore the API will only work if it is installed
func (a AlertFilter) AlertFilterList(contextid string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
	}
	return a.c.Request("alertFilter/view/alertFilterList/", m)
}

// Lists the global alert filters.
//
// This component is optional and therefore the API will only work if it is installed
func (a AlertFilter) GlobalAlertFilterList() (map[string]interface{}, error) {
	return a.c.Request("alertFilter/view/globalAlertFilterList/", nil)
}

// Adds a new alert filter for the context with the given ID. 
//
// This component is optional and therefore the API will only work if it is installed
func (a AlertFilter) AddAlertFilter(contextid string, ruleid string, newlevel string, url string, urlisregex string, parameter string, enabled string, parameterisregex string, attack string, attackisregex string, evidence string, evidenceisregex string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
		"ruleId": ruleid,
		"newLevel": newlevel,
		"url": url,
		"urlIsRegex": urlisregex,
		"parameter": parameter,
		"enabled": enabled,
		"parameterIsRegex": parameterisregex,
		"attack": attack,
		"attackIsRegex": attackisregex,
		"evidence": evidence,
		"evidenceIsRegex": evidenceisregex,
	}
	return a.c.Request("alertFilter/action/addAlertFilter/", m)
}

// Removes an alert filter from the context with the given ID.
//
// This component is optional and therefore the API will only work if it is installed
func (a AlertFilter) RemoveAlertFilter(contextid string, ruleid string, newlevel string, url string, urlisregex string, parameter string, enabled string, parameterisregex string, attack string, attackisregex string, evidence string, evidenceisregex string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
		"ruleId": ruleid,
		"newLevel": newlevel,
		"url": url,
		"urlIsRegex": urlisregex,
		"parameter": parameter,
		"enabled": enabled,
		"parameterIsRegex": parameterisregex,
		"attack": attack,
		"attackIsRegex": attackisregex,
		"evidence": evidence,
		"evidenceIsRegex": evidenceisregex,
	}
	return a.c.Request("alertFilter/action/removeAlertFilter/", m)
}

// Adds a new global alert filter. 
//
// This component is optional and therefore the API will only work if it is installed
func (a AlertFilter) AddGlobalAlertFilter(ruleid string, newlevel string, url string, urlisregex string, parameter string, enabled string, parameterisregex string, attack string, attackisregex string, evidence string, evidenceisregex string) (map[string]interface{}, error) {
	m := map[string]string{
		"ruleId": ruleid,
		"newLevel": newlevel,
		"url": url,
		"urlIsRegex": urlisregex,
		"parameter": parameter,
		"enabled": enabled,
		"parameterIsRegex": parameterisregex,
		"attack": attack,
		"attackIsRegex": attackisregex,
		"evidence": evidence,
		"evidenceIsRegex": evidenceisregex,
	}
	return a.c.Request("alertFilter/action/addGlobalAlertFilter/", m)
}

// Removes a global alert filter.
//
// This component is optional and therefore the API will only work if it is installed
func (a AlertFilter) RemoveGlobalAlertFilter(ruleid string, newlevel string, url string, urlisregex string, parameter string, enabled string, parameterisregex string, attack string, attackisregex string, evidence string, evidenceisregex string) (map[string]interface{}, error) {
	m := map[string]string{
		"ruleId": ruleid,
		"newLevel": newlevel,
		"url": url,
		"urlIsRegex": urlisregex,
		"parameter": parameter,
		"enabled": enabled,
		"parameterIsRegex": parameterisregex,
		"attack": attack,
		"attackIsRegex": attackisregex,
		"evidence": evidence,
		"evidenceIsRegex": evidenceisregex,
	}
	return a.c.Request("alertFilter/action/removeGlobalAlertFilter/", m)
}

// Applies all currently enabled Global and Context alert filters.
//
// This component is optional and therefore the API will only work if it is installed
func (a AlertFilter) ApplyAll() (map[string]interface{}, error) {
	return a.c.Request("alertFilter/action/applyAll/", nil)
}

// Applies all currently enabled Context alert filters.
//
// This component is optional and therefore the API will only work if it is installed
func (a AlertFilter) ApplyContext() (map[string]interface{}, error) {
	return a.c.Request("alertFilter/action/applyContext/", nil)
}

// Applies all currently enabled Global alert filters.
//
// This component is optional and therefore the API will only work if it is installed
func (a AlertFilter) ApplyGlobal() (map[string]interface{}, error) {
	return a.c.Request("alertFilter/action/applyGlobal/", nil)
}

// Tests all currently enabled Global and Context alert filters.
//
// This component is optional and therefore the API will only work if it is installed
func (a AlertFilter) TestAll() (map[string]interface{}, error) {
	return a.c.Request("alertFilter/action/testAll/", nil)
}

// Tests all currently enabled Context alert filters.
//
// This component is optional and therefore the API will only work if it is installed
func (a AlertFilter) TestContext() (map[string]interface{}, error) {
	return a.c.Request("alertFilter/action/testContext/", nil)
}

// Tests all currently enabled Global alert filters.
//
// This component is optional and therefore the API will only work if it is installed
func (a AlertFilter) TestGlobal() (map[string]interface{}, error) {
	return a.c.Request("alertFilter/action/testGlobal/", nil)
}

