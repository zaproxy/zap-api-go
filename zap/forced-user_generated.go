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

type ForcedUser struct {
	c *Client
}

// Returns 'true' if 'forced user' mode is enabled, 'false' otherwise
func (f ForcedUser) IsForcedUserModeEnabled() (map[string]interface{}, error) {
	return f.c.Request("forcedUser/view/isForcedUserModeEnabled/", nil)
}

// Gets the user (ID) set as 'forced user' for the given context (ID)
func (f ForcedUser) GetForcedUser(contextid string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
	}
	return f.c.Request("forcedUser/view/getForcedUser/", m)
}

// Sets the user (ID) that should be used in 'forced user' mode for the given context (ID)
func (f ForcedUser) SetForcedUser(contextid string, userid string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
		"userId": userid,
	}
	return f.c.Request("forcedUser/action/setForcedUser/", m)
}

// Sets if 'forced user' mode should be enabled or not
func (f ForcedUser) SetForcedUserModeEnabled(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"boolean": strconv.FormatBool(boolean),
	}
	return f.c.Request("forcedUser/action/setForcedUserModeEnabled/", m)
}

