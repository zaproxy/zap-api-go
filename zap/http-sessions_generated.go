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

type HttpSessions struct {
	c *Client
}

// Gets all of the sites that have sessions.
func (h HttpSessions) Sites() (map[string]interface{}, error) {
	return h.c.Request("httpSessions/view/sites/", nil)
}

// Gets the sessions for the given site. Optionally returning just the session with the given name.
func (h HttpSessions) Sessions(site string, session string) (map[string]interface{}, error) {
	m := map[string]string{
		"site":    site,
		"session": session,
	}
	return h.c.Request("httpSessions/view/sessions/", m)
}

// Gets the name of the active session for the given site.
func (h HttpSessions) ActiveSession(site string) (map[string]interface{}, error) {
	m := map[string]string{
		"site": site,
	}
	return h.c.Request("httpSessions/view/activeSession/", m)
}

// Gets the names of the session tokens for the given site.
func (h HttpSessions) SessionTokens(site string) (map[string]interface{}, error) {
	m := map[string]string{
		"site": site,
	}
	return h.c.Request("httpSessions/view/sessionTokens/", m)
}

// Gets the default session tokens.
func (h HttpSessions) DefaultSessionTokens() (map[string]interface{}, error) {
	return h.c.Request("httpSessions/view/defaultSessionTokens/", nil)
}

// Creates an empty session for the given site. Optionally with the given name.
func (h HttpSessions) CreateEmptySession(site string, session string) (map[string]interface{}, error) {
	m := map[string]string{
		"site":    site,
		"session": session,
	}
	return h.c.Request("httpSessions/action/createEmptySession/", m)
}

// Removes the session from the given site.
func (h HttpSessions) RemoveSession(site string, session string) (map[string]interface{}, error) {
	m := map[string]string{
		"site":    site,
		"session": session,
	}
	return h.c.Request("httpSessions/action/removeSession/", m)
}

// Sets the given session as active for the given site.
func (h HttpSessions) SetActiveSession(site string, session string) (map[string]interface{}, error) {
	m := map[string]string{
		"site":    site,
		"session": session,
	}
	return h.c.Request("httpSessions/action/setActiveSession/", m)
}

// Unsets the active session of the given site.
func (h HttpSessions) UnsetActiveSession(site string) (map[string]interface{}, error) {
	m := map[string]string{
		"site": site,
	}
	return h.c.Request("httpSessions/action/unsetActiveSession/", m)
}

// Adds the session token to the given site.
func (h HttpSessions) AddSessionToken(site string, sessiontoken string) (map[string]interface{}, error) {
	m := map[string]string{
		"site":         site,
		"sessionToken": sessiontoken,
	}
	return h.c.Request("httpSessions/action/addSessionToken/", m)
}

// Removes the session token from the given site.
func (h HttpSessions) RemoveSessionToken(site string, sessiontoken string) (map[string]interface{}, error) {
	m := map[string]string{
		"site":         site,
		"sessionToken": sessiontoken,
	}
	return h.c.Request("httpSessions/action/removeSessionToken/", m)
}

// Sets the value of the session token of the given session for the given site.
func (h HttpSessions) SetSessionTokenValue(site string, session string, sessiontoken string, tokenvalue string) (map[string]interface{}, error) {
	m := map[string]string{
		"site":         site,
		"session":      session,
		"sessionToken": sessiontoken,
		"tokenValue":   tokenvalue,
	}
	return h.c.Request("httpSessions/action/setSessionTokenValue/", m)
}

// Renames the session of the given site.
func (h HttpSessions) RenameSession(site string, oldsessionname string, newsessionname string) (map[string]interface{}, error) {
	m := map[string]string{
		"site":           site,
		"oldSessionName": oldsessionname,
		"newSessionName": newsessionname,
	}
	return h.c.Request("httpSessions/action/renameSession/", m)
}

// Adds a default session token with the given name and enabled state.
func (h HttpSessions) AddDefaultSessionToken(sessiontoken string, tokenenabled string) (map[string]interface{}, error) {
	m := map[string]string{
		"sessionToken": sessiontoken,
		"tokenEnabled": tokenenabled,
	}
	return h.c.Request("httpSessions/action/addDefaultSessionToken/", m)
}

// Sets whether or not the default session token with the given name is enabled.
func (h HttpSessions) SetDefaultSessionTokenEnabled(sessiontoken string, tokenenabled string) (map[string]interface{}, error) {
	m := map[string]string{
		"sessionToken": sessiontoken,
		"tokenEnabled": tokenenabled,
	}
	return h.c.Request("httpSessions/action/setDefaultSessionTokenEnabled/", m)
}

// Removes the default session token with the given name.
func (h HttpSessions) RemoveDefaultSessionToken(sessiontoken string) (map[string]interface{}, error) {
	m := map[string]string{
		"sessionToken": sessiontoken,
	}
	return h.c.Request("httpSessions/action/removeDefaultSessionToken/", m)
}
