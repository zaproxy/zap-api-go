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

type Selenium struct {
	c *Client
}

// This component is optional and therefore the API will only work if it is installed
func (s Selenium) OptionBrowserExtensions() (map[string]interface{}, error) {
	return s.c.Request("selenium/view/optionBrowserExtensions/", nil)
}

// Returns the current path to Chrome binary
//
// This component is optional and therefore the API will only work if it is installed
func (s Selenium) OptionChromeBinaryPath() (map[string]interface{}, error) {
	return s.c.Request("selenium/view/optionChromeBinaryPath/", nil)
}

// Returns the current path to ChromeDriver
//
// This component is optional and therefore the API will only work if it is installed
func (s Selenium) OptionChromeDriverPath() (map[string]interface{}, error) {
	return s.c.Request("selenium/view/optionChromeDriverPath/", nil)
}

// Returns the current path to Firefox binary
//
// This component is optional and therefore the API will only work if it is installed
func (s Selenium) OptionFirefoxBinaryPath() (map[string]interface{}, error) {
	return s.c.Request("selenium/view/optionFirefoxBinaryPath/", nil)
}

// Returns the current path to Firefox driver (geckodriver)
//
// This component is optional and therefore the API will only work if it is installed
func (s Selenium) OptionFirefoxDriverPath() (map[string]interface{}, error) {
	return s.c.Request("selenium/view/optionFirefoxDriverPath/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (s Selenium) OptionIeDriverPath() (map[string]interface{}, error) {
	return s.c.Request("selenium/view/optionIeDriverPath/", nil)
}

// This component is optional and therefore the API will only work if it is installed
func (s Selenium) OptionLastDirectory() (map[string]interface{}, error) {
	return s.c.Request("selenium/view/optionLastDirectory/", nil)
}

// Returns the current path to PhantomJS binary
//
// This component is optional and therefore the API will only work if it is installed
func (s Selenium) OptionPhantomJsBinaryPath() (map[string]interface{}, error) {
	return s.c.Request("selenium/view/optionPhantomJsBinaryPath/", nil)
}

// Sets the current path to Chrome binary
//
// This component is optional and therefore the API will only work if it is installed
func (s Selenium) SetOptionChromeBinaryPath(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return s.c.Request("selenium/action/setOptionChromeBinaryPath/", m)
}

// Sets the current path to ChromeDriver
//
// This component is optional and therefore the API will only work if it is installed
func (s Selenium) SetOptionChromeDriverPath(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return s.c.Request("selenium/action/setOptionChromeDriverPath/", m)
}

// Sets the current path to Firefox binary
//
// This component is optional and therefore the API will only work if it is installed
func (s Selenium) SetOptionFirefoxBinaryPath(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return s.c.Request("selenium/action/setOptionFirefoxBinaryPath/", m)
}

// Sets the current path to Firefox driver (geckodriver)
//
// This component is optional and therefore the API will only work if it is installed
func (s Selenium) SetOptionFirefoxDriverPath(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return s.c.Request("selenium/action/setOptionFirefoxDriverPath/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Selenium) SetOptionIeDriverPath(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return s.c.Request("selenium/action/setOptionIeDriverPath/", m)
}

// This component is optional and therefore the API will only work if it is installed
func (s Selenium) SetOptionLastDirectory(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return s.c.Request("selenium/action/setOptionLastDirectory/", m)
}

// Sets the current path to PhantomJS binary
//
// This component is optional and therefore the API will only work if it is installed
func (s Selenium) SetOptionPhantomJsBinaryPath(str string) (map[string]interface{}, error) {
	m := map[string]string{
		"String": str,
	}
	return s.c.Request("selenium/action/setOptionPhantomJsBinaryPath/", m)
}

