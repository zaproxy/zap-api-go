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

type Network struct {
	c *Client
}

// Gets the Root CA certificate validity, in days. Used when generating a new Root CA certificate.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) GetRootCaCertValidity() (map[string]interface{}, error) {
	return n.c.Request("network/view/getRootCaCertValidity/", nil)
}

// Gets the server certificate validity, in days. Used when generating server certificates.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) GetServerCertValidity() (map[string]interface{}, error) {
	return n.c.Request("network/view/getServerCertValidity/", nil)
}

// Gets the aliases used to identify the local servers/proxies.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) GetAliases() (map[string]interface{}, error) {
	return n.c.Request("network/view/getAliases/", nil)
}

// Gets the local servers/proxies.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) GetLocalServers() (map[string]interface{}, error) {
	return n.c.Request("network/view/getLocalServers/", nil)
}

// Gets the authorities that will pass-through the local proxies.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) GetPassThroughs() (map[string]interface{}, error) {
	return n.c.Request("network/view/getPassThroughs/", nil)
}

// Gets the connection timeout, in seconds.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) GetConnectionTimeout() (map[string]interface{}, error) {
	return n.c.Request("network/view/getConnectionTimeout/", nil)
}

// Gets the default user-agent.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) GetDefaultUserAgent() (map[string]interface{}, error) {
	return n.c.Request("network/view/getDefaultUserAgent/", nil)
}

// Gets the TTL (in seconds) of successful DNS queries.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) GetDnsTtlSuccessfulQueries() (map[string]interface{}, error) {
	return n.c.Request("network/view/getDnsTtlSuccessfulQueries/", nil)
}

// Gets the HTTP proxy.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) GetHttpProxy() (map[string]interface{}, error) {
	return n.c.Request("network/view/getHttpProxy/", nil)
}

// Gets the HTTP proxy exclusions.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) GetHttpProxyExclusions() (map[string]interface{}, error) {
	return n.c.Request("network/view/getHttpProxyExclusions/", nil)
}

// Gets the SOCKS proxy.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) GetSocksProxy() (map[string]interface{}, error) {
	return n.c.Request("network/view/getSocksProxy/", nil)
}

// Tells whether or not the HTTP proxy authentication is enabled.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) IsHttpProxyAuthEnabled() (map[string]interface{}, error) {
	return n.c.Request("network/view/isHttpProxyAuthEnabled/", nil)
}

// Tells whether or not the HTTP proxy is enabled.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) IsHttpProxyEnabled() (map[string]interface{}, error) {
	return n.c.Request("network/view/isHttpProxyEnabled/", nil)
}

// Tells whether or not the SOCKS proxy is enabled.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) IsSocksProxyEnabled() (map[string]interface{}, error) {
	return n.c.Request("network/view/isSocksProxyEnabled/", nil)
}

// Tells whether or not to use global HTTP state.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) IsUseGlobalHttpState() (map[string]interface{}, error) {
	return n.c.Request("network/view/isUseGlobalHttpState/", nil)
}

// Generates a new Root CA certificate, used to issue server certificates.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) GenerateRootCaCert() (map[string]interface{}, error) {
	return n.c.Request("network/action/generateRootCaCert/", nil)
}

// Imports a Root CA certificate to be used to issue server certificates.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) ImportRootCaCert(filepath string) (map[string]interface{}, error) {
	m := map[string]string{
		"filePath": filepath,
	}
	return n.c.Request("network/action/importRootCaCert/", m)
}

// Sets the Root CA certificate validity. Used when generating a new Root CA certificate.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) SetRootCaCertValidity(validity string) (map[string]interface{}, error) {
	m := map[string]string{
		"validity": validity,
	}
	return n.c.Request("network/action/setRootCaCertValidity/", m)
}

// Sets the server certificate validity. Used when generating server certificates.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) SetServerCertValidity(validity string) (map[string]interface{}, error) {
	m := map[string]string{
		"validity": validity,
	}
	return n.c.Request("network/action/setServerCertValidity/", m)
}

// Adds an alias for the local servers/proxies.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) AddAlias(name string, enabled string) (map[string]interface{}, error) {
	m := map[string]string{
		"name":    name,
		"enabled": enabled,
	}
	return n.c.Request("network/action/addAlias/", m)
}

// Adds a local server/proxy.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) AddLocalServer(address string, port string, api string, proxy string, behindnat string, decoderesponse string, removeacceptencoding string) (map[string]interface{}, error) {
	m := map[string]string{
		"address":              address,
		"port":                 port,
		"api":                  api,
		"proxy":                proxy,
		"behindNat":            behindnat,
		"decodeResponse":       decoderesponse,
		"removeAcceptEncoding": removeacceptencoding,
	}
	return n.c.Request("network/action/addLocalServer/", m)
}

// Adds an authority to pass-through the local proxies.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) AddPassThrough(authority string, enabled string) (map[string]interface{}, error) {
	m := map[string]string{
		"authority": authority,
		"enabled":   enabled,
	}
	return n.c.Request("network/action/addPassThrough/", m)
}

// Removes an alias.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) RemoveAlias(name string) (map[string]interface{}, error) {
	m := map[string]string{
		"name": name,
	}
	return n.c.Request("network/action/removeAlias/", m)
}

// Removes a local server/proxy.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) RemoveLocalServer(address string, port string) (map[string]interface{}, error) {
	m := map[string]string{
		"address": address,
		"port":    port,
	}
	return n.c.Request("network/action/removeLocalServer/", m)
}

