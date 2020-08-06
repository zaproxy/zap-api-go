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

type Websocket struct {
	c *Client
}

// Returns all of the registered web socket channels
//
// This component is optional and therefore the API will only work if it is installed
func (w Websocket) Channels() (map[string]interface{}, error) {
	return w.c.Request("websocket/view/channels/", nil)
}

// Returns full details of the message specified by the channelId and messageId
//
// This component is optional and therefore the API will only work if it is installed
func (w Websocket) Message(channelid string, messageid string) (map[string]interface{}, error) {
	m := map[string]string{
		"channelId": channelid,
		"messageId": messageid,
	}
	return w.c.Request("websocket/view/message/", m)
}

// Returns a list of all of the messages that meet the given criteria (all optional), where channelId is a channel identifier, start is the offset to start returning messages from (starting from 0), count is the number of messages to return (default no limit) and payloadPreviewLength is the maximum number bytes to return for the payload contents
//
// This component is optional and therefore the API will only work if it is installed
func (w Websocket) Messages(channelid string, start string, count string, payloadpreviewlength string) (map[string]interface{}, error) {
	m := map[string]string{
		"channelId":            channelid,
		"start":                start,
		"count":                count,
		"payloadPreviewLength": payloadpreviewlength,
	}
	return w.c.Request("websocket/view/messages/", m)
}

// Returns a text representation of an intercepted websockets message
//
// This component is optional and therefore the API will only work if it is installed
func (w Websocket) BreakTextMessage() (map[string]interface{}, error) {
	return w.c.Request("websocket/view/breakTextMessage/", nil)
}

// Sends the specified message on the channel specified by channelId, if outgoing is 'True' then the message will be sent to the server and if it is 'False' then it will be sent to the client
//
// This component is optional and therefore the API will only work if it is installed
func (w Websocket) SendTextMessage(channelid string, outgoing string, message string) (map[string]interface{}, error) {
	m := map[string]string{
		"channelId": channelid,
		"outgoing":  outgoing,
		"message":   message,
	}
	return w.c.Request("websocket/action/sendTextMessage/", m)
}

// Sets the text message for an intercepted websockets message
//
// This component is optional and therefore the API will only work if it is installed
func (w Websocket) SetBreakTextMessage(message string, outgoing string) (map[string]interface{}, error) {
	m := map[string]string{
		"message":  message,
		"outgoing": outgoing,
	}
	return w.c.Request("websocket/action/setBreakTextMessage/", m)
}
