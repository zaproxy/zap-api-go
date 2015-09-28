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

type Users struct {
	c *Client
}

func (u Users) UsersList(contextid string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
	}
	return u.c.Request("users/view/usersList/", m)
}

func (u Users) GetUserById(contextid string, userid string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
		"userId":    userid,
	}
	return u.c.Request("users/view/getUserById/", m)
}

func (u Users) GetAuthenticationCredentialsConfigParams(contextid string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
	}
	return u.c.Request("users/view/getAuthenticationCredentialsConfigParams/", m)
}

func (u Users) GetAuthenticationCredentials(contextid string, userid string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
		"userId":    userid,
	}
	return u.c.Request("users/view/getAuthenticationCredentials/", m)
}

func (u Users) NewUser(contextid string, name string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
		"name":      name,
	}
	return u.c.Request("users/action/newUser/", m)
}

func (u Users) RemoveUser(contextid string, userid string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
		"userId":    userid,
	}
	return u.c.Request("users/action/removeUser/", m)
}

func (u Users) SetUserEnabled(contextid string, userid string, enabled string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
		"userId":    userid,
		"enabled":   enabled,
	}
	return u.c.Request("users/action/setUserEnabled/", m)
}

func (u Users) SetUserName(contextid string, userid string, name string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
		"userId":    userid,
		"name":      name,
	}
	return u.c.Request("users/action/setUserName/", m)
}

func (u Users) SetAuthenticationCredentials(contextid string, userid string, authcredentialsconfigparams string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId":                   contextid,
		"userId":                      userid,
		"authCredentialsConfigParams": authcredentialsconfigparams,
	}
	return u.c.Request("users/action/setAuthenticationCredentials/", m)
}
