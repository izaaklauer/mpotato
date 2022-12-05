# mpotato

izaak's great potato serving microservice

### First clone

Run `make start` to download go modules and regenerate protobufs.

### Local testing

Start your server with `make serve`

```
$ make serve

go run cmd/mpotato-api/main.go config/config_local.hcl
starting mpotato.......
2022/12/02 20:58:46 Serving on "localhost:8080"

```

Then send a request via grpcurl:

```
$ grpcurl -plaintext -d '{"message": "hello from local development"}' localhost:8080 mpotato.v1.MpotatoService/HelloWorld

{
  "configMessage": "hello from mpotato",
  "requestMessage": "hello from local development",
  "now": "2022-12-03T02:01:19.505743Z"
}
```
