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

type Authentication struct {
	c *Client
}

// Gets the name of the authentication methods.
func (a Authentication) GetSupportedAuthenticationMethods() (map[string]interface{}, error) {
	return a.c.Request("authentication/view/getSupportedAuthenticationMethods/", nil)
}

// Gets the configuration parameters for the authentication method with the given name.
func (a Authentication) GetAuthenticationMethodConfigParams(authmethodname string) (map[string]interface{}, error) {
	m := map[string]string{
		"authMethodName": authmethodname,
	}
	return a.c.Request("authentication/view/getAuthenticationMethodConfigParams/", m)
}

// Gets the name of the authentication method for the context with the given ID.
func (a Authentication) GetAuthenticationMethod(contextid string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
	}
	return a.c.Request("authentication/view/getAuthenticationMethod/", m)
}

// Gets the logged in indicator for the context with the given ID.
func (a Authentication) GetLoggedInIndicator(contextid string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
	}
	return a.c.Request("authentication/view/getLoggedInIndicator/", m)
}

// Gets the logged out indicator for the context with the given ID.
func (a Authentication) GetLoggedOutIndicator(contextid string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
	}
	return a.c.Request("authentication/view/getLoggedOutIndicator/", m)
}

// Sets the authentication method for the context with the given ID.
func (a Authentication) SetAuthenticationMethod(contextid string, authmethodname string, authmethodconfigparams string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId":              contextid,
		"authMethodName":         authmethodname,
		"authMethodConfigParams": authmethodconfigparams,
	}
	return a.c.Request("authentication/action/setAuthenticationMethod/", m)
}

// Sets the logged in indicator for the context with the given ID.
func (a Authentication) SetLoggedInIndicator(contextid string, loggedinindicatorregex string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId":              contextid,
		"loggedInIndicatorRegex": loggedinindicatorregex,
	}
	return a.c.Request("authentication/action/setLoggedInIndicator/", m)
}

// Sets the logged out indicator for the context with the given ID.
func (a Authentication) SetLoggedOutIndicator(contextid string, loggedoutindicatorregex string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId":               contextid,
		"loggedOutIndicatorRegex": loggedoutindicatorregex,
	}
	return a.c.Request("authentication/action/setLoggedOutIndicator/", m)
}
