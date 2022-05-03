# OWASP ZAP GO API

The Go implementation to access the [OWASP ZAP API](https://www.zaproxy.org/docs/api/). For more information
about OWASP ZAP consult the (main) [OWASP ZAP project](https://github.com/zaproxy/zaproxy/).

## How to Obtain

The latest released version can be downloaded using:

    go get github.com/zaproxy/zap-api-go

## Getting Help

For help using OWASP ZAP API refer to:
  * [Examples](https://github.com/zaproxy/zap-api-go/tree/master/example) - collection of examples using the library;
  * [API Documentation](https://www.zaproxy.org/docs/api/)
  * [OWASP ZAP User Group](https://groups.google.com/group/zaproxy-users) - for asking questions;

## Updating the Generated Files

Most of the API code is generated from the ZAP java source code.

To regenerate the API code you will need the repos [zaproxy](https://github.com/zaproxy/zaproxy) and [zap-extensions](https://github.com/zaproxy/zap-extensions) checked out at the same level as this one.

You should typically generate the core API calls from the latest release tag e.g.:

```
cd zaproxy
git fetch upstream -t
git checkout tags/v2.11.1
./gradlew generateGoApiEndpoints
cd ..
```

The add-on APIs can be generated from the zap-extensions `main` branch:

```
cd zap-extensions
git pull upstream main
./gradlew generateGoZapApiClientFiles --continue
cd ..
```

The above commands will update the files in `zap-api-go/zap`.

If any new files are created then they should be manually added to `zap-api-go/zap/interface.go` as per the existing files.