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

type Context struct {
	c *Client
}

// List context names of current session
func (c Context) ContextList() (map[string]interface{}, error) {
	return c.c.Request("context/view/contextList/", nil)
}

// List excluded regexs for context
func (c Context) ExcludeRegexs(contextname string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextName": contextname,
	}
	return c.c.Request("context/view/excludeRegexs/", m)
}

// List included regexs for context
func (c Context) IncludeRegexs(contextname string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextName": contextname,
	}
	return c.c.Request("context/view/includeRegexs/", m)
}

// List the information about the named context
func (c Context) Context(contextname string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextName": contextname,
	}
	return c.c.Request("context/view/context/", m)
}

// Lists the names of all built in technologies
func (c Context) TechnologyList() (map[string]interface{}, error) {
	return c.c.Request("context/view/technologyList/", nil)
}

// Lists the names of all technologies included in a context
func (c Context) IncludedTechnologyList(contextname string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextName": contextname,
	}
	return c.c.Request("context/view/includedTechnologyList/", m)
}

// Lists the names of all technologies excluded from a context
func (c Context) ExcludedTechnologyList(contextname string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextName": contextname,
	}
	return c.c.Request("context/view/excludedTechnologyList/", m)
}

// Lists the URLs accessed through/by ZAP, that belong to the context with the given name.
func (c Context) Urls(contextname string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextName": contextname,
	}
	return c.c.Request("context/view/urls/", m)
}

// Add exclude regex to context
func (c Context) ExcludeFromContext(contextname string, regex string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextName": contextname,
		"regex":       regex,
	}
	return c.c.Request("context/action/excludeFromContext/", m)
}

// Add include regex to context
func (c Context) IncludeInContext(contextname string, regex string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextName": contextname,
		"regex":       regex,
	}
	return c.c.Request("context/action/includeInContext/", m)
}

// Set the regexs to include and exclude for a context, both supplied as JSON string arrays
func (c Context) SetContextRegexs(contextname string, incregexs string, excregexs string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextName": contextname,
		"incRegexs":   incregexs,
		"excRegexs":   excregexs,
	}
	return c.c.Request("context/action/setContextRegexs/", m)
}

// Set the checking strategy for a context - this defines how ZAP checks that a request is authenticated
func (c Context) SetContextCheckingStrategy(contextname string, checkingstrategy string, pollurl string, polldata string, pollheaders string, pollfrequency string, pollfrequencyunits string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextName":        contextname,
		"checkingStrategy":   checkingstrategy,
		"pollUrl":            pollurl,
		"pollData":           polldata,
		"pollHeaders":        pollheaders,
		"pollFrequency":      pollfrequency,
		"pollFrequencyUnits": pollfrequencyunits,
	}
	return c.c.Request("context/action/setContextCheckingStrategy/", m)
}

// Creates a new context with the given name in the current session
func (c Context) NewContext(contextname string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextName": contextname,
	}
	return c.c.Request("context/action/newContext/", m)
}

// Removes a context in the current session
func (c Context) RemoveContext(contextname string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextName": contextname,
	}
	return c.c.Request("context/action/removeContext/", m)
}

// Exports the context with the given name to a file. If a relative file path is specified it will be resolved against the "contexts" directory in ZAP "home" dir.
func (c Context) ExportContext(contextname string, contextfile string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextName": contextname,
		"contextFile": contextfile,
	}
	return c.c.Request("context/action/exportContext/", m)
}

// Imports a context from a file. If a relative file path is specified it will be resolved against the "contexts" directory in ZAP "home" dir.
func (c Context) ImportContext(contextfile string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextFile": contextfile,
	}
	return c.c.Request("context/action/importContext/", m)
}

// Includes technologies with the given names, separated by a comma, to a context
func (c Context) IncludeContextTechnologies(contextname string, technologynames string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextName":     contextname,
		"technologyNames": technologynames,
	}
	return c.c.Request("context/action/includeContextTechnologies/", m)
}

// Includes all built in technologies in to a context
func (c Context) IncludeAllContextTechnologies(contextname string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextName": contextname,
	}
	return c.c.Request("context/action/includeAllContextTechnologies/", m)
}

// Excludes technologies with the given names, separated by a comma, from a context
func (c Context) ExcludeContextTechnologies(contextname string, technologynames string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextName":     contextname,
		"technologyNames": technologynames,
	}
	return c.c.Request("context/action/excludeContextTechnologies/", m)
}

// Excludes all built in technologies from a context
func (c Context) ExcludeAllContextTechnologies(contextname string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextName": contextname,
	}
	return c.c.Request("context/action/excludeAllContextTechnologies/", m)
}

// Sets a context to in scope (contexts are in scope by default)
func (c Context) SetContextInScope(contextname string, booleaninscope string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextName":    contextname,
		"booleanInScope": booleaninscope,
	}
	return c.c.Request("context/action/setContextInScope/", m)
}
