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

type Alert struct {
	c *Client
}

// Gets the alert with the given ID, the corresponding HTTP message can be obtained with the 'messageId' field and 'message' API method
func (a Alert) Alert(id string) (map[string]interface{}, error) {
	m := map[string]string{
		"id": id,
	}
	return a.c.Request("alert/view/alert/", m)
}

// Gets the alerts raised by ZAP, optionally filtering by URL or riskId, and paginating with 'start' position and 'count' of alerts
func (a Alert) Alerts(baseurl string, start string, count string, riskid string) (map[string]interface{}, error) {
	m := map[string]string{
		"baseurl": baseurl,
		"start":   start,
		"count":   count,
		"riskId":  riskid,
	}
	return a.c.Request("alert/view/alerts/", m)
}

// Gets number of alerts grouped by each risk level, optionally filtering by URL
func (a Alert) AlertsSummary(baseurl string) (map[string]interface{}, error) {
	m := map[string]string{
		"baseurl": baseurl,
	}
	return a.c.Request("alert/view/alertsSummary/", m)
}

// Gets the number of alerts, optionally filtering by URL or riskId
func (a Alert) NumberOfAlerts(baseurl string, riskid string) (map[string]interface{}, error) {
	m := map[string]string{
		"baseurl": baseurl,
		"riskId":  riskid,
	}
	return a.c.Request("alert/view/numberOfAlerts/", m)
}

// Gets a summary of the alerts, optionally filtered by a 'url'. If 'recurse' is true then all alerts that apply to urls that start with the specified 'url' will be returned, otherwise only those on exactly the same 'url' (ignoring url parameters)
func (a Alert) AlertsByRisk(url string, recurse string) (map[string]interface{}, error) {
	m := map[string]string{
		"url":     url,
		"recurse": recurse,
	}
	return a.c.Request("alert/view/alertsByRisk/", m)
}

// Gets a count of the alerts, optionally filtered as per alertsPerRisk
func (a Alert) AlertCountsByRisk(url string, recurse string) (map[string]interface{}, error) {
	m := map[string]string{
		"url":     url,
		"recurse": recurse,
	}
	return a.c.Request("alert/view/alertCountsByRisk/", m)
}

// Deletes all alerts of the current session.
func (a Alert) DeleteAllAlerts() (map[string]interface{}, error) {
	return a.c.Request("alert/action/deleteAllAlerts/", nil)
}

// Deletes the alert with the given ID.
func (a Alert) DeleteAlert(id string) (map[string]interface{}, error) {
	m := map[string]string{
		"id": id,
	}
	return a.c.Request("alert/action/deleteAlert/", m)
}

// Update the confidence of the alerts.
func (a Alert) UpdateAlertsConfidence(ids string, confidenceid string) (map[string]interface{}, error) {
	m := map[string]string{
		"ids":          ids,
		"confidenceId": confidenceid,
	}
	return a.c.Request("alert/action/updateAlertsConfidence/", m)
}

// Update the risk of the alerts.
func (a Alert) UpdateAlertsRisk(ids string, riskid string) (map[string]interface{}, error) {
	m := map[string]string{
		"ids":    ids,
		"riskId": riskid,
	}
	return a.c.Request("alert/action/updateAlertsRisk/", m)
}

// Update the alert with the given ID, with the provided details.
func (a Alert) UpdateAlert(id string, name string, riskid string, confidenceid string, description string, param string, attack string, otherinfo string, solution string, references string, evidence string, cweid string, wascid string) (map[string]interface{}, error) {
	m := map[string]string{
		"id":           id,
		"name":         name,
		"riskId":       riskid,
		"confidenceId": confidenceid,
		"description":  description,
		"param":        param,
		"attack":       attack,
		"otherInfo":    otherinfo,
		"solution":     solution,
		"references":   references,
		"evidence":     evidence,
		"cweId":        cweid,
		"wascId":       wascid,
	}
	return a.c.Request("alert/action/updateAlert/", m)
}

// Add an alert associated with the given message ID, with the provided details. (The ID of the created alert is returned.)
func (a Alert) AddAlert(messageid string, name string, riskid string, confidenceid string, description string, param string, attack string, otherinfo string, solution string, references string, evidence string, cweid string, wascid string) (map[string]interface{}, error) {
	m := map[string]string{
		"messageId":    messageid,
		"name":         name,
		"riskId":       riskid,
		"confidenceId": confidenceid,
		"description":  description,
		"param":        param,
		"attack":       attack,
		"otherInfo":    otherinfo,
		"solution":     solution,
		"references":   references,
		"evidence":     evidence,
		"cweId":        cweid,
		"wascId":       wascid,
	}
	return a.c.Request("alert/action/addAlert/", m)
}
