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

// Returns the URLs of the HTTP messages that match the given regular expression in the URL optionally filtered by URL and paginated with 'start' position and 'count' of messages.
func (s Search) UrlsByUrlRegex(regex string, baseurl string, start string, count string) (map[string]interface{}, error) {
	m := map[string]string{
		"regex": regex,
		"baseurl": baseurl,
		"start": start,
		"count": count,
	}
	return s.c.Request("search/view/urlsByUrlRegex/", m)
}

// Returns the URLs of the HTTP messages that match the given regular expression in the request optionally filtered by URL and paginated with 'start' position and 'count' of messages.
func (s Search) UrlsByRequestRegex(regex string, baseurl string, start string, count string) (map[string]interface{}, error) {
	m := map[string]string{
		"regex": regex,
		"baseurl": baseurl,
		"start": start,
		"count": count,
	}
	return s.c.Request("search/view/urlsByRequestRegex/", m)
}

// Returns the URLs of the HTTP messages that match the given regular expression in the response optionally filtered by URL and paginated with 'start' position and 'count' of messages.
func (s Search) UrlsByResponseRegex(regex string, baseurl string, start string, count string) (map[string]interface{}, error) {
	m := map[string]string{
		"regex": regex,
		"baseurl": baseurl,
		"start": start,
		"count": count,
	}
	return s.c.Request("search/view/urlsByResponseRegex/", m)
}

// Returns the URLs of the HTTP messages that match the given regular expression in the header(s) optionally filtered by URL and paginated with 'start' position and 'count' of messages.
func (s Search) UrlsByHeaderRegex(regex string, baseurl string, start string, count string) (map[string]interface{}, error) {
	m := map[string]string{
		"regex": regex,
		"baseurl": baseurl,
		"start": start,
		"count": count,
	}
	return s.c.Request("search/view/urlsByHeaderRegex/", m)
}

// Returns the HTTP messages that match the given regular expression in the URL optionally filtered by URL and paginated with 'start' position and 'count' of messages.
func (s Search) MessagesByUrlRegex(regex string, baseurl string, start string, count string) (map[string]interface{}, error) {
	m := map[string]string{
		"regex": regex,
		"baseurl": baseurl,
		"start": start,
		"count": count,
	}
	return s.c.Request("search/view/messagesByUrlRegex/", m)
}

// Returns the HTTP messages that match the given regular expression in the request optionally filtered by URL and paginated with 'start' position and 'count' of messages.
func (s Search) MessagesByRequestRegex(regex string, baseurl string, start string, count string) (map[string]interface{}, error) {
	m := map[string]string{
		"regex": regex,
		"baseurl": baseurl,
		"start": start,
		"count": count,
	}
	return s.c.Request("search/view/messagesByRequestRegex/", m)
}

// Returns the HTTP messages that match the given regular expression in the response optionally filtered by URL and paginated with 'start' position and 'count' of messages.
func (s Search) MessagesByResponseRegex(regex string, baseurl string, start string, count string) (map[string]interface{}, error) {
	m := map[string]string{
		"regex": regex,
		"baseurl": baseurl,
		"start": start,
		"count": count,
	}
	return s.c.Request("search/view/messagesByResponseRegex/", m)
}

// Returns the HTTP messages that match the given regular expression in the header(s) optionally filtered by URL and paginated with 'start' position and 'count' of messages.
func (s Search) MessagesByHeaderRegex(regex string, baseurl string, start string, count string) (map[string]interface{}, error) {
	m := map[string]string{
		"regex": regex,
		"baseurl": baseurl,
		"start": start,
		"count": count,
	}
	return s.c.Request("search/view/messagesByHeaderRegex/", m)
}

// Returns the HTTP messages, in HAR format, that match the given regular expression in the URL optionally filtered by URL and paginated with 'start' position and 'count' of messages.
func (s Search) HarByUrlRegex(regex string, baseurl string, start string, count string) ([]byte, error) {
	m := map[string]string{
		"regex": regex,
		"baseurl": baseurl,
		"start": start,
		"count": count,
	}
	return s.c.RequestOther("search/other/harByUrlRegex/", m)
}

// Returns the HTTP messages, in HAR format, that match the given regular expression in the request optionally filtered by URL and paginated with 'start' position and 'count' of messages.
func (s Search) HarByRequestRegex(regex string, baseurl string, start string, count string) ([]byte, error) {
	m := map[string]string{
		"regex": regex,
		"baseurl": baseurl,
		"start": start,
		"count": count,
	}
	return s.c.RequestOther("search/other/harByRequestRegex/", m)
}

// Returns the HTTP messages, in HAR format, that match the given regular expression in the response optionally filtered by URL and paginated with 'start' position and 'count' of messages.
func (s Search) HarByResponseRegex(regex string, baseurl string, start string, count string) ([]byte, error) {
	m := map[string]string{
		"regex": regex,
		"baseurl": baseurl,
		"start": start,
		"count": count,
	}
	return s.c.RequestOther("search/other/harByResponseRegex/", m)
}

// Returns the HTTP messages, in HAR format, that match the given regular expression in the header(s) optionally filtered by URL and paginated with 'start' position and 'count' of messages.
func (s Search) HarByHeaderRegex(regex string, baseurl string, start string, count string) ([]byte, error) {
	m := map[string]string{
		"regex": regex,
		"baseurl": baseurl,
		"start": start,
		"count": count,
	}
	return s.c.RequestOther("search/other/harByHeaderRegex/", m)
}

