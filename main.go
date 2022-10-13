package main

import (
	_ "git.woa.com/opentelemetry/opentelemetry-go-ecosystem/instrumentation/oteltrpc"

	_ "git.code.oa.com/trpc-go/trpc-config-rainbow"
	_ "git.code.oa.com/trpc-go/trpc-filter/debuglog"
	_ "git.code.oa.com/trpc-go/trpc-filter/recovery"
	trpc "git.code.oa.com/trpc-go/trpc-go"
	"git.code.oa.com/trpc-go/trpc-go/log"
	_ "git.code.oa.com/trpc-go/trpc-log-atta"
	_ "git.code.oa.com/trpc-go/trpc-metrics-m007"
	_ "git.code.oa.com/trpc-go/trpc-metrics-runtime"
	_ "git.code.oa.com/trpc-go/trpc-naming-polaris"

	pb "github.com/houjichao/trpcprotocol"
)

func main() {
	s := trpc.NewServer()
	pb.RegisterGreeterService(s, &greeterImpl{})
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
