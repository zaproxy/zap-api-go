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

type Autoupdate struct {
	c *Client
}

// Returns the latest version number
func (a Autoupdate) LatestVersionNumber() (map[string]interface{}, error) {
	return a.c.Request("autoupdate/view/latestVersionNumber/", nil)
}

// Returns 'true' if ZAP is on the latest version
func (a Autoupdate) IsLatestVersion() (map[string]interface{}, error) {
	return a.c.Request("autoupdate/view/isLatestVersion/", nil)
}

// Return a list of all of the installed add-ons
func (a Autoupdate) InstalledAddons() (map[string]interface{}, error) {
	return a.c.Request("autoupdate/view/installedAddons/", nil)
}

// Returns a list with all local add-ons, installed or not.
func (a Autoupdate) LocalAddons() (map[string]interface{}, error) {
	return a.c.Request("autoupdate/view/localAddons/", nil)
}

// Return a list of any add-ons that have been added to the Marketplace since the last check for updates
func (a Autoupdate) NewAddons() (map[string]interface{}, error) {
	return a.c.Request("autoupdate/view/newAddons/", nil)
}

// Return a list of any add-ons that have been changed in the Marketplace since the last check for updates
func (a Autoupdate) UpdatedAddons() (map[string]interface{}, error) {
	return a.c.Request("autoupdate/view/updatedAddons/", nil)
}

// Return a list of all of the add-ons on the ZAP Marketplace (this information is read once and then cached)
func (a Autoupdate) MarketplaceAddons() (map[string]interface{}, error) {
	return a.c.Request("autoupdate/view/marketplaceAddons/", nil)
}

// 
func (a Autoupdate) OptionAddonDirectories() (map[string]interface{}, error) {
	return a.c.Request("autoupdate/view/optionAddonDirectories/", nil)
}

// 
func (a Autoupdate) OptionDayLastChecked() (map[string]interface{}, error) {
	return a.c.Request("autoupdate/view/optionDayLastChecked/", nil)
}

// 
func (a Autoupdate) OptionDayLastInstallWarned() (map[string]interface{}, error) {
	return a.c.Request("autoupdate/view/optionDayLastInstallWarned/", nil)
}

// 
func (a Autoupdate) OptionDayLastUpdateWarned() (map[string]interface{}, error) {
	return a.c.Request("autoupdate/view/optionDayLastUpdateWarned/", nil)
}

// 
func (a Autoupdate) OptionDownloadDirectory() (map[string]interface{}, error) {
	return a.c.Request("autoupdate/view/optionDownloadDirectory/", nil)
}

// 
func (a Autoupdate) OptionCheckAddonUpdates() (map[string]interface{}, error) {
	return a.c.Request("autoupdate/view/optionCheckAddonUpdates/", nil)
}

// 
func (a Autoupdate) OptionCheckOnStart() (map[string]interface{}, error) {
	return a.c.Request("autoupdate/view/optionCheckOnStart/", nil)
}

// 
func (a Autoupdate) OptionDownloadNewRelease() (map[string]interface{}, error) {
	return a.c.Request("autoupdate/view/optionDownloadNewRelease/", nil)
}

// 
func (a Autoupdate) OptionInstallAddonUpdates() (map[string]interface{}, error) {
	return a.c.Request("autoupdate/view/optionInstallAddonUpdates/", nil)
}

// 
func (a Autoupdate) OptionInstallScannerRules() (map[string]interface{}, error) {
	return a.c.Request("autoupdate/view/optionInstallScannerRules/", nil)
}

// 
func (a Autoupdate) OptionReportAlphaAddons() (map[string]interface{}, error) {
	return a.c.Request("autoupdate/view/optionReportAlphaAddons/", nil)
}

// 
func (a Autoupdate) OptionReportBetaAddons() (map[string]interface{}, error) {
	return a.c.Request("autoupdate/view/optionReportBetaAddons/", nil)
}

// 
func (a Autoupdate) OptionReportReleaseAddons() (map[string]interface{}, error) {
	return a.c.Request("autoupdate/view/optionReportReleaseAddons/", nil)
}

// Downloads the latest release, if any 
func (a Autoupdate) DownloadLatestRelease() (map[string]interface{}, error) {
	return a.c.Request("autoupdate/action/downloadLatestRelease/", nil)
}

// Installs or updates the specified add-on, returning when complete (i.e. not asynchronously)
func (a Autoupdate) InstallAddon(id string) (map[string]interface{}, error) {
	m := map[string]string{
		"id": id,
	}
	return a.c.Request("autoupdate/action/installAddon/", m)
}

// 
func (a Autoupdate) InstallLocalAddon(file string) (map[string]interface{}, error) {
	m := map[string]string{
		"file": file,
	}
	return a.c.Request("autoupdate/action/installLocalAddon/", m)
}

// Uninstalls the specified add-on 
func (a Autoupdate) UninstallAddon(id string) (map[string]interface{}, error) {
	m := map[string]string{
		"id": id,
	}
	return a.c.Request("autoupdate/action/uninstallAddon/", m)
}

// 
func (a Autoupdate) SetOptionCheckAddonUpdates(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return a.c.Request("autoupdate/action/setOptionCheckAddonUpdates/", m)
}

// 
func (a Autoupdate) SetOptionCheckOnStart(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return a.c.Request("autoupdate/action/setOptionCheckOnStart/", m)
}

// 
func (a Autoupdate) SetOptionDownloadNewRelease(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return a.c.Request("autoupdate/action/setOptionDownloadNewRelease/", m)
}

// 
func (a Autoupdate) SetOptionInstallAddonUpdates(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return a.c.Request("autoupdate/action/setOptionInstallAddonUpdates/", m)
}

// 
func (a Autoupdate) SetOptionInstallScannerRules(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return a.c.Request("autoupdate/action/setOptionInstallScannerRules/", m)
}

// 
func (a Autoupdate) SetOptionReportAlphaAddons(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return a.c.Request("autoupdate/action/setOptionReportAlphaAddons/", m)
}

// 
func (a Autoupdate) SetOptionReportBetaAddons(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return a.c.Request("autoupdate/action/setOptionReportBetaAddons/", m)
}

// 
func (a Autoupdate) SetOptionReportReleaseAddons(boolean bool) (map[string]interface{}, error) {
	m := map[string]string{
		"Boolean": strconv.FormatBool(boolean),
	}
	return a.c.Request("autoupdate/action/setOptionReportReleaseAddons/", m)
}

