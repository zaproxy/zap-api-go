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

type Stats struct {
	c *Client
}

// Statistics
func (s Stats) Stats(keyprefix string) (map[string]interface{}, error) {
	m := map[string]string{
		"keyPrefix": keyprefix,
	}
	return s.c.Request("stats/view/stats/", m)
}

// Gets all of the site based statistics, optionally filtered by a key prefix
func (s Stats) AllSitesStats(keyprefix string) (map[string]interface{}, error) {
	m := map[string]string{
		"keyPrefix": keyprefix,
	}
	return s.c.Request("stats/view/allSitesStats/", m)
}

// Gets all of the global statistics, optionally filtered by a key prefix
func (s Stats) SiteStats(site string, keyprefix string) (map[string]interface{}, error) {
	m := map[string]string{
		"site": site,
		"keyPrefix": keyprefix,
	}
	return s.c.Request("stats/view/siteStats/", m)
}

// Gets the Statsd service hostname
func (s Stats) OptionStatsdHost() (map[string]interface{}, error) {
	return s.c.Request("stats/view/optionStatsdHost/", nil)
}

// Gets the Statsd service port
func (s Stats) OptionStatsdPort() (map[string]interface{}, error) {
	return s.c.Request("stats/view/optionStatsdPort/", nil)
}

// Gets the prefix to be applied to all stats sent to the configured Statsd service
func (s Stats) OptionStatsdPrefix() (map[string]interface{}, error) {
	return s.c.Request("stats/view/optionStatsdPrefix/", nil)
}

// Returns 'true' if in memory statistics are enabled, otherwise returns 'false'
func (s Stats) OptionInMemoryEnabled() (map[string]interface{}, error) {
	return s.c.Request("stats/view/optionInMemoryEnabled/", nil)
}

// Returns 'true' if a Statsd server has been correctly configured, otherwise returns 'false'
func (s Stats) OptionStatsdEnabled() (map[string]interface{}, error) {
	return s.c.Request("stats/view/optionStatsdEnabled/", nil)
}

// Clears all of the statistics
func (s Stats) ClearStats(keyprefix string) (map[string]interface{}, error) {
	m := map[string]string{
		"keyPrefix": keyprefix,
	}
	return s.c.Request("stats/action/clearStats/", m)
}

// Sets the Statsd service hostname, supply an empty string to stop using a Statsd service
func (s Stats) SetOptionStatsdHost(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return s.c.Request("stats/action/setOptionStatsdHost/", m)
}

// Sets the prefix to be applied to all stats sent to the configured Statsd service
func (s Stats) SetOptionStatsdPrefix(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return s.c.Request("stats/action/setOptionStatsdPrefix/", m)
}

// Sets whether in memory statistics are enabled
func (s Stats) SetOptionInMemoryEnabled(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return s.c.Request("stats/action/setOptionInMemoryEnabled/", m)
}

// Sets the Statsd service port
func (s Stats) SetOptionStatsdPort(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return s.c.Request("stats/action/setOptionStatsdPort/", m)
}

