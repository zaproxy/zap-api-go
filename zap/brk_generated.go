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

type Break struct {
	c *Client
}

// Returns True if ZAP will break on both requests and responses
func (b Break) IsBreakAll() (map[string]interface{}, error) {
	return b.c.Request("break/view/isBreakAll/", nil)
}

// Returns True if ZAP will break on requests
func (b Break) IsBreakRequest() (map[string]interface{}, error) {
	return b.c.Request("break/view/isBreakRequest/", nil)
}

// Returns True if ZAP will break on responses
func (b Break) IsBreakResponse() (map[string]interface{}, error) {
	return b.c.Request("break/view/isBreakResponse/", nil)
}

// Returns the HTTP message currently intercepted (if any)
func (b Break) HttpMessage() (map[string]interface{}, error) {
	return b.c.Request("break/view/httpMessage/", nil)
}

// Controls the global break functionality. The type may be one of: http-all, http-request or http-response. The state may be true (for turning break on for the specified type) or false (for turning break off). Scope is not currently used.
func (b Break) Brk(t string, state string, scope string) (map[string]interface{}, error) {
	m := map[string]string{
		"type":  t,
		"state": state,
		"scope": scope,
	}
	return b.c.Request("break/action/break/", m)
}

// Overwrites the currently intercepted message with the data provided
func (b Break) SetHttpMessage(httpheader string, httpbody string) (map[string]interface{}, error) {
	m := map[string]string{
		"httpHeader": httpheader,
		"httpBody":   httpbody,
	}
	return b.c.Request("break/action/setHttpMessage/", m)
}

// Submits the currently intercepted message and unsets the global request/response breakpoints
func (b Break) Cont() (map[string]interface{}, error) {
	return b.c.Request("break/action/continue/", nil)
}

// Submits the currently intercepted message, the next request or response will automatically be intercepted
func (b Break) Step() (map[string]interface{}, error) {
	return b.c.Request("break/action/step/", nil)
}

// Drops the currently intercepted message
func (b Break) Drop() (map[string]interface{}, error) {
	return b.c.Request("break/action/drop/", nil)
}

// Adds a custom HTTP breakpoint. The string is the string to match. Location may be one of: url, request_header, request_body, response_header or response_body. Match may be: contains or regex. Inverse (match) may be true or false. Lastly, ignorecase (when matching the string) may be true or false.
func (b Break) AddHttpBreakpoint(str string, location string, match string, inverse string, ignorecase string) (map[string]interface{}, error) {
	m := map[string]string{
		"string":     str,
		"location":   location,
		"match":      match,
		"inverse":    inverse,
		"ignorecase": ignorecase,
	}
	return b.c.Request("break/action/addHttpBreakpoint/", m)
}

// Removes the specified breakpoint
func (b Break) RemoveHttpBreakpoint(str string, location string, match string, inverse string, ignorecase string) (map[string]interface{}, error) {
	m := map[string]string{
		"string":     str,
		"location":   location,
		"match":      match,
		"inverse":    inverse,
		"ignorecase": ignorecase,
	}
	return b.c.Request("break/action/removeHttpBreakpoint/", m)
}
