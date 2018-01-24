## How to run example

#### run zaproxy daemon
```shell
docker run --name zap -u zap -p 8080:8080 -i owasp/zap2docker-stable zap.sh -daemon -host 0.0.0.0 -port 8080 -config api.disablekey=true -config api.addrs.addr.name=.\* -config api.addrs.addr.regex=true
```

#### run bodgeit
```shell
docker run --name bodgeit --rm -it psiinon/bodgeit
```

#### get the ip address of bodgeit
```shell
docker inspect bodgeit | grep IPAddress
```

#### run example
```shell
go run example.go -target=http://{$bodgeit_IP}:8080/bodgeit
```

#### How to connect to ZAP server using HTTPS

1. Client connect to ZAP server using HTTP by default. To instruct client to connect ZAP server using HTTPS, you should:
* change the ZAP base path using HTTPS
* pass in your TLS config

2. example:
```golang
	cfg := &zap.Config{
		Base:      zap.DefaultHTTPSBase,
		BaseOther: zap.DefaultHTTPSBaseOther,
		// you can set your custom certificates here
		TLSConfig: tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client, err := zap.NewClient(cfg)
	if err != nil {
	    log.Fatal(err)
	}
	/*
	use the client...
	*/
```