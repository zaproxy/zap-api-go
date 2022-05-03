// Zed Attack Proxy (ZAP) and its related class files.
//
// ZAP is an HTTP/HTTPS proxy for assessing web application security.
//
// Copyright 2022 the ZAP development team
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

type LocalProxies struct {
	c *Client
}

// Gets all of the additional proxies that have been configured.
func (l LocalProxies) AdditionalProxies() (map[string]interface{}, error) {
	return l.c.Request("localProxies/view/additionalProxies/", nil)
}

// Adds an new proxy using the details supplied.
func (l LocalProxies) AddAdditionalProxy(address string, port string, behindnat string, alwaysdecodezip string, removeunsupportedencodings string) (map[string]interface{}, error) {
	m := map[string]string{
		"address":                    address,
		"port":                       port,
		"behindNat":                  behindnat,
		"alwaysDecodeZip":            alwaysdecodezip,
		"removeUnsupportedEncodings": removeunsupportedencodings,
	}
	return l.c.Request("localProxies/action/addAdditionalProxy/", m)
}

// Removes the additional proxy with the specified address and port.
func (l LocalProxies) RemoveAdditionalProxy(address string, port string) (map[string]interface{}, error) {
	m := map[string]string{
		"address": address,
		"port":    port,
	}
	return l.c.Request("localProxies/action/removeAdditionalProxy/", m)
}
