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

type Ascan struct {
	c *Client
}

// 
func (a Ascan) Status(scanid string) (map[string]interface{}, error) {
	m := map[string]string{
		"scanId": scanid,
	}
	return a.c.Request("ascan/view/status/", m)
}

// 
func (a Ascan) ScanProgress(scanid string) (map[string]interface{}, error) {
	m := map[string]string{
		"scanId": scanid,
	}
	return a.c.Request("ascan/view/scanProgress/", m)
}

// Gets the IDs of the messages sent during the scan with the given ID. A message can be obtained with 'message' core view.
func (a Ascan) MessagesIds(scanid string) (map[string]interface{}, error) {
	m := map[string]string{
		"scanId": scanid,
	}
	return a.c.Request("ascan/view/messagesIds/", m)
}

// Gets the IDs of the alerts raised during the scan with the given ID. An alert can be obtained with 'alert' core view.
func (a Ascan) AlertsIds(scanid string) (map[string]interface{}, error) {
	m := map[string]string{
		"scanId": scanid,
	}
	return a.c.Request("ascan/view/alertsIds/", m)
}

// 
func (a Ascan) Scans() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/scans/", nil)
}

// 
func (a Ascan) ScanPolicyNames() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/scanPolicyNames/", nil)
}

// Gets the regexes of URLs excluded from the active scans.
func (a Ascan) ExcludedFromScan() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/excludedFromScan/", nil)
}

// Gets the scan rules, optionally, of the given scan policy or scanner policy/category ID.
func (a Ascan) Scanners(scanpolicyname string, policyid string) (map[string]interface{}, error) {
	m := map[string]string{
		"scanPolicyName": scanpolicyname,
		"policyId": policyid,
	}
	return a.c.Request("ascan/view/scanners/", m)
}

// 
func (a Ascan) Policies(scanpolicyname string, policyid string) (map[string]interface{}, error) {
	m := map[string]string{
		"scanPolicyName": scanpolicyname,
		"policyId": policyid,
	}
	return a.c.Request("ascan/view/policies/", m)
}

// 
func (a Ascan) AttackModeQueue() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/attackModeQueue/", nil)
}

// Gets all the parameters that are excluded. For each parameter the following are shown: the name, the URL, and the parameter type.
func (a Ascan) ExcludedParams() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/excludedParams/", nil)
}

// Use view excludedParams instead.
func (a Ascan) OptionExcludedParamList() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/optionExcludedParamList/", nil)
}

// Gets all the types of excluded parameters. For each type the following are shown: the ID and the name.
func (a Ascan) ExcludedParamTypes() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/excludedParamTypes/", nil)
}

// 
func (a Ascan) OptionAttackPolicy() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/optionAttackPolicy/", nil)
}

// 
func (a Ascan) OptionDefaultPolicy() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/optionDefaultPolicy/", nil)
}

// 
func (a Ascan) OptionDelayInMs() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/optionDelayInMs/", nil)
}

// 
func (a Ascan) OptionHandleAntiCSRFTokens() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/optionHandleAntiCSRFTokens/", nil)
}

// 
func (a Ascan) OptionHostPerScan() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/optionHostPerScan/", nil)
}

// 
func (a Ascan) OptionMaxChartTimeInMins() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/optionMaxChartTimeInMins/", nil)
}

// 
func (a Ascan) OptionMaxResultsToList() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/optionMaxResultsToList/", nil)
}

// 
func (a Ascan) OptionMaxRuleDurationInMins() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/optionMaxRuleDurationInMins/", nil)
}

// 
func (a Ascan) OptionMaxScanDurationInMins() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/optionMaxScanDurationInMins/", nil)
}

// 
func (a Ascan) OptionMaxScansInUI() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/optionMaxScansInUI/", nil)
}

// 
func (a Ascan) OptionTargetParamsEnabledRPC() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/optionTargetParamsEnabledRPC/", nil)
}

// 
func (a Ascan) OptionTargetParamsInjectable() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/optionTargetParamsInjectable/", nil)
}

// 
func (a Ascan) OptionThreadPerHost() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/optionThreadPerHost/", nil)
}

// Tells whether or not the active scanner should add a query parameter to GET request that don't have parameters to start with.
func (a Ascan) OptionAddQueryParam() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/optionAddQueryParam/", nil)
}

// 
func (a Ascan) OptionAllowAttackOnStart() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/optionAllowAttackOnStart/", nil)
}

// Tells whether or not the active scanner should inject the HTTP request header X-ZAP-Scan-ID, with the ID of the scanner that's sending the requests.
func (a Ascan) OptionInjectPluginIdInHeader() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/optionInjectPluginIdInHeader/", nil)
}

// 
func (a Ascan) OptionPromptInAttackMode() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/optionPromptInAttackMode/", nil)
}

