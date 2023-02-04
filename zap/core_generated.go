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

type Core struct {
	c *Client
}

// Gets the name of the hosts accessed through/by ZAP
func (c Core) Hosts() (map[string]interface{}, error) {
	return c.c.Request("core/view/hosts/", nil)
}

// Gets the sites accessed through/by ZAP (scheme and domain)
func (c Core) Sites() (map[string]interface{}, error) {
	return c.c.Request("core/view/sites/", nil)
}

// Gets the URLs accessed through/by ZAP, optionally filtering by (base) URL.
func (c Core) Urls(baseurl string) (map[string]interface{}, error) {
	m := map[string]string{
		"baseurl": baseurl,
	}
	return c.c.Request("core/view/urls/", m)
}

// Gets the child nodes underneath the specified URL in the Sites tree
func (c Core) ChildNodes(url string) (map[string]interface{}, error) {
	m := map[string]string{
		"url": url,
	}
	return c.c.Request("core/view/childNodes/", m)
}

// Gets the HTTP message with the given ID. Returns the ID, request/response headers and bodies, cookies, note, type, RTT, and timestamp.
func (c Core) Message(id string) (map[string]interface{}, error) {
	m := map[string]string{
		"id": id,
	}
	return c.c.Request("core/view/message/", m)
}

// Gets the HTTP messages sent by ZAP, request and response, optionally filtered by URL and paginated with 'start' position and 'count' of messages
func (c Core) Messages(baseurl string, start string, count string) (map[string]interface{}, error) {
	m := map[string]string{
		"baseurl": baseurl,
		"start":   start,
		"count":   count,
	}
	return c.c.Request("core/view/messages/", m)
}

// Gets the HTTP messages with the given IDs.
func (c Core) MessagesById(ids string) (map[string]interface{}, error) {
	m := map[string]string{
		"ids": ids,
	}
	return c.c.Request("core/view/messagesById/", m)
}

// Gets the number of messages, optionally filtering by URL
func (c Core) NumberOfMessages(baseurl string) (map[string]interface{}, error) {
	m := map[string]string{
		"baseurl": baseurl,
	}
	return c.c.Request("core/view/numberOfMessages/", m)
}

// Gets the mode
func (c Core) Mode() (map[string]interface{}, error) {
	return c.c.Request("core/view/mode/", nil)
}

// Gets ZAP version
func (c Core) Version() (map[string]interface{}, error) {
	return c.c.Request("core/view/version/", nil)
}

// Gets the regular expressions, applied to URLs, to exclude from the local proxies.
func (c Core) ExcludedFromProxy() (map[string]interface{}, error) {
	return c.c.Request("core/view/excludedFromProxy/", nil)
}

// Gets the location of the current session file
func (c Core) SessionLocation() (map[string]interface{}, error) {
	return c.c.Request("core/view/sessionLocation/", nil)
}

// Gets all the domains that are excluded from the outgoing proxy. For each domain the following are shown: the index, the value (domain), if enabled, and if specified as a regex.
func (c Core) ProxyChainExcludedDomains() (map[string]interface{}, error) {
	return c.c.Request("core/view/proxyChainExcludedDomains/", nil)
}

// Gets the path to ZAP's home directory.
func (c Core) ZapHomePath() (map[string]interface{}, error) {
	return c.c.Request("core/view/zapHomePath/", nil)
}

// Gets the maximum number of alert instances to include in a report.
func (c Core) OptionMaximumAlertInstances() (map[string]interface{}, error) {
	return c.c.Request("core/view/optionMaximumAlertInstances/", nil)
}

// Gets whether or not related alerts will be merged in any reports generated.
func (c Core) OptionMergeRelatedAlerts() (map[string]interface{}, error) {
	return c.c.Request("core/view/optionMergeRelatedAlerts/", nil)
}

// Gets the path to the file with alert overrides.
func (c Core) OptionAlertOverridesFilePath() (map[string]interface{}, error) {
	return c.c.Request("core/view/optionAlertOverridesFilePath/", nil)
}

func (c Core) HomeDirectory() (map[string]interface{}, error) {
	return c.c.Request("core/view/homeDirectory/", nil)
}

// Use view proxyChainExcludedDomains instead.
func (c Core) OptionProxyChainSkipName() (map[string]interface{}, error) {
	return c.c.Request("core/view/optionProxyChainSkipName/", nil)
}

