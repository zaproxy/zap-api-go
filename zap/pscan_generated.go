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

type Pscan struct {
	c *Client
}

// Tells whether or not the passive scan should be performed only on messages that are in scope.
func (p Pscan) ScanOnlyInScope() (map[string]interface{}, error) {
	return p.c.Request("pscan/view/scanOnlyInScope/", nil)
}

// The number of records the passive scanner still has to scan
func (p Pscan) RecordsToScan() (map[string]interface{}, error) {
	return p.c.Request("pscan/view/recordsToScan/", nil)
}

// Lists all passive scan rules with their ID, name, enabled state, and alert threshold.
func (p Pscan) Scanners() (map[string]interface{}, error) {
	return p.c.Request("pscan/view/scanners/", nil)
}

// Show information about the passive scan rule currently being run (if any).
func (p Pscan) CurrentRule() (map[string]interface{}, error) {
	return p.c.Request("pscan/view/currentRule/", nil)
}

// Show information about the passive scan tasks currently being run (if any).
func (p Pscan) CurrentTasks() (map[string]interface{}, error) {
	return p.c.Request("pscan/view/currentTasks/", nil)
}

// Gets the maximum number of alerts a passive scan rule should raise.
func (p Pscan) MaxAlertsPerRule() (map[string]interface{}, error) {
	return p.c.Request("pscan/view/maxAlertsPerRule/", nil)
}

// Sets whether or not the passive scanning is enabled (Note: the enabled state is not persisted).
func (p Pscan) SetEnabled(enabled string) (map[string]interface{}, error) {
	m := map[string]string{
		"enabled": enabled,
	}
	return p.c.Request("pscan/action/setEnabled/", m)
}

// Sets whether or not the passive scan should be performed only on messages that are in scope.
func (p Pscan) SetScanOnlyInScope(onlyinscope string) (map[string]interface{}, error) {
	m := map[string]string{
		"onlyInScope": onlyinscope,
	}
	return p.c.Request("pscan/action/setScanOnlyInScope/", m)
}

// Enables all passive scan rules
func (p Pscan) EnableAllScanners() (map[string]interface{}, error) {
	return p.c.Request("pscan/action/enableAllScanners/", nil)
}

// Disables all passive scan rules
func (p Pscan) DisableAllScanners() (map[string]interface{}, error) {
	return p.c.Request("pscan/action/disableAllScanners/", nil)
}

// Enables all passive scan rules with the given IDs (comma separated list of IDs)
func (p Pscan) EnableScanners(ids string) (map[string]interface{}, error) {
	m := map[string]string{
		"ids": ids,
	}
	return p.c.Request("pscan/action/enableScanners/", m)
}

// Disables all passive scan rules with the given IDs (comma separated list of IDs)
func (p Pscan) DisableScanners(ids string) (map[string]interface{}, error) {
	m := map[string]string{
		"ids": ids,
	}
	return p.c.Request("pscan/action/disableScanners/", m)
}

// Sets the alert threshold of the passive scanner with the given ID, accepted values for alert threshold: OFF, DEFAULT, LOW, MEDIUM and HIGH
func (p Pscan) SetScannerAlertThreshold(id string, alertthreshold string) (map[string]interface{}, error) {
	m := map[string]string{
		"id": id,
		"alertThreshold": alertthreshold,
	}
	return p.c.Request("pscan/action/setScannerAlertThreshold/", m)
}

// Sets the maximum number of alerts a passive scan rule should raise.
func (p Pscan) SetMaxAlertsPerRule(maxalerts string) (map[string]interface{}, error) {
	m := map[string]string{
		"maxAlerts": maxalerts,
	}
	return p.c.Request("pscan/action/setMaxAlertsPerRule/", m)
}

// Disables all passive scan tags.
func (p Pscan) DisableAllTags() (map[string]interface{}, error) {
	return p.c.Request("pscan/action/disableAllTags/", nil)
}

// Enables all passive scan tags.
func (p Pscan) EnableAllTags() (map[string]interface{}, error) {
	return p.c.Request("pscan/action/enableAllTags/", nil)
}

// Clears the passive scan queue.
func (p Pscan) ClearQueue() (map[string]interface{}, error) {
	return p.c.Request("pscan/action/clearQueue/", nil)
}

