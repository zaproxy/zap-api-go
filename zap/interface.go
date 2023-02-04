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

// TODO: auto generate this file
package zap

// Interface defines the interface a ZAP client should implement
type Interface interface {
	AccessControl() *AccessControl
	Acsrf() *Acsrf
	AjaxSpider() *AjaxSpider
	Alert() *Alert
	AlertFilter() *AlertFilter
	Ascan() *Ascan
	Authentication() *Authentication
	Authorization() *Authorization
	Automation() *Automation
	Autoupdate() *Autoupdate
	Break() *Break
	Context() *Context
	Core() *Core
	Exim() *Exim
	Exportreport() *Exportreport
	ForcedUser() *ForcedUser
	Graphql() *Graphql
	HttpSessions() *HttpSessions
	ImportLogFiles() *ImportLogFiles
	Importurls() *Importurls
	LocalProxies() *LocalProxies
	Network() *Network
	Openapi() *Openapi
	Params() *Params
	Pnh() *Pnh
	Pscan() *Pscan
	Replacer() *Replacer
	Reports() *Reports
	Retest() *Retest
	Reveal() *Reveal
	Revisit() *Revisit
	RuleConfig() *RuleConfig
	Script() *Script
	Search() *Search
	Selenium() *Selenium
	Soap() *Soap
	Spider() *Spider
	Stats() *Stats
	Users() *Users
	Wappalyzer() *Wappalyzer
	Websocket() *Websocket
}

// AccessControl() returns a AccessControl client
func (c *Client) AccessControl() *AccessControl {
	return &AccessControl{c}
}

// Acsrf() returns a Acsrf client
func (c *Client) Acsrf() *Acsrf {
	return &Acsrf{c}
}

// AjaxSpider() returns a AjaxSpider client
func (c *Client) AjaxSpider() *AjaxSpider {
	return &AjaxSpider{c}
}

// Alert() returns a Alert client
func (c *Client) Alert() *Alert {
	return &Alert{c}
}

// AlertFilter() returns a AlertFilter client
func (c *Client) AlertFilter() *AlertFilter {
	return &AlertFilter{c}
}

// Ascan() returns a Ascan client
func (c *Client) Ascan() *Ascan {
	return &Ascan{c}
}

// Authentication() returns a Authentication client
func (c *Client) Authentication() *Authentication {
	return &Authentication{c}
}

// Authorization() returns a Authorization client
func (c *Client) Authorization() *Authorization {
	return &Authorization{c}
}

// Autoupdate returns an Autoupdate client
func (c *Client) Autoupdate() *Autoupdate {
	return &Autoupdate{c}
}

// Automation() returns an Automation client
func (c *Client) Automation() *Automation {
	return &Automation{c}
}

// Break() returns a Break client
func (c *Client) Break() *Break {
	return &Break{c}
}

// Context() returns a Context client
func (c *Client) Context() *Context {
	return &Context{c}
}

// Core() returns a Core client
func (c *Client) Core() *Core {
	return &Core{c}
}

// Exim() returns a Exim client
func (c *Client) Exim() *Exim {
	return &Exim{c}
}

// Exportreport() returns a Exportreport client
func (c *Client) Exportreport() *Exportreport {
	return &Exportreport{c}
}

// ForcedUser() returns a ForcedUser client
func (c *Client) ForcedUser() *ForcedUser {
	return &ForcedUser{c}
}

// Graphql() returns a Graphql client
func (c *Client) Graphql() *Graphql {
	return &Graphql{c}
}

// HttpSessions() returns a HttpSessions client
func (c *Client) HttpSessions() *HttpSessions {
	return &HttpSessions{c}
}

// ImportLogFiles() returns a ImportLogFiles client
func (c *Client) ImportLogFiles() *ImportLogFiles {
	return &ImportLogFiles{c}
}

// Importurls() returns a Importurls client
func (c *Client) Importurls() *Importurls {
	return &Importurls{c}
}

// LocalProxies() returns a LocalProxies client
func (c *Client) LocalProxies() *LocalProxies {
	return &LocalProxies{c}
}

// Network() returns a Network clinet
func (c *Client) Network() *Network {
	return &Network{c}
}

// Openapi() returns a Openapi clinet
func (c *Client) Openapi() *Openapi {
	return &Openapi{c}
}

// Params() returns a Params client
func (c *Client) Params() *Params {
	return &Params{c}
}

// Pnh() returns a Pnh client
func (c *Client) Pnh() *Pnh {
	return &Pnh{c}
}

// Pscan() returns a Pscan client
func (c *Client) Pscan() *Pscan {
	return &Pscan{c}
}

// Replacer() returns a Replacer client
func (c *Client) Replacer() *Replacer {
	return &Replacer{c}
}

// Reports() returns a Reports client
func (c *Client) Reports() *Reports {
	return &Reports{c}
}

// Retest() returns a Retest client
func (c *Client) Retest() *Retest {
	return &Retest{c}
}

// Reveal() returns a Reveal client
func (c *Client) Reveal() *Reveal {
	return &Reveal{c}
}

// Revisit() returns a Revisit client
func (c *Client) Revisit() *Revisit {
	return &Revisit{c}
}

// RuleConfig() returns a RuleConfig client
func (c *Client) RuleConfig() *RuleConfig {
	return &RuleConfig{c}
}

// Script() returns a Script client
func (c *Client) Script() *Script {
	return &Script{c}
}

// Search() returns a Search client
func (c *Client) Search() *Search {
	return &Search{c}
}

// Selenium() returns a Selenium client
func (c *Client) Selenium() *Selenium {
	return &Selenium{c}
}

// Soap() returns a Soap client
func (c *Client) Soap() *Soap {
	return &Soap{c}
}

// Spider() returns a Spider client
func (c *Client) Spider() *Spider {
	return &Spider{c}
}

// Stats() returns a Stats client
func (c *Client) Stats() *Stats {
	return &Stats{c}
}

// Users() returns a Users client
func (c *Client) Users() *Users {
	return &Users{c}
}

// Wappalyzer() returns a Wappalyzer client
func (c *Client) Wappalyzer() *Wappalyzer {
	return &Wappalyzer{c}
}

// Websocket() returns a Websocket client
func (c *Client) Websocket() *Websocket {
	return &Websocket{c}
}
