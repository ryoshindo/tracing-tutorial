# k6 build
FROM golang:1.20 AS build

ARG K6_LOCAL_BUILD_PATH="k6/build"

WORKDIR /go/src/github.com/ryoshindo/tracing-tutorial/backend/k6/build

COPY $K6_LOCAL_BUILD_PATH/go.mod $K6_LOCAL_BUILD_PATH/go.sum $K6_LOCAL_BUILD_PATH/main.go ./

RUN go build main.go