// 
func (a Ascan) OptionPromptToClearFinishedScans() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/optionPromptToClearFinishedScans/", nil)
}

// 
func (a Ascan) OptionRescanInAttackMode() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/optionRescanInAttackMode/", nil)
}

// Tells whether or not the HTTP Headers of all requests should be scanned. Not just requests that send parameters, through the query or request body.
func (a Ascan) OptionScanHeadersAllRequests() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/optionScanHeadersAllRequests/", nil)
}

// Tells whether or not the active scanner should scan null JSON values.
func (a Ascan) OptionScanNullJsonValues() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/optionScanNullJsonValues/", nil)
}

// 
func (a Ascan) OptionShowAdvancedDialog() (map[string]interface{}, error) {
	return a.c.Request("ascan/view/optionShowAdvancedDialog/", nil)
}

// Runs the active scanner against the given URL or Context. Optionally, the 'recurse' parameter can be used to scan URLs under the given URL, the parameter 'inScopeOnly' can be used to constrain the scan to URLs that are in scope (ignored if a Context is specified), the parameter 'scanPolicyName' allows to specify the scan policy (if none is given it uses the default scan policy), the parameters 'method' and 'postData' allow to select a given request in conjunction with the given URL.
func (a Ascan) Scan(url string, recurse string, inscopeonly string, scanpolicyname string, method string, postdata string, contextid string) (map[string]interface{}, error) {
	m := map[string]string{
		"url": url,
		"recurse": recurse,
		"inScopeOnly": inscopeonly,
		"scanPolicyName": scanpolicyname,
		"method": method,
		"postData": postdata,
		"contextId": contextid,
	}
	return a.c.Request("ascan/action/scan/", m)
}

// Active Scans from the perspective of a User, obtained using the given Context ID and User ID. See 'scan' action for more details.
func (a Ascan) ScanAsUser(url string, contextid string, userid string, recurse string, scanpolicyname string, method string, postdata string) (map[string]interface{}, error) {
	m := map[string]string{
		"url": url,
		"contextId": contextid,
		"userId": userid,
		"recurse": recurse,
		"scanPolicyName": scanpolicyname,
		"method": method,
		"postData": postdata,
	}
	return a.c.Request("ascan/action/scanAsUser/", m)
}

// 
func (a Ascan) Pause(scanid string) (map[string]interface{}, error) {
	m := map[string]string{
		"scanId": scanid,
	}
	return a.c.Request("ascan/action/pause/", m)
}

// 
func (a Ascan) Resume(scanid string) (map[string]interface{}, error) {
	m := map[string]string{
		"scanId": scanid,
	}
	return a.c.Request("ascan/action/resume/", m)
}

// 
func (a Ascan) Stop(scanid string) (map[string]interface{}, error) {
	m := map[string]string{
		"scanId": scanid,
	}
	return a.c.Request("ascan/action/stop/", m)
}

// 
func (a Ascan) RemoveScan(scanid string) (map[string]interface{}, error) {
	m := map[string]string{
		"scanId": scanid,
	}
	return a.c.Request("ascan/action/removeScan/", m)
}

// 
func (a Ascan) PauseAllScans() (map[string]interface{}, error) {
	return a.c.Request("ascan/action/pauseAllScans/", nil)
}

// 
func (a Ascan) ResumeAllScans() (map[string]interface{}, error) {
	return a.c.Request("ascan/action/resumeAllScans/", nil)
}

// 
func (a Ascan) StopAllScans() (map[string]interface{}, error) {
	return a.c.Request("ascan/action/stopAllScans/", nil)
}

// 
func (a Ascan) RemoveAllScans() (map[string]interface{}, error) {
	return a.c.Request("ascan/action/removeAllScans/", nil)
}

// Clears the regexes of URLs excluded from the active scans.
func (a Ascan) ClearExcludedFromScan() (map[string]interface{}, error) {
	return a.c.Request("ascan/action/clearExcludedFromScan/", nil)
}

// Adds a regex of URLs that should be excluded from the active scans.
func (a Ascan) ExcludeFromScan(regex string) (map[string]interface{}, error) {
	m := map[string]string{
		"regex": regex,
	}
	return a.c.Request("ascan/action/excludeFromScan/", m)
}

// Enables all scan rules of the scan policy with the given name, or the default if none given.
func (a Ascan) EnableAllScanners(scanpolicyname string) (map[string]interface{}, error) {
	m := map[string]string{
		"scanPolicyName": scanpolicyname,
	}
	return a.c.Request("ascan/action/enableAllScanners/", m)
}

// Disables all scan rules of the scan policy with the given name, or the default if none given.
func (a Ascan) DisableAllScanners(scanpolicyname string) (map[string]interface{}, error) {
	m := map[string]string{
		"scanPolicyName": scanpolicyname,
	}
	return a.c.Request("ascan/action/disableAllScanners/", m)
}

