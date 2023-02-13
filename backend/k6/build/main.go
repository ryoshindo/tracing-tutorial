package main

import (
	"context"

	"go.k6.io/xk6"
)

func main() {
	builder := xk6.Builder{
		K6Version: "v0.42.0",
		Extensions: []xk6.Dependency{
			{
				PackagePath: "github.com/grafana/xk6-client-tracing",
			},
		},
	}
	err := builder.Build(context.Background(), "/usr/bin/k6")
	if err != nil {
		panic(err)
	}
}