// Use view proxyChainExcludedDomains instead.
func (c Core) OptionProxyExcludedDomains() (map[string]interface{}, error) {
	return c.c.Request("core/view/optionProxyExcludedDomains/", nil)
}

// Use view proxyChainExcludedDomains instead.
func (c Core) OptionProxyExcludedDomainsEnabled() (map[string]interface{}, error) {
	return c.c.Request("core/view/optionProxyExcludedDomainsEnabled/", nil)
}

// Gets the alert with the given ID, the corresponding HTTP message can be obtained with the 'messageId' field and 'message' API method
func (c Core) Alert(id string) (map[string]interface{}, error) {
	m := map[string]string{
		"id": id,
	}
	return c.c.Request("core/view/alert/", m)
}

// Gets the alerts raised by ZAP, optionally filtering by URL or riskId, and paginating with 'start' position and 'count' of alerts
func (c Core) Alerts(baseurl string, start string, count string, riskid string) (map[string]interface{}, error) {
	m := map[string]string{
		"baseurl": baseurl,
		"start":   start,
		"count":   count,
		"riskId":  riskid,
	}
	return c.c.Request("core/view/alerts/", m)
}

// Gets number of alerts grouped by each risk level, optionally filtering by URL
func (c Core) AlertsSummary(baseurl string) (map[string]interface{}, error) {
	m := map[string]string{
		"baseurl": baseurl,
	}
	return c.c.Request("core/view/alertsSummary/", m)
}

// Gets the number of alerts, optionally filtering by URL or riskId
func (c Core) NumberOfAlerts(baseurl string, riskid string) (map[string]interface{}, error) {
	m := map[string]string{
		"baseurl": baseurl,
		"riskId":  riskid,
	}
	return c.c.Request("core/view/numberOfAlerts/", m)
}

// Gets the user agent that ZAP should use when creating HTTP messages (for example, spider messages or CONNECT requests to outgoing proxy).
func (c Core) OptionDefaultUserAgent() (map[string]interface{}, error) {
	return c.c.Request("core/view/optionDefaultUserAgent/", nil)
}

// Gets the TTL (in seconds) of successful DNS queries.
func (c Core) OptionDnsTtlSuccessfulQueries() (map[string]interface{}, error) {
	return c.c.Request("core/view/optionDnsTtlSuccessfulQueries/", nil)
}

func (c Core) OptionHttpState() (map[string]interface{}, error) {
	return c.c.Request("core/view/optionHttpState/", nil)
}

func (c Core) OptionHttpStateEnabled() (map[string]interface{}, error) {
	return c.c.Request("core/view/optionHttpStateEnabled/", nil)
}

func (c Core) OptionProxyChainName() (map[string]interface{}, error) {
	return c.c.Request("core/view/optionProxyChainName/", nil)
}

func (c Core) OptionProxyChainPassword() (map[string]interface{}, error) {
	return c.c.Request("core/view/optionProxyChainPassword/", nil)
}

func (c Core) OptionProxyChainPort() (map[string]interface{}, error) {
	return c.c.Request("core/view/optionProxyChainPort/", nil)
}

func (c Core) OptionProxyChainPrompt() (map[string]interface{}, error) {
	return c.c.Request("core/view/optionProxyChainPrompt/", nil)
}

func (c Core) OptionProxyChainRealm() (map[string]interface{}, error) {
	return c.c.Request("core/view/optionProxyChainRealm/", nil)
}

func (c Core) OptionProxyChainUserName() (map[string]interface{}, error) {
	return c.c.Request("core/view/optionProxyChainUserName/", nil)
}

func (c Core) OptionSingleCookieRequestHeader() (map[string]interface{}, error) {
	return c.c.Request("core/view/optionSingleCookieRequestHeader/", nil)
}

// Gets the connection time out (in seconds).
func (c Core) OptionTimeoutInSecs() (map[string]interface{}, error) {
	return c.c.Request("core/view/optionTimeoutInSecs/", nil)
}

func (c Core) OptionUseProxyChain() (map[string]interface{}, error) {
	return c.c.Request("core/view/optionUseProxyChain/", nil)
}

func (c Core) OptionUseProxyChainAuth() (map[string]interface{}, error) {
	return c.c.Request("core/view/optionUseProxyChainAuth/", nil)
}

// Gets whether or not the SOCKS proxy should be used.
func (c Core) OptionUseSocksProxy() (map[string]interface{}, error) {
	return c.c.Request("core/view/optionUseSocksProxy/", nil)
}

