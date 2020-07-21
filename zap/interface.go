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

// TODO: auto generate this file
package zap

// Interface defines the interface a ZAP client should implement
type Interface interface {
	Acsrf() *Acsrf
	Ascan() *Ascan
	Authentication() *Authentication
	Authorization() *Authorization
	Autoupdate() *Autoupdate
	Break() *Break
	Context() *Context
	Core() *Core
	ForcedUser() *ForcedUser
	HttpSessions() *HttpSessions
	Openapi() *Openapi
	Params() *Params
	Pscan() *Pscan
	Script() *Script
	Search() *Search
	Spider() *Spider
	Stats() *Stats
	Users() *Users
}

// Acsrf() returns a Acsrf client
func (c *Client) Acsrf() *Acsrf {
	return &Acsrf{c}
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

// ForcedUser() returns a ForcedUser client
func (c *Client) ForcedUser() *ForcedUser {
	return &ForcedUser{c}
}

// HttpSessions() returns a HttpSessions client
func (c *Client) HttpSessions() *HttpSessions {
	return &HttpSessions{c}
}

// Openapi() returns a Openapi clinet
func (c *Client) Openapi() *Openapi {
	return &Openapi{c}
}

// Params() returns a Params client
func (c *Client) Params() *Params {
	return &Params{c}
}

// Pscan() returns a Pscan client
func (c *Client) Pscan() *Pscan {
	return &Pscan{c}
}

// Script() returns a Script client
func (c *Client) Script() *Script {
	return &Script{c}
}

// Search() returns a Search client
func (c *Client) Search() *Search {
	return &Search{c}
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
