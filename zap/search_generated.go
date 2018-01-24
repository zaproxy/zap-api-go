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

type Search struct {
	c *Client
}

func (s Search) UrlsByUrlRegex(regex string, baseurl string, start string, count string) (map[string]interface{}, error) {
	m := map[string]string{
		"regex":   regex,
		"baseurl": baseurl,
		"start":   start,
		"count":   count,
	}
	return s.c.Request("search/view/urlsByUrlRegex/", m)
}

func (s Search) UrlsByRequestRegex(regex string, baseurl string, start string, count string) (map[string]interface{}, error) {
	m := map[string]string{
		"regex":   regex,
		"baseurl": baseurl,
		"start":   start,
		"count":   count,
	}
	return s.c.Request("search/view/urlsByRequestRegex/", m)
}

func (s Search) UrlsByResponseRegex(regex string, baseurl string, start string, count string) (map[string]interface{}, error) {
	m := map[string]string{
		"regex":   regex,
		"baseurl": baseurl,
		"start":   start,
		"count":   count,
	}
	return s.c.Request("search/view/urlsByResponseRegex/", m)
}

func (s Search) UrlsByHeaderRegex(regex string, baseurl string, start string, count string) (map[string]interface{}, error) {
	m := map[string]string{
		"regex":   regex,
		"baseurl": baseurl,
		"start":   start,
		"count":   count,
	}
	return s.c.Request("search/view/urlsByHeaderRegex/", m)
}

func (s Search) MessagesByUrlRegex(regex string, baseurl string, start string, count string) (map[string]interface{}, error) {
	m := map[string]string{
		"regex":   regex,
		"baseurl": baseurl,
		"start":   start,
		"count":   count,
	}
	return s.c.Request("search/view/messagesByUrlRegex/", m)
}

func (s Search) MessagesByRequestRegex(regex string, baseurl string, start string, count string) (map[string]interface{}, error) {
	m := map[string]string{
		"regex":   regex,
		"baseurl": baseurl,
		"start":   start,
		"count":   count,
	}
	return s.c.Request("search/view/messagesByRequestRegex/", m)
}

func (s Search) MessagesByResponseRegex(regex string, baseurl string, start string, count string) (map[string]interface{}, error) {
	m := map[string]string{
		"regex":   regex,
		"baseurl": baseurl,
		"start":   start,
		"count":   count,
	}
	return s.c.Request("search/view/messagesByResponseRegex/", m)
}

func (s Search) MessagesByHeaderRegex(regex string, baseurl string, start string, count string) (map[string]interface{}, error) {
	m := map[string]string{
		"regex":   regex,
		"baseurl": baseurl,
		"start":   start,
		"count":   count,
	}
	return s.c.Request("search/view/messagesByHeaderRegex/", m)
}

func (s Search) HarByUrlRegex(regex string, baseurl string, start string, count string) ([]byte, error) {
	m := map[string]string{
		"regex":   regex,
		"baseurl": baseurl,
		"start":   start,
		"count":   count,
	}
	return s.c.RequestOther("search/other/harByUrlRegex/", m)
}

func (s Search) HarByRequestRegex(regex string, baseurl string, start string, count string) ([]byte, error) {
	m := map[string]string{
		"regex":   regex,
		"baseurl": baseurl,
		"start":   start,
		"count":   count,
	}
	return s.c.RequestOther("search/other/harByRequestRegex/", m)
}

func (s Search) HarByResponseRegex(regex string, baseurl string, start string, count string) ([]byte, error) {
	m := map[string]string{
		"regex":   regex,
		"baseurl": baseurl,
		"start":   start,
		"count":   count,
	}
	return s.c.RequestOther("search/other/harByResponseRegex/", m)
}

func (s Search) HarByHeaderRegex(regex string, baseurl string, start string, count string) ([]byte, error) {
	m := map[string]string{
		"regex":   regex,
		"baseurl": baseurl,
		"start":   start,
		"count":   count,
	}
	return s.c.RequestOther("search/other/harByHeaderRegex/", m)
}