// Convenient and simple action to access a URL, optionally following redirections. Returns the request sent and response received and followed redirections, if any. Other actions are available which offer more control on what is sent, like, 'sendRequest' or 'sendHarRequest'.
func (c Core) AccessUrl(url string, followredirects string) (map[string]interface{}, error) {
	m := map[string]string{
		"url":             url,
		"followRedirects": followredirects,
	}
	return c.c.Request("core/action/accessUrl/", m)
}

// Shuts down ZAP
func (c Core) Shutdown() (map[string]interface{}, error) {
	return c.c.Request("core/action/shutdown/", nil)
}

// Creates a new session, optionally overwriting existing files. If a relative path is specified it will be resolved against the "session" directory in ZAP "home" dir.
func (c Core) NewSession(name string, overwrite string) (map[string]interface{}, error) {
	m := map[string]string{
		"name":      name,
		"overwrite": overwrite,
	}
	return c.c.Request("core/action/newSession/", m)
}

// Loads the session with the given name. If a relative path is specified it will be resolved against the "session" directory in ZAP "home" dir.
func (c Core) LoadSession(name string) (map[string]interface{}, error) {
	m := map[string]string{
		"name": name,
	}
	return c.c.Request("core/action/loadSession/", m)
}

// Saves the session.
func (c Core) SaveSession(name string, overwrite string) (map[string]interface{}, error) {
	m := map[string]string{
		"name":      name,
		"overwrite": overwrite,
	}
	return c.c.Request("core/action/saveSession/", m)
}

// Snapshots the session, optionally with the given name, and overwriting existing files. If no name is specified the name of the current session with a timestamp appended is used. If a relative path is specified it will be resolved against the "session" directory in ZAP "home" dir.
func (c Core) SnapshotSession(name string, overwrite string) (map[string]interface{}, error) {
	m := map[string]string{
		"name":      name,
		"overwrite": overwrite,
	}
	return c.c.Request("core/action/snapshotSession/", m)
}

// Clears the regexes of URLs excluded from the local proxies.
func (c Core) ClearExcludedFromProxy() (map[string]interface{}, error) {
	return c.c.Request("core/action/clearExcludedFromProxy/", nil)
}

// Adds a regex of URLs that should be excluded from the local proxies.
func (c Core) ExcludeFromProxy(regex string) (map[string]interface{}, error) {
	m := map[string]string{
		"regex": regex,
	}
	return c.c.Request("core/action/excludeFromProxy/", m)
}

func (c Core) SetHomeDirectory(dir string) (map[string]interface{}, error) {
	m := map[string]string{
		"dir": dir,
	}
	return c.c.Request("core/action/setHomeDirectory/", m)
}

// Sets the mode, which may be one of [safe, protect, standard, attack]
func (c Core) SetMode(mode string) (map[string]interface{}, error) {
	m := map[string]string{
		"mode": mode,
	}
	return c.c.Request("core/action/setMode/", m)
}

// Generates a new Root CA certificate for the local proxies.
func (c Core) GenerateRootCA() (map[string]interface{}, error) {
	return c.c.Request("core/action/generateRootCA/", nil)
}

// Sends the HTTP request, optionally following redirections. Returns the request sent and response received and followed redirections, if any. The Mode is enforced when sending the request (and following redirections), custom manual requests are not allowed in 'Safe' mode nor in 'Protected' mode if out of scope.
func (c Core) SendRequest(request string, followredirects string) (map[string]interface{}, error) {
	m := map[string]string{
		"request":         request,
		"followRedirects": followredirects,
	}
	return c.c.Request("core/action/sendRequest/", m)
}

func (c Core) RunGarbageCollection() (map[string]interface{}, error) {
	return c.c.Request("core/action/runGarbageCollection/", nil)
}

// Deletes the site node found in the Sites Tree on the basis of the URL, HTTP method, and post data (if applicable and specified).
func (c Core) DeleteSiteNode(url string, method string, postdata string) (map[string]interface{}, error) {
	m := map[string]string{
		"url":      url,
		"method":   method,
		"postData": postdata,
	}
	return c.c.Request("core/action/deleteSiteNode/", m)
}

