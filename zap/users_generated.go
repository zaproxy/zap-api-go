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

// Gets a list of users that belong to the context with the given ID, or all users if none provided.
func (u Users) UsersList(contextid string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
	}
	return u.c.Request("users/view/usersList/", m)
}

// Gets the data of the user with the given ID that belongs to the context with the given ID.
func (u Users) GetUserById(contextid string, userid string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
		"userId":    userid,
	}
	return u.c.Request("users/view/getUserById/", m)
}

// Gets the configuration parameters for the credentials of the context with the given ID.
func (u Users) GetAuthenticationCredentialsConfigParams(contextid string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
	}
	return u.c.Request("users/view/getAuthenticationCredentialsConfigParams/", m)
}

// Gets the authentication credentials of the user with given ID that belongs to the context with the given ID.
func (u Users) GetAuthenticationCredentials(contextid string, userid string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
		"userId":    userid,
	}
	return u.c.Request("users/view/getAuthenticationCredentials/", m)
}

// Gets the authentication state information for the user identified by the Context and User Ids.
func (u Users) GetAuthenticationState(contextid string, userid string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
		"userId":    userid,
	}
	return u.c.Request("users/view/getAuthenticationState/", m)
}

// Gets the authentication session information for the user identified by the Context and User Ids, e.g. cookies and realm credentials.
func (u Users) GetAuthenticationSession(contextid string, userid string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
		"userId":    userid,
	}
	return u.c.Request("users/view/getAuthenticationSession/", m)
}

// Creates a new user with the given name for the context with the given ID.
func (u Users) NewUser(contextid string, name string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
		"name":      name,
	}
	return u.c.Request("users/action/newUser/", m)
}

// Removes the user with the given ID that belongs to the context with the given ID.
func (u Users) RemoveUser(contextid string, userid string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
		"userId":    userid,
	}
	return u.c.Request("users/action/removeUser/", m)
}

// Sets whether or not the user, with the given ID that belongs to the context with the given ID, should be enabled.
func (u Users) SetUserEnabled(contextid string, userid string, enabled string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
		"userId":    userid,
		"enabled":   enabled,
	}
	return u.c.Request("users/action/setUserEnabled/", m)
}

// Renames the user with the given ID that belongs to the context with the given ID.
func (u Users) SetUserName(contextid string, userid string, name string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
		"userId":    userid,
		"name":      name,
	}
	return u.c.Request("users/action/setUserName/", m)
}

// Sets the authentication credentials for the user with the given ID that belongs to the context with the given ID.
func (u Users) SetAuthenticationCredentials(contextid string, userid string, authcredentialsconfigparams string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId":                   contextid,
		"userId":                      userid,
		"authCredentialsConfigParams": authcredentialsconfigparams,
	}
	return u.c.Request("users/action/setAuthenticationCredentials/", m)
}

// Tries to authenticate as the identified user, returning the authentication request and whether it appears to have succeeded.
func (u Users) AuthenticateAsUser(contextid string, userid string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
		"userId":    userid,
	}
	return u.c.Request("users/action/authenticateAsUser/", m)
}

// Tries to poll as the identified user, returning the authentication request and whether it appears to have succeeded. This will only work if the polling verification strategy has been configured.
func (u Users) PollAsUser(contextid string, userid string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
		"userId":    userid,
	}
	return u.c.Request("users/action/pollAsUser/", m)
}

// Sets fields in the authentication state for the user identified by the Context and User Ids.
func (u Users) SetAuthenticationState(contextid string, userid string, lastpollresult string, lastpolltimeinms string, requestssincelastpoll string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId":             contextid,
		"userId":                userid,
		"lastPollResult":        lastpollresult,
		"lastPollTimeInMs":      lastpolltimeinms,
		"requestsSinceLastPoll": requestssincelastpoll,
	}
	return u.c.Request("users/action/setAuthenticationState/", m)
}

// Sets the specified cookie for the user identified by the Context and User Ids.
func (u Users) SetCookie(contextid string, userid string, domain string, name string, value string, path string, secure string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
		"userId":    userid,
		"domain":    domain,
		"name":      name,
		"value":     value,
		"path":      path,
		"secure":    secure,
	}
	return u.c.Request("users/action/setCookie/", m)
}
