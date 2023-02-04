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

import "strconv"

type Graphql struct {
	c *Client
}

// Returns how arguments are currently specified.
//
// This component is optional and therefore the API will only work if it is installed
func (g Graphql) OptionArgsType() (map[string]interface{}, error) {
	return g.c.Request("graphql/view/optionArgsType/", nil)
}

// Returns whether or not lenient maximum query generation depth is enabled.
//
// This component is optional and therefore the API will only work if it is installed
func (g Graphql) OptionLenientMaxQueryDepthEnabled() (map[string]interface{}, error) {
	return g.c.Request("graphql/view/optionLenientMaxQueryDepthEnabled/", nil)
}

// Returns the current maximum additional query generation depth.
//
// This component is optional and therefore the API will only work if it is installed
func (g Graphql) OptionMaxAdditionalQueryDepth() (map[string]interface{}, error) {
	return g.c.Request("graphql/view/optionMaxAdditionalQueryDepth/", nil)
}

// Returns the current maximum arguments generation depth.
//
// This component is optional and therefore the API will only work if it is installed
func (g Graphql) OptionMaxArgsDepth() (map[string]interface{}, error) {
	return g.c.Request("graphql/view/optionMaxArgsDepth/", nil)
}

// Returns the current maximum query generation depth.
//
// This component is optional and therefore the API will only work if it is installed
func (g Graphql) OptionMaxQueryDepth() (map[string]interface{}, error) {
	return g.c.Request("graphql/view/optionMaxQueryDepth/", nil)
}

// Returns whether or not optional arguments are currently specified.
//
// This component is optional and therefore the API will only work if it is installed
func (g Graphql) OptionOptionalArgsEnabled() (map[string]interface{}, error) {
	return g.c.Request("graphql/view/optionOptionalArgsEnabled/", nil)
}

// Returns the current level for which a single query is generated.
//
// This component is optional and therefore the API will only work if it is installed
func (g Graphql) OptionQuerySplitType() (map[string]interface{}, error) {
	return g.c.Request("graphql/view/optionQuerySplitType/", nil)
}

// Returns the current request method.
//
// This component is optional and therefore the API will only work if it is installed
func (g Graphql) OptionRequestMethod() (map[string]interface{}, error) {
	return g.c.Request("graphql/view/optionRequestMethod/", nil)
}

// Imports a GraphQL Schema from a File.
//
// This component is optional and therefore the API will only work if it is installed
func (g Graphql) ImportFile(endurl string, file string) (map[string]interface{}, error) {
	m := map[string]string{
		"endurl": endurl,
		"file":   file,
	}
	return g.c.Request("graphql/action/importFile/", m)
}

// Imports a GraphQL Schema from a URL.
//
// This component is optional and therefore the API will only work if it is installed
func (g Graphql) ImportUrl(endurl string, url string) (map[string]interface{}, error) {
	m := map[string]string{
		"endurl": endurl,
		"url":    url,
	}
	return g.c.Request("graphql/action/importUrl/", m)
}

// Sets how arguments are specified.
//
// This component is optional and therefore the API will only work if it is installed
func (g Graphql) SetOptionArgsType(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return g.c.Request("graphql/action/setOptionArgsType/", m)
}

// Sets the level for which a single query is generated.
//
// This component is optional and therefore the API will only work if it is installed
func (g Graphql) SetOptionQuerySplitType(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return g.c.Request("graphql/action/setOptionQuerySplitType/", m)
}

// Sets the request method.
//
// This component is optional and therefore the API will only work if it is installed
func (g Graphql) SetOptionRequestMethod(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return g.c.Request("graphql/action/setOptionRequestMethod/", m)
}

// Sets whether or not Maximum Query Depth is enforced leniently.
//
// This component is optional and therefore the API will only work if it is installed
func (g Graphql) SetOptionLenientMaxQueryDepthEnabled(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return g.c.Request("graphql/action/setOptionLenientMaxQueryDepthEnabled/", m)
}

// Sets the maximum additional query generation depth (used if enforced leniently).
//
// This component is optional and therefore the API will only work if it is installed
func (g Graphql) SetOptionMaxAdditionalQueryDepth(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return g.c.Request("graphql/action/setOptionMaxAdditionalQueryDepth/", m)
}

// Sets the maximum arguments generation depth.
//
// This component is optional and therefore the API will only work if it is installed
func (g Graphql) SetOptionMaxArgsDepth(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return g.c.Request("graphql/action/setOptionMaxArgsDepth/", m)
}

// Sets the maximum query generation depth.
//
// This component is optional and therefore the API will only work if it is installed
func (g Graphql) SetOptionMaxQueryDepth(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return g.c.Request("graphql/action/setOptionMaxQueryDepth/", m)
}

// Sets whether or not Optional Arguments should be specified.
//
// This component is optional and therefore the API will only work if it is installed
func (g Graphql) SetOptionOptionalArgsEnabled(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return g.c.Request("graphql/action/setOptionOptionalArgsEnabled/", m)
}