// Adds a domain to be excluded from the outgoing proxy, using the specified value. Optionally sets if the new entry is enabled (default, true) and whether or not the new value is specified as a regex (default, false).
func (c Core) AddProxyChainExcludedDomain(value string, isregex string, isenabled string) (map[string]interface{}, error) {
	m := map[string]string{
		"value":     value,
		"isRegex":   isregex,
		"isEnabled": isenabled,
	}
	return c.c.Request("core/action/addProxyChainExcludedDomain/", m)
}

// Modifies a domain excluded from the outgoing proxy. Allows to modify the value, if enabled or if a regex. The domain is selected with its index, which can be obtained with the view proxyChainExcludedDomains.
func (c Core) ModifyProxyChainExcludedDomain(idx string, value string, isregex string, isenabled string) (map[string]interface{}, error) {
	m := map[string]string{
		"idx":       idx,
		"value":     value,
		"isRegex":   isregex,
		"isEnabled": isenabled,
	}
	return c.c.Request("core/action/modifyProxyChainExcludedDomain/", m)
}

// Removes a domain excluded from the outgoing proxy, with the given index. The index can be obtained with the view proxyChainExcludedDomains.
func (c Core) RemoveProxyChainExcludedDomain(idx string) (map[string]interface{}, error) {
	m := map[string]string{
		"idx": idx,
	}
	return c.c.Request("core/action/removeProxyChainExcludedDomain/", m)
}

// Enables all domains excluded from the outgoing proxy.
func (c Core) EnableAllProxyChainExcludedDomains() (map[string]interface{}, error) {
	return c.c.Request("core/action/enableAllProxyChainExcludedDomains/", nil)
}

// Disables all domains excluded from the outgoing proxy.
func (c Core) DisableAllProxyChainExcludedDomains() (map[string]interface{}, error) {
	return c.c.Request("core/action/disableAllProxyChainExcludedDomains/", nil)
}

// Sets the maximum number of alert instances to include in a report. A value of zero is treated as unlimited.
func (c Core) SetOptionMaximumAlertInstances(numberofinstances string) (map[string]interface{}, error) {
	m := map[string]string{
		"numberOfInstances": numberofinstances,
	}
	return c.c.Request("core/action/setOptionMaximumAlertInstances/", m)
}

// Sets whether or not related alerts will be merged in any reports generated.
func (c Core) SetOptionMergeRelatedAlerts(enabled string) (map[string]interface{}, error) {
	m := map[string]string{
		"enabled": enabled,
	}
	return c.c.Request("core/action/setOptionMergeRelatedAlerts/", m)
}

// Sets (or clears, if empty) the path to the file with alert overrides.
func (c Core) SetOptionAlertOverridesFilePath(filepath string) (map[string]interface{}, error) {
	m := map[string]string{
		"filePath": filepath,
	}
	return c.c.Request("core/action/setOptionAlertOverridesFilePath/", m)
}

// Enables use of a PKCS12 client certificate for the certificate with the given file system path, password, and optional index.
func (c Core) EnablePKCS12ClientCertificate(filepath string, password string, index string) (map[string]interface{}, error) {
	m := map[string]string{
		"filePath": filepath,
		"password": password,
		"index":    index,
	}
	return c.c.Request("core/action/enablePKCS12ClientCertificate/", m)
}

// Disables the option for use of client certificates.
func (c Core) DisableClientCertificate() (map[string]interface{}, error) {
	return c.c.Request("core/action/disableClientCertificate/", nil)
}

// Deletes all alerts of the current session.
func (c Core) DeleteAllAlerts() (map[string]interface{}, error) {
	return c.c.Request("core/action/deleteAllAlerts/", nil)
}

// Deletes the alert with the given ID.
func (c Core) DeleteAlert(id string) (map[string]interface{}, error) {
	m := map[string]string{
		"id": id,
	}
	return c.c.Request("core/action/deleteAlert/", m)
}

// Sets the user agent that ZAP should use when creating HTTP messages (for example, spider messages or CONNECT requests to outgoing proxy).
func (c Core) SetOptionDefaultUserAgent(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return c.c.Request("core/action/setOptionDefaultUserAgent/", m)
}

// Sets the TTL (in seconds) of successful DNS queries (applies after ZAP restart).
func (c Core) SetOptionDnsTtlSuccessfulQueries(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return c.c.Request("core/action/setOptionDnsTtlSuccessfulQueries/", m)
}

func (c Core) SetOptionHttpStateEnabled(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return c.c.Request("core/action/setOptionHttpStateEnabled/", m)
}