// Removes a pass-through.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) RemovePassThrough(authority string) (map[string]interface{}, error) {
	m := map[string]string{
		"authority": authority,
	}
	return n.c.Request("network/action/removePassThrough/", m)
}

// Sets whether or not an alias is enabled.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) SetAliasEnabled(name string, enabled string) (map[string]interface{}, error) {
	m := map[string]string{
		"name":    name,
		"enabled": enabled,
	}
	return n.c.Request("network/action/setAliasEnabled/", m)
}

// Sets whether or not a pass-through is enabled.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) SetPassThroughEnabled(authority string, enabled string) (map[string]interface{}, error) {
	m := map[string]string{
		"authority": authority,
		"enabled":   enabled,
	}
	return n.c.Request("network/action/setPassThroughEnabled/", m)
}

// Sets the timeout, for reads and connects.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) SetConnectionTimeout(timeout string) (map[string]interface{}, error) {
	m := map[string]string{
		"timeout": timeout,
	}
	return n.c.Request("network/action/setConnectionTimeout/", m)
}

// Sets the default user-agent.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) SetDefaultUserAgent(useragent string) (map[string]interface{}, error) {
	m := map[string]string{
		"userAgent": useragent,
	}
	return n.c.Request("network/action/setDefaultUserAgent/", m)
}

// Sets the TTL of successful DNS queries.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) SetDnsTtlSuccessfulQueries(ttl string) (map[string]interface{}, error) {
	m := map[string]string{
		"ttl": ttl,
	}
	return n.c.Request("network/action/setDnsTtlSuccessfulQueries/", m)
}

// Adds a host to be excluded from the HTTP proxy.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) AddHttpProxyExclusion(host string, enabled string) (map[string]interface{}, error) {
	m := map[string]string{
		"host":    host,
		"enabled": enabled,
	}
	return n.c.Request("network/action/addHttpProxyExclusion/", m)
}

// Removes an HTTP proxy exclusion.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) RemoveHttpProxyExclusion(host string) (map[string]interface{}, error) {
	m := map[string]string{
		"host": host,
	}
	return n.c.Request("network/action/removeHttpProxyExclusion/", m)
}

// Sets the HTTP proxy configuration.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) SetHttpProxy(host string, port string, realm string, username string, password string) (map[string]interface{}, error) {
	m := map[string]string{
		"host":     host,
		"port":     port,
		"realm":    realm,
		"username": username,
		"password": password,
	}
	return n.c.Request("network/action/setHttpProxy/", m)
}

// Sets whether or not the HTTP proxy authentication is enabled.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) SetHttpProxyAuthEnabled(enabled string) (map[string]interface{}, error) {
	m := map[string]string{
		"enabled": enabled,
	}
	return n.c.Request("network/action/setHttpProxyAuthEnabled/", m)
}

// Sets whether or not the HTTP proxy is enabled.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) SetHttpProxyEnabled(enabled string) (map[string]interface{}, error) {
	m := map[string]string{
		"enabled": enabled,
	}
	return n.c.Request("network/action/setHttpProxyEnabled/", m)
}

// Sets whether or not an HTTP proxy exclusion is enabled.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) SetHttpProxyExclusionEnabled(host string, enabled string) (map[string]interface{}, error) {
	m := map[string]string{
		"host":    host,
		"enabled": enabled,
	}
	return n.c.Request("network/action/setHttpProxyExclusionEnabled/", m)
}

// Sets the SOCKS proxy configuration.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) SetSocksProxy(host string, port string, version string, usedns string, username string, password string) (map[string]interface{}, error) {
	m := map[string]string{
		"host":     host,
		"port":     port,
		"version":  version,
		"useDns":   usedns,
		"username": username,
		"password": password,
	}
	return n.c.Request("network/action/setSocksProxy/", m)
}

// Sets whether or not the SOCKS proxy is enabled.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) SetSocksProxyEnabled(enabled string) (map[string]interface{}, error) {
	m := map[string]string{
		"enabled": enabled,
	}
	return n.c.Request("network/action/setSocksProxyEnabled/", m)
}

// Sets whether or not to use the global HTTP state.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) SetUseGlobalHttpState(use string) (map[string]interface{}, error) {
	m := map[string]string{
		"use": use,
	}
	return n.c.Request("network/action/setUseGlobalHttpState/", m)
}

// Adds a client certificate contained in a PKCS#12 file, the certificate is automatically set as active and used.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) AddPkcs12ClientCertificate(filepath string, password string, index string) (map[string]interface{}, error) {
	m := map[string]string{
		"filePath": filepath,
		"password": password,
		"index":    index,
	}
	return n.c.Request("network/action/addPkcs12ClientCertificate/", m)
}

// Sets whether or not to use the active client certificate.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) SetUseClientCertificate(use string) (map[string]interface{}, error) {
	m := map[string]string{
		"use": use,
	}
	return n.c.Request("network/action/setUseClientCertificate/", m)
}

// Provides a PAC file, proxying through the main proxy.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) Proxypac() ([]byte, error) {
	return n.c.RequestOther("network/other/proxy.pac/", nil)
}

// Sets the HTTP proxy configuration.
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) SetProxy(proxy string) ([]byte, error) {
	m := map[string]string{
		"proxy": proxy,
	}
	return n.c.RequestOther("network/other/setProxy/", m)
}

// Gets the Root CA certificate used to issue server certificates. Suitable to import into client applications (e.g. browsers).
//
// This component is optional and therefore the API will only work if it is installed
func (n Network) RootCaCert() ([]byte, error) {
	return n.c.RequestOther("network/other/rootCaCert/", nil)
}
