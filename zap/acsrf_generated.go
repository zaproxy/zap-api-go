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

import "strconv"

type Acsrf struct {
	c *Client
}

// Lists the names of all anti-CSRF tokens
func (a Acsrf) OptionTokensNames() (map[string]interface{}, error) {
	return a.c.Request("acsrf/view/optionTokensNames/", nil)
}

// Define if ZAP should detect CSRF tokens by searching for partial matches
func (a Acsrf) OptionPartialMatchingEnabled() (map[string]interface{}, error) {
	return a.c.Request("acsrf/view/optionPartialMatchingEnabled/", nil)
}

// Adds an anti-CSRF token with the given name, enabled by default
func (a Acsrf) AddOptionToken(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return a.c.Request("acsrf/action/addOptionToken/", m)
}

// Removes the anti-CSRF token with the given name
func (a Acsrf) RemoveOptionToken(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return a.c.Request("acsrf/action/removeOptionToken/", m)
}

// Define if ZAP should detect CSRF tokens by searching for partial matches.
func (a Acsrf) SetOptionPartialMatchingEnabled(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return a.c.Request("acsrf/action/setOptionPartialMatchingEnabled/", m)
}

// Generate a form for testing lack of anti-CSRF tokens - typically invoked via ZAP
func (a Acsrf) GenForm(hrefid string, actionurl string) ([]byte, error) {
	m := map[string]string{
		"hrefId":    hrefid,
		"actionUrl": actionurl,
	}
	return a.c.RequestOther("acsrf/other/genForm/", m)
}
