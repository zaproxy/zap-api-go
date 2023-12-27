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

// Lists the script types available.
func (s Script) ListTypes() (map[string]interface{}, error) {
	return s.c.Request("script/view/listTypes/", nil)
}

// Lists the scripts available, with its engine, name, description, type and error state.
func (s Script) ListScripts() (map[string]interface{}, error) {
	return s.c.Request("script/view/listScripts/", nil)
}

// Gets the value of the global variable with the given key. Returns an API error (DOES_NOT_EXIST) if no value was previously set.
func (s Script) GlobalVar(varkey string) (map[string]interface{}, error) {
	m := map[string]string{
		"varKey": varkey,
	}
	return s.c.Request("script/view/globalVar/", m)
}

// Gets the value (string representation) of a global custom variable. Returns an API error (DOES_NOT_EXIST) if no value was previously set.
func (s Script) GlobalCustomVar(varkey string) (map[string]interface{}, error) {
	m := map[string]string{
		"varKey": varkey,
	}
	return s.c.Request("script/view/globalCustomVar/", m)
}

// Gets all the global variables (key/value pairs).
func (s Script) GlobalVars() (map[string]interface{}, error) {
	return s.c.Request("script/view/globalVars/", nil)
}

// Gets all the global custom variables (key/value pairs, the value is the string representation).
func (s Script) GlobalCustomVars() (map[string]interface{}, error) {
	return s.c.Request("script/view/globalCustomVars/", nil)
}

// Gets the value of the variable with the given key for the given script. Returns an API error (DOES_NOT_EXIST) if no script with the given name exists or if no value was previously set.
func (s Script) ScriptVar(scriptname string, varkey string) (map[string]interface{}, error) {
	m := map[string]string{
		"scriptName": scriptname,
		"varKey":     varkey,
	}
	return s.c.Request("script/view/scriptVar/", m)
}

// Gets the value (string representation) of a custom variable. Returns an API error (DOES_NOT_EXIST) if no script with the given name exists or if no value was previously set.
func (s Script) ScriptCustomVar(scriptname string, varkey string) (map[string]interface{}, error) {
	m := map[string]string{
		"scriptName": scriptname,
		"varKey":     varkey,
	}
	return s.c.Request("script/view/scriptCustomVar/", m)
}

// Gets all the variables (key/value pairs) of the given script. Returns an API error (DOES_NOT_EXIST) if no script with the given name exists.
func (s Script) ScriptVars(scriptname string) (map[string]interface{}, error) {
	m := map[string]string{
		"scriptName": scriptname,
	}
	return s.c.Request("script/view/scriptVars/", m)
}

// Gets all the custom variables (key/value pairs, the value is the string representation) of a script. Returns an API error (DOES_NOT_EXIST) if no script with the given name exists.
func (s Script) ScriptCustomVars(scriptname string) (map[string]interface{}, error) {
	m := map[string]string{
		"scriptName": scriptname,
	}
	return s.c.Request("script/view/scriptCustomVars/", m)
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

// Runs the stand alone script with the given name
func (s Script) RunStandAloneScript(scriptname string) (map[string]interface{}, error) {
	m := map[string]string{
		"scriptName": scriptname,
	}
	return s.c.Request("script/action/runStandAloneScript/", m)
}

// Clears the global variable with the given key.
func (s Script) ClearGlobalVar(varkey string) (map[string]interface{}, error) {
	m := map[string]string{
		"varKey": varkey,
	}
	return s.c.Request("script/action/clearGlobalVar/", m)
}

// Clears a global custom variable.
func (s Script) ClearGlobalCustomVar(varkey string) (map[string]interface{}, error) {
	m := map[string]string{
		"varKey": varkey,
	}
	return s.c.Request("script/action/clearGlobalCustomVar/", m)
}

// Clears the global variables.
func (s Script) ClearGlobalVars() (map[string]interface{}, error) {
	return s.c.Request("script/action/clearGlobalVars/", nil)
}

// Clears the variable with the given key of the given script. Returns an API error (DOES_NOT_EXIST) if no script with the given name exists.
func (s Script) ClearScriptVar(scriptname string, varkey string) (map[string]interface{}, error) {
	m := map[string]string{
		"scriptName": scriptname,
		"varKey":     varkey,
	}
	return s.c.Request("script/action/clearScriptVar/", m)
}

// Clears a script custom variable.
func (s Script) ClearScriptCustomVar(scriptname string, varkey string) (map[string]interface{}, error) {
	m := map[string]string{
		"scriptName": scriptname,
		"varKey":     varkey,
	}
	return s.c.Request("script/action/clearScriptCustomVar/", m)
}

// Clears the variables of the given script. Returns an API error (DOES_NOT_EXIST) if no script with the given name exists.
func (s Script) ClearScriptVars(scriptname string) (map[string]interface{}, error) {
	m := map[string]string{
		"scriptName": scriptname,
	}
	return s.c.Request("script/action/clearScriptVars/", m)
}

// Sets the value of the variable with the given key of the given script. Returns an API error (DOES_NOT_EXIST) if no script with the given name exists.
func (s Script) SetScriptVar(scriptname string, varkey string, varvalue string) (map[string]interface{}, error) {
	m := map[string]string{
		"scriptName": scriptname,
		"varKey":     varkey,
		"varValue":   varvalue,
	}
	return s.c.Request("script/action/setScriptVar/", m)
}

// Sets the value of the global variable with the given key.
func (s Script) SetGlobalVar(varkey string, varvalue string) (map[string]interface{}, error) {
	m := map[string]string{
		"varKey":   varkey,
		"varValue": varvalue,
	}
	return s.c.Request("script/action/setGlobalVar/", m)
}
