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

type Pnh struct {
	c *Client
}

// This component is optional and therefore the API will only work if it is installed
func (p Pnh) Monitor(id string, message string) (map[string]interface{}, error) {
	m := map[string]string{
		"id": id,
		"message": message,
	}
	return p.c.Request("pnh/action/monitor/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (p Pnh) Oracle(id string) (map[string]interface{}, error) {
	m := map[string]string{
		"id": id,
	}
	return p.c.Request("pnh/action/oracle/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (p Pnh) StartMonitoring(url string) (map[string]interface{}, error) {
	m := map[string]string{
		"url": url,
	}
	return p.c.Request("pnh/action/startMonitoring/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (p Pnh) StopMonitoring(id string) (map[string]interface{}, error) {
	m := map[string]string{
		"id": id,
	}
	return p.c.Request("pnh/action/stopMonitoring/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (p Pnh) Pnh() ([]byte, error) {
	return p.c.RequestOther("pnh/other/pnh/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (p Pnh) Manifest() ([]byte, error) {
	return p.c.RequestOther("pnh/other/manifest/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (p Pnh) Service() ([]byte, error) {
	return p.c.RequestOther("pnh/other/service/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (p Pnh) Fx_pnhxpi() ([]byte, error) {
	return p.c.RequestOther("pnh/other/fx_pnh.xpi/", nil)
}

