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

type Authorization struct {
	c *Client
}

// Obtains all the configuration of the authorization detection method that is currently set for a context.
func (a Authorization) GetAuthorizationDetectionMethod(contextid string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
	}
	return a.c.Request("authorization/view/getAuthorizationDetectionMethod/", m)
}

// Sets the authorization detection method for a context as one that identifies un-authorized messages based on: the message's status code or a regex pattern in the response's header or body. Also, whether all conditions must match or just some can be specified via the logicalOperator parameter, which accepts two values: "AND" (default), "OR".  
func (a Authorization) SetBasicAuthorizationDetectionMethod(contextid string, headerregex string, bodyregex string, statuscode string, logicaloperator string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextId": contextid,
		"headerRegex": headerregex,
		"bodyRegex": bodyregex,
		"statusCode": statuscode,
		"logicalOperator": logicaloperator,
	}
	return a.c.Request("authorization/action/setBasicAuthorizationDetectionMethod/", m)
}

