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

type AccessControl struct {
	c *Client
}

// Gets the Access Control scan progress (percentage integer) for the given context ID.
//
// This component is optional and therefore the API will only work if it is installed
func (a AccessControl) GetScanProgress(contextid string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
	}
	return a.c.Request("accessControl/view/getScanProgress/", m)
}

// Gets the Access Control scan status (description string) for the given context ID.
//
// This component is optional and therefore the API will only work if it is installed
func (a AccessControl) GetScanStatus(contextid string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
	}
	return a.c.Request("accessControl/view/getScanStatus/", m)
}

// Starts an Access Control scan with the given context ID and user ID. (Optional parameters: user ID for Unauthenticated user, boolean identifying whether or not Alerts are raised, and the Risk level for the Alerts.) [This assumes the Access Control rules were previously established via ZAP gui and the necessary Context exported/imported.]
//
// This component is optional and therefore the API will only work if it is installed
func (a AccessControl) Scan(contextid string, userid string, scanasunauthuser string, raisealert string, alertrisklevel string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId":        contextid,
		"userId":           userid,
		"scanAsUnAuthUser": scanasunauthuser,
		"raiseAlert":       raisealert,
		"alertRiskLevel":   alertrisklevel,
	}
	return a.c.Request("accessControl/action/scan/", m)
}

// Generates an Access Control report for the given context ID and saves it based on the provided filename (path).
//
// This component is optional and therefore the API will only work if it is installed
func (a AccessControl) WriteHTMLreport(contextid string, filename string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
		"fileName":  filename,
	}
	return a.c.Request("accessControl/action/writeHTMLreport/", m)
}