func (c Core) SetOptionProxyChainName(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return c.c.Request("core/action/setOptionProxyChainName/", m)
}

func (c Core) SetOptionProxyChainPassword(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return c.c.Request("core/action/setOptionProxyChainPassword/", m)
}

func (c Core) SetOptionProxyChainPort(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return c.c.Request("core/action/setOptionProxyChainPort/", m)
}

func (c Core) SetOptionProxyChainPrompt(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return c.c.Request("core/action/setOptionProxyChainPrompt/", m)
}

func (c Core) SetOptionProxyChainRealm(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return c.c.Request("core/action/setOptionProxyChainRealm/", m)
}

// Use actions [add|modify|remove]ProxyChainExcludedDomain instead.
func (c Core) SetOptionProxyChainSkipName(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return c.c.Request("core/action/setOptionProxyChainSkipName/", m)
}

func (c Core) SetOptionProxyChainUserName(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return c.c.Request("core/action/setOptionProxyChainUserName/", m)
}

func (c Core) SetOptionSingleCookieRequestHeader(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return c.c.Request("core/action/setOptionSingleCookieRequestHeader/", m)
}

// Sets the connection time out (in seconds).
func (c Core) SetOptionTimeoutInSecs(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return c.c.Request("core/action/setOptionTimeoutInSecs/", m)
}

// Sets whether or not the outgoing proxy should be used. The address/hostname of the outgoing proxy must be set to enable this option.
func (c Core) SetOptionUseProxyChain(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return c.c.Request("core/action/setOptionUseProxyChain/", m)
}

func (c Core) SetOptionUseProxyChainAuth(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return c.c.Request("core/action/setOptionUseProxyChainAuth/", m)
}

// Sets whether or not the SOCKS proxy should be used.
func (c Core) SetOptionUseSocksProxy(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return c.c.Request("core/action/setOptionUseSocksProxy/", m)
}

func (c Core) Proxypac() ([]byte, error) {
	return c.c.RequestOther("core/other/proxy.pac/", nil)
}

// Gets the Root CA certificate used by the local proxies.
func (c Core) Rootcert() ([]byte, error) {
	return c.c.RequestOther("core/other/rootcert/", nil)
}

func (c Core) Setproxy(proxy string) ([]byte, error) {
	m := map[string]string{
		"proxy": proxy,
	}
	return c.c.RequestOther("core/other/setproxy/", m)
}

// Generates a report in XML format
func (c Core) Xmlreport() ([]byte, error) {
	return c.c.RequestOther("core/other/xmlreport/", nil)
}

// Generates a report in HTML format
func (c Core) Htmlreport() ([]byte, error) {
	return c.c.RequestOther("core/other/htmlreport/", nil)
}

// Generates a report in JSON format
func (c Core) Jsonreport() ([]byte, error) {
	return c.c.RequestOther("core/other/jsonreport/", nil)
}

// Generates a report in Markdown format
func (c Core) Mdreport() ([]byte, error) {
	return c.c.RequestOther("core/other/mdreport/", nil)
}

// Gets the message with the given ID in HAR format
func (c Core) MessageHar(id string) ([]byte, error) {
	m := map[string]string{
		"id": id,
	}
	return c.c.RequestOther("core/other/messageHar/", m)
}

// Gets the HTTP messages sent through/by ZAP, in HAR format, optionally filtered by URL and paginated with 'start' position and 'count' of messages
func (c Core) MessagesHar(baseurl string, start string, count string) ([]byte, error) {
	m := map[string]string{
		"baseurl": baseurl,
		"start":   start,
		"count":   count,
	}
	return c.c.RequestOther("core/other/messagesHar/", m)
}

// Gets the HTTP messages with the given IDs, in HAR format.
func (c Core) MessagesHarById(ids string) ([]byte, error) {
	m := map[string]string{
		"ids": ids,
	}
	return c.c.RequestOther("core/other/messagesHarById/", m)
}

// Sends the first HAR request entry, optionally following redirections. Returns, in HAR format, the request sent and response received and followed redirections, if any. The Mode is enforced when sending the request (and following redirections), custom manual requests are not allowed in 'Safe' mode nor in 'Protected' mode if out of scope.
func (c Core) SendHarRequest(request string, followredirects string) ([]byte, error) {
	m := map[string]string{
		"request":         request,
		"followRedirects": followredirects,
	}
	return c.c.RequestOther("core/other/sendHarRequest/", m)
}
