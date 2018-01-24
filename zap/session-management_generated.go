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

type SessionManagement struct {
	c *Client
}

func (s SessionManagement) GetSupportedSessionManagementMethods() (map[string]interface{}, error) {
	return s.c.Request("sessionManagement/view/getSupportedSessionManagementMethods/", nil)
}

func (s SessionManagement) GetSessionManagementMethodConfigParams(methodname string) (map[string]interface{}, error) {
	m := map[string]string{
		"methodName": methodname,
	}
	return s.c.Request("sessionManagement/view/getSessionManagementMethodConfigParams/", m)
}

func (s SessionManagement) GetSessionManagementMethod(contextid string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
	}
	return s.c.Request("sessionManagement/view/getSessionManagementMethod/", m)
}

func (s SessionManagement) SetSessionManagementMethod(contextid string, methodname string, methodconfigparams string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId":          contextid,
		"methodName":         methodname,
		"methodConfigParams": methodconfigparams,
	}
	return s.c.Request("sessionManagement/action/setSessionManagementMethod/", m)
}
