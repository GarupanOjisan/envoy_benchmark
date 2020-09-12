# Envoy Benchmark

## local test

```bash
# grpcurl
$ grpcurl \
	-d '{"message": "hello"}' \
	-import-path ./proto \
	-proto ./proto/echo/echo.proto \
	-plaintext \
	localhost:8080 \
	echo.EchoServer.Echo


# ghz
$ ghz \
	--import-paths=./proto \
	--insecure \
	--proto=./proto/echo/echo.proto \
	--call=echo.EchoServer.Echo \
	--total=100 \
	--connections=1 \
	--keepalive=30s \
	--concurrency=1 \
	localhost:8080
```