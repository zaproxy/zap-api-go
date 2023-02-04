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

type Wappalyzer struct {
	c *Client
}

// Lists all the sites recognized by the wappalyzer addon.
//
// This component is optional and therefore the API will only work if it is installed
func (w Wappalyzer) ListSites() (map[string]interface{}, error) {
	return w.c.Request("wappalyzer/view/listSites/", nil)
}

// Lists all sites and their associated applications (technologies).
//
// This component is optional and therefore the API will only work if it is installed
func (w Wappalyzer) ListAll() (map[string]interface{}, error) {
	return w.c.Request("wappalyzer/view/listAll/", nil)
}

// Lists all the applications (technologies) associated with a specific site.
//
// This component is optional and therefore the API will only work if it is installed
func (w Wappalyzer) ListSite(site string) (map[string]interface{}, error) {
	m := map[string]string{
		"site": site,
	}
	return w.c.Request("wappalyzer/view/listSite/", m)
}