// Enables the scan rules with the given IDs (comma separated list of IDs) of the scan policy with the given name, or the default if none given.
func (a Ascan) EnableScanners(ids string, scanpolicyname string) (map[string]interface{}, error) {
	m := map[string]string{
		"ids": ids,
		"scanPolicyName": scanpolicyname,
	}
	return a.c.Request("ascan/action/enableScanners/", m)
}

// Disables the scan rules with the given IDs (comma separated list of IDs) of the scan policy with the given name, or the default if none given.
func (a Ascan) DisableScanners(ids string, scanpolicyname string) (map[string]interface{}, error) {
	m := map[string]string{
		"ids": ids,
		"scanPolicyName": scanpolicyname,
	}
	return a.c.Request("ascan/action/disableScanners/", m)
}

// 
func (a Ascan) SetEnabledPolicies(ids string, scanpolicyname string) (map[string]interface{}, error) {
	m := map[string]string{
		"ids": ids,
		"scanPolicyName": scanpolicyname,
	}
	return a.c.Request("ascan/action/setEnabledPolicies/", m)
}

// 
func (a Ascan) SetPolicyAttackStrength(id string, attackstrength string, scanpolicyname string) (map[string]interface{}, error) {
	m := map[string]string{
		"id": id,
		"attackStrength": attackstrength,
		"scanPolicyName": scanpolicyname,
	}
	return a.c.Request("ascan/action/setPolicyAttackStrength/", m)
}

// 
func (a Ascan) SetPolicyAlertThreshold(id string, alertthreshold string, scanpolicyname string) (map[string]interface{}, error) {
	m := map[string]string{
		"id": id,
		"alertThreshold": alertthreshold,
		"scanPolicyName": scanpolicyname,
	}
	return a.c.Request("ascan/action/setPolicyAlertThreshold/", m)
}

// 
func (a Ascan) SetScannerAttackStrength(id string, attackstrength string, scanpolicyname string) (map[string]interface{}, error) {
	m := map[string]string{
		"id": id,
		"attackStrength": attackstrength,
		"scanPolicyName": scanpolicyname,
	}
	return a.c.Request("ascan/action/setScannerAttackStrength/", m)
}

// 
func (a Ascan) SetScannerAlertThreshold(id string, alertthreshold string, scanpolicyname string) (map[string]interface{}, error) {
	m := map[string]string{
		"id": id,
		"alertThreshold": alertthreshold,
		"scanPolicyName": scanpolicyname,
	}
	return a.c.Request("ascan/action/setScannerAlertThreshold/", m)
}

// 
func (a Ascan) AddScanPolicy(scanpolicyname string, alertthreshold string, attackstrength string) (map[string]interface{}, error) {
	m := map[string]string{
		"scanPolicyName": scanpolicyname,
		"alertThreshold": alertthreshold,
		"attackStrength": attackstrength,
	}
	return a.c.Request("ascan/action/addScanPolicy/", m)
}

// 
func (a Ascan) RemoveScanPolicy(scanpolicyname string) (map[string]interface{}, error) {
	m := map[string]string{
		"scanPolicyName": scanpolicyname,
	}
	return a.c.Request("ascan/action/removeScanPolicy/", m)
}

// 
func (a Ascan) UpdateScanPolicy(scanpolicyname string, alertthreshold string, attackstrength string) (map[string]interface{}, error) {
	m := map[string]string{
		"scanPolicyName": scanpolicyname,
		"alertThreshold": alertthreshold,
		"attackStrength": attackstrength,
	}
	return a.c.Request("ascan/action/updateScanPolicy/", m)
}

// Imports a Scan Policy using the given file system path.
func (a Ascan) ImportScanPolicy(path string) (map[string]interface{}, error) {
	m := map[string]string{
		"path": path,
	}
	return a.c.Request("ascan/action/importScanPolicy/", m)
}

// Adds a new parameter excluded from the scan, using the specified name. Optionally sets if the new entry applies to a specific URL (default, all URLs) and sets the ID of the type of the parameter (default, ID of any type). The type IDs can be obtained with the view excludedParamTypes. 
func (a Ascan) AddExcludedParam(name string, t string, url string) (map[string]interface{}, error) {
	m := map[string]string{
		"name": name,
		"type": t,
		"url": url,
	}
	return a.c.Request("ascan/action/addExcludedParam/", m)
}

// Modifies a parameter excluded from the scan. Allows to modify the name, the URL and the type of parameter. The parameter is selected with its index, which can be obtained with the view excludedParams.
func (a Ascan) ModifyExcludedParam(idx string, name string, t string, url string) (map[string]interface{}, error) {
	m := map[string]string{
		"idx": idx,
		"name": name,
		"type": t,
		"url": url,
	}
	return a.c.Request("ascan/action/modifyExcludedParam/", m)
}

