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

type Exim struct {
	c *Client
}

// Imports a HAR file.
//
// This component is optional and therefore the API will only work if it is installed
func (e Exim) ImportHar(filepath string) (map[string]interface{}, error) {
	m := map[string]string{
		"filePath": filepath,
	}
	return e.c.Request("exim/action/importHar/", m)
}

// Imports URLs (one per line) from the file with the given file system path.
//
// This component is optional and therefore the API will only work if it is installed
func (e Exim) ImportUrls(filepath string) (map[string]interface{}, error) {
	m := map[string]string{
		"filePath": filepath,
	}
	return e.c.Request("exim/action/importUrls/", m)
}

// Imports previously exported ZAP messages from the file with the given file system path.
//
// This component is optional and therefore the API will only work if it is installed
func (e Exim) ImportZapLogs(filepath string) (map[string]interface{}, error) {
	m := map[string]string{
		"filePath": filepath,
	}
	return e.c.Request("exim/action/importZapLogs/", m)
}

// Imports ModSecurity2 logs from the file with the given file system path.
//
// This component is optional and therefore the API will only work if it is installed
func (e Exim) ImportModsec2Logs(filepath string) (map[string]interface{}, error) {
	m := map[string]string{
		"filePath": filepath,
	}
	return e.c.Request("exim/action/importModsec2Logs/", m)
}

// Gets the HTTP messages sent through/by ZAP, in HAR format, optionally filtered by URL and paginated with 'start' position and 'count' of messages
//
// This component is optional and therefore the API will only work if it is installed
func (e Exim) ExportHar(baseurl string, start string, count string) ([]byte, error) {
	m := map[string]string{
		"baseurl": baseurl,
		"start": start,
		"count": count,
	}
	return e.c.RequestOther("exim/other/exportHar/", m)
}

// Gets the HTTP messages with the given IDs, in HAR format.
//
// This component is optional and therefore the API will only work if it is installed
func (e Exim) ExportHarById(ids string) ([]byte, error) {
	m := map[string]string{
		"ids": ids,
	}
	return e.c.RequestOther("exim/other/exportHarById/", m)
}

// Sends the first HAR request entry, optionally following redirections. Returns, in HAR format, the request sent and response received and followed redirections, if any. The Mode is enforced when sending the request (and following redirections), custom manual requests are not allowed in 'Safe' mode nor in 'Protected' mode if out of scope.
//
// This component is optional and therefore the API will only work if it is installed
func (e Exim) SendHarRequest(request string, followredirects string) ([]byte, error) {
	m := map[string]string{
		"request": request,
		"followRedirects": followredirects,
	}
	return e.c.RequestOther("exim/other/sendHarRequest/", m)
}

