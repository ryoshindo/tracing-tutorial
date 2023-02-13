FROM golang:1.20

WORKDIR /go/src/github.com/ryoshindo/tracing-tutorial/backend

COPY go.mod api/main.go ./

RUN go build main.go

CMD ["./main"]
