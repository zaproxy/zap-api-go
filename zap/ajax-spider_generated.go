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

type AjaxSpider struct {
	c *Client
}

// Gets the allowed resources. The allowed resources are always fetched even if out of scope, allowing to include necessary resources (e.g. scripts) from 3rd-parties.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) AllowedResources() (map[string]interface{}, error) {
	return a.c.Request("ajaxSpider/view/allowedResources/", nil)
}

// Gets the current status of the crawler. Actual values are Stopped and Running.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) Status() (map[string]interface{}, error) {
	return a.c.Request("ajaxSpider/view/status/", nil)
}

// Gets the current results of the crawler.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) Results(start string, count string) (map[string]interface{}, error) {
	m := map[string]string{
		"start": start,
		"count": count,
	}
	return a.c.Request("ajaxSpider/view/results/", m)
}

// Gets the number of resources found.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) NumberOfResults() (map[string]interface{}, error) {
	return a.c.Request("ajaxSpider/view/numberOfResults/", nil)
}

// Gets the full crawled content detected by the AJAX Spider. Returns a set of values based on 'inScope' URLs, 'outOfScope' URLs, and 'errors' encountered during the last/current run of the AJAX Spider.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) FullResults() (map[string]interface{}, error) {
	return a.c.Request("ajaxSpider/view/fullResults/", nil)
}

// Gets the configured browser to use for crawling.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) OptionBrowserId() (map[string]interface{}, error) {
	return a.c.Request("ajaxSpider/view/optionBrowserId/", nil)
}

// Gets the time to wait after an event (in milliseconds). For example: the wait delay after the cursor hovers over an element, in order for a menu to display, etc.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) OptionEventWait() (map[string]interface{}, error) {
	return a.c.Request("ajaxSpider/view/optionEventWait/", nil)
}

// Gets the configured value for the max crawl depth.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) OptionMaxCrawlDepth() (map[string]interface{}, error) {
	return a.c.Request("ajaxSpider/view/optionMaxCrawlDepth/", nil)
}

// Gets the configured value for the maximum crawl states allowed.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) OptionMaxCrawlStates() (map[string]interface{}, error) {
	return a.c.Request("ajaxSpider/view/optionMaxCrawlStates/", nil)
}

// Gets the configured max duration of the crawl, the value is in minutes.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) OptionMaxDuration() (map[string]interface{}, error) {
	return a.c.Request("ajaxSpider/view/optionMaxDuration/", nil)
}

// Gets the configured number of browsers to be used.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) OptionNumberOfBrowsers() (map[string]interface{}, error) {
	return a.c.Request("ajaxSpider/view/optionNumberOfBrowsers/", nil)
}

// Gets the configured time to wait after reloading the page, this value is in milliseconds.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) OptionReloadWait() (map[string]interface{}, error) {
	return a.c.Request("ajaxSpider/view/optionReloadWait/", nil)
}

// Gets the configured value for 'Click Default Elements Only', HTML elements such as 'a', 'button', 'input', all associated with some action or links on the page.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) OptionClickDefaultElems() (map[string]interface{}, error) {
	return a.c.Request("ajaxSpider/view/optionClickDefaultElems/", nil)
}

// Gets the value configured for the AJAX Spider to know if it should click on the elements only once.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) OptionClickElemsOnce() (map[string]interface{}, error) {
	return a.c.Request("ajaxSpider/view/optionClickElemsOnce/", nil)
}

// Gets if the AJAX Spider will use random values in form fields when crawling, if set to true.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) OptionRandomInputs() (map[string]interface{}, error) {
	return a.c.Request("ajaxSpider/view/optionRandomInputs/", nil)
}

// Runs the AJAX Spider against a given target.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) Scan(url string, inscope string, contextname string, subtreeonly string) (map[string]interface{}, error) {
	m := map[string]string{
		"url": url,
		"inScope": inscope,
		"contextName": contextname,
		"subtreeOnly": subtreeonly,
	}
	return a.c.Request("ajaxSpider/action/scan/", m)
}

// Runs the AJAX Spider from the perspective of a User of the web application.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) ScanAsUser(contextname string, username string, url string, subtreeonly string) (map[string]interface{}, error) {
	m := map[string]string{
		"contextName": contextname,
		"userName": username,
		"url": url,
		"subtreeOnly": subtreeonly,
	}
	return a.c.Request("ajaxSpider/action/scanAsUser/", m)
}

// Stops the AJAX Spider.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) Stop() (map[string]interface{}, error) {
	return a.c.Request("ajaxSpider/action/stop/", nil)
}

// Adds an allowed resource.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) AddAllowedResource(regex string, enabled string) (map[string]interface{}, error) {
	m := map[string]string{
		"regex": regex,
		"enabled": enabled,
	}
	return a.c.Request("ajaxSpider/action/addAllowedResource/", m)
}

// Removes an allowed resource.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) RemoveAllowedResource(regex string) (map[string]interface{}, error) {
	m := map[string]string{
		"regex": regex,
	}
	return a.c.Request("ajaxSpider/action/removeAllowedResource/", m)
}

// Sets whether or not an allowed resource is enabled.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) SetEnabledAllowedResource(regex string, enabled string) (map[string]interface{}, error) {
	m := map[string]string{
		"regex": regex,
		"enabled": enabled,
	}
	return a.c.Request("ajaxSpider/action/setEnabledAllowedResource/", m)
}

// Sets the configuration of the AJAX Spider to use one of the supported browsers.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) SetOptionBrowserId(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return a.c.Request("ajaxSpider/action/setOptionBrowserId/", m)
}

// Sets whether or not the the AJAX Spider will only click on the default HTML elements.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) SetOptionClickDefaultElems(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return a.c.Request("ajaxSpider/action/setOptionClickDefaultElems/", m)
}

// When enabled, the crawler attempts to interact with each element (e.g., by clicking) only once.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) SetOptionClickElemsOnce(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return a.c.Request("ajaxSpider/action/setOptionClickElemsOnce/", m)
}

// Sets the time to wait after an event (in milliseconds). For example: the wait delay after the cursor hovers over an element, in order for a menu to display, etc.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) SetOptionEventWait(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return a.c.Request("ajaxSpider/action/setOptionEventWait/", m)
}

// Sets the maximum depth that the crawler can reach.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) SetOptionMaxCrawlDepth(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return a.c.Request("ajaxSpider/action/setOptionMaxCrawlDepth/", m)
}

// Sets the maximum number of states that the crawler should crawl.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) SetOptionMaxCrawlStates(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return a.c.Request("ajaxSpider/action/setOptionMaxCrawlStates/", m)
}

// The maximum time that the crawler is allowed to run.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) SetOptionMaxDuration(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return a.c.Request("ajaxSpider/action/setOptionMaxDuration/", m)
}

// Sets the number of windows to be used by AJAX Spider.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) SetOptionNumberOfBrowsers(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return a.c.Request("ajaxSpider/action/setOptionNumberOfBrowsers/", m)
}

// When enabled, inserts random values into form fields.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) SetOptionRandomInputs(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return a.c.Request("ajaxSpider/action/setOptionRandomInputs/", m)
}

// Sets the time to wait after the page is loaded before interacting with it.
//
// This component is optional and therefore the API will only work if it is installed
func (a AjaxSpider) SetOptionReloadWait(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return a.c.Request("ajaxSpider/action/setOptionReloadWait/", m)
}

