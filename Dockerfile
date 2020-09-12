FROM golang:1.14.2-alpine AS base

COPY . /go/src/github.com/garupanojisan/envoy_benchmark

RUN go install github.com/garupanojisan/envoy_benchmark

FROM scratch

WORKDIR /app

COPY --from=base /go/bin .

ENTRYPOINT ["/app/spinnaker-example"]

EXPOSE 8080