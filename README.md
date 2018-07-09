# http2-check

## Build
```
go get "golang.org/x/net/http2"
env GOOS=linux go build .
```

## Transfer to Linux box and run
```
(local)$ curl --upload-file http2-check https://transfer.sh
https://transfer.sh/tYdsi/http2-check

(remote)$ wget https://transfer.sh/tYdsi/http2-check
(remote)$ chmod +x ./http2-check
(remote)$ ./http2-check

```