// Removes a parameter excluded from the scan, with the given index. The index can be obtained with the view excludedParams.
func (a Ascan) RemoveExcludedParam(idx string) (map[string]interface{}, error) {
	m := map[string]string{
		"idx": idx,
	}
	return a.c.Request("ascan/action/removeExcludedParam/", m)
}

// Skips the scanner using the given IDs of the scan and the scanner.
func (a Ascan) SkipScanner(scanid string, scannerid string) (map[string]interface{}, error) {
	m := map[string]string{
		"scanId": scanid,
		"scannerId": scannerid,
	}
	return a.c.Request("ascan/action/skipScanner/", m)
}

// 
func (a Ascan) SetOptionAttackPolicy(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return a.c.Request("ascan/action/setOptionAttackPolicy/", m)
}

// 
func (a Ascan) SetOptionDefaultPolicy(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return a.c.Request("ascan/action/setOptionDefaultPolicy/", m)
}

// Sets whether or not the active scanner should add a query param to GET requests which do not have parameters to start with.
func (a Ascan) SetOptionAddQueryParam(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return a.c.Request("ascan/action/setOptionAddQueryParam/", m)
}

// 
func (a Ascan) SetOptionAllowAttackOnStart(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return a.c.Request("ascan/action/setOptionAllowAttackOnStart/", m)
}

// 
func (a Ascan) SetOptionDelayInMs(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return a.c.Request("ascan/action/setOptionDelayInMs/", m)
}

// 
func (a Ascan) SetOptionHandleAntiCSRFTokens(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return a.c.Request("ascan/action/setOptionHandleAntiCSRFTokens/", m)
}

// 
func (a Ascan) SetOptionHostPerScan(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return a.c.Request("ascan/action/setOptionHostPerScan/", m)
}

// Sets whether or not the active scanner should inject the HTTP request header X-ZAP-Scan-ID, with the ID of the scanner that's sending the requests.
func (a Ascan) SetOptionInjectPluginIdInHeader(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return a.c.Request("ascan/action/setOptionInjectPluginIdInHeader/", m)
}

// 
func (a Ascan) SetOptionMaxChartTimeInMins(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return a.c.Request("ascan/action/setOptionMaxChartTimeInMins/", m)
}

// 
func (a Ascan) SetOptionMaxResultsToList(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return a.c.Request("ascan/action/setOptionMaxResultsToList/", m)
}

// 
func (a Ascan) SetOptionMaxRuleDurationInMins(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return a.c.Request("ascan/action/setOptionMaxRuleDurationInMins/", m)
}

// 
func (a Ascan) SetOptionMaxScanDurationInMins(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return a.c.Request("ascan/action/setOptionMaxScanDurationInMins/", m)
}

// 
func (a Ascan) SetOptionMaxScansInUI(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return a.c.Request("ascan/action/setOptionMaxScansInUI/", m)
}

// 
func (a Ascan) SetOptionPromptInAttackMode(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return a.c.Request("ascan/action/setOptionPromptInAttackMode/", m)
}

// 
func (a Ascan) SetOptionPromptToClearFinishedScans(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return a.c.Request("ascan/action/setOptionPromptToClearFinishedScans/", m)
}

// 
func (a Ascan) SetOptionRescanInAttackMode(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return a.c.Request("ascan/action/setOptionRescanInAttackMode/", m)
}

// Sets whether or not the HTTP Headers of all requests should be scanned. Not just requests that send parameters, through the query or request body.
func (a Ascan) SetOptionScanHeadersAllRequests(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return a.c.Request("ascan/action/setOptionScanHeadersAllRequests/", m)
}

// Sets whether or not the active scanner should scan null JSON values.
func (a Ascan) SetOptionScanNullJsonValues(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return a.c.Request("ascan/action/setOptionScanNullJsonValues/", m)
}

// 
func (a Ascan) SetOptionShowAdvancedDialog(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return a.c.Request("ascan/action/setOptionShowAdvancedDialog/", m)
}

// 
func (a Ascan) SetOptionTargetParamsEnabledRPC(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return a.c.Request("ascan/action/setOptionTargetParamsEnabledRPC/", m)
}

// 
func (a Ascan) SetOptionTargetParamsInjectable(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return a.c.Request("ascan/action/setOptionTargetParamsInjectable/", m)
}

// 
func (a Ascan) SetOptionThreadPerHost(i int) (map[string]interface{}, error) {
	m := map[string]string{
		"Integer": strconv.Itoa(i),
	}
	return a.c.Request("ascan/action/setOptionThreadPerHost/", m)
}

