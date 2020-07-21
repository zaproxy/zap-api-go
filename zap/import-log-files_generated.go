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

type ImportLogFiles struct {
	c *Client
}

// This component is optional and therefore the API will only work if it is installed
func (i ImportLogFiles) ImportZAPLogFromFile(filepath string) (map[string]interface{}, error) {
	m := map[string]string{
		"FilePath": filepath,
	}
	return i.c.Request("importLogFiles/action/ImportZAPLogFromFile/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (i ImportLogFiles) ImportModSecurityLogFromFile(filepath string) (map[string]interface{}, error) {
	m := map[string]string{
		"FilePath": filepath,
	}
	return i.c.Request("importLogFiles/action/ImportModSecurityLogFromFile/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (i ImportLogFiles) ImportZAPHttpRequestResponsePair(httprequest string, httpresponse string) (map[string]interface{}, error) {
	m := map[string]string{
		"HTTPRequest":  httprequest,
		"HTTPResponse": httpresponse,
	}
	return i.c.Request("importLogFiles/action/ImportZAPHttpRequestResponsePair/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (i ImportLogFiles) PostModSecurityAuditEvent(auditeventstring string) (map[string]interface{}, error) {
	m := map[string]string{
		"AuditEventString": auditeventstring,
	}
	return i.c.Request("importLogFiles/action/PostModSecurityAuditEvent/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (i ImportLogFiles) OtherPostModSecurityAuditEvent(auditeventstring string) ([]byte, error) {
	m := map[string]string{
		"AuditEventString": auditeventstring,
	}
	return i.c.RequestOther("importLogFiles/other/OtherPostModSecurityAuditEvent/", m)
}
