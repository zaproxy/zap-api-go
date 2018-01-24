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

type Script struct {
	c *Client
}

// Lists the script engines available
func (s Script) ListEngines() (map[string]interface{}, error) {
	return s.c.Request("script/view/listEngines/", nil)
}

// Lists the scripts available, with its engine, name, description, type and error state.
func (s Script) ListScripts() (map[string]interface{}, error) {
	return s.c.Request("script/view/listScripts/", nil)
}

// Enables the script with the given name
func (s Script) Enable(scriptname string) (map[string]interface{}, error) {
	m := map[string]string{
		"scriptName": scriptname,
	}
	return s.c.Request("script/action/enable/", m)
}

// Disables the script with the given name
func (s Script) Disable(scriptname string) (map[string]interface{}, error) {
	m := map[string]string{
		"scriptName": scriptname,
	}
	return s.c.Request("script/action/disable/", m)
}

// Loads a script into ZAP from the given local file, with the given name, type and engine, optionally with a description, and a charset name to read the script (the charset name is required if the script is not in UTF-8, for example, in ISO-8859-1).
func (s Script) Load(scriptname string, scripttype string, scriptengine string, filename string, scriptdescription string, charset string) (map[string]interface{}, error) {
	m := map[string]string{
		"scriptName":        scriptname,
		"scriptType":        scripttype,
		"scriptEngine":      scriptengine,
		"fileName":          filename,
		"scriptDescription": scriptdescription,
		"charset":           charset,
	}
	return s.c.Request("script/action/load/", m)
}

// Removes the script with the given name
func (s Script) Remove(scriptname string) (map[string]interface{}, error) {
	m := map[string]string{
		"scriptName": scriptname,
	}
	return s.c.Request("script/action/remove/", m)
}

// Runs the stand alone script with the give name
func (s Script) RunStandAloneScript(scriptname string) (map[string]interface{}, error) {
	m := map[string]string{
		"scriptName": scriptname,
	}
	return s.c.Request("script/action/runStandAloneScript/", m)
}
