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

type Replacer struct {
	c *Client
}

// Returns full details of all of the rules
//
// This component is optional and therefore the API will only work if it is installed
func (r Replacer) Rules() (map[string]interface{}, error) {
	return r.c.Request("replacer/view/rules/", nil)
}

// Adds a replacer rule. For the parameters: desc is a user friendly description, enabled is true or false, matchType is one of [REQ_HEADER, REQ_HEADER_STR, REQ_BODY_STR, RESP_HEADER, RESP_HEADER_STR, RESP_BODY_STR], matchRegex should be true if the matchString should be treated as a regex otherwise false, matchString is the string that will be matched against, replacement is the replacement string, initiators may be blank (for all initiators) or a comma separated list of integers as defined in <a href="https://github.com/zaproxy/zaproxy/blob/main/zap/src/main/java/org/parosproxy/paros/network/HttpSender.java">HttpSender</a>
//
// This component is optional and therefore the API will only work if it is installed
func (r Replacer) AddRule(description string, enabled string, matchtype string, matchregex string, matchstring string, replacement string, initiators string, url string) (map[string]interface{}, error) {
	m := map[string]string{
		"description": description,
		"enabled":     enabled,
		"matchType":   matchtype,
		"matchRegex":  matchregex,
		"matchString": matchstring,
		"replacement": replacement,
		"initiators":  initiators,
		"url":         url,
	}
	return r.c.Request("replacer/action/addRule/", m)
}

// Removes the rule with the given description
//
// This component is optional and therefore the API will only work if it is installed
func (r Replacer) RemoveRule(description string) (map[string]interface{}, error) {
	m := map[string]string{
		"description": description,
	}
	return r.c.Request("replacer/action/removeRule/", m)
}

// Enables or disables the rule with the given description based on the bool parameter
//
// This component is optional and therefore the API will only work if it is installed
func (r Replacer) SetEnabled(description string, bool string) (map[string]interface{}, error) {
	m := map[string]string{
		"description": description,
		"bool":        bool,
	}
	return r.c.Request("replacer/action/setEnabled/", m)
}
