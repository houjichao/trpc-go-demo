module github.com/houjichao/trpc-go-demo

go 1.12

replace github.com/houjichao/trpcprotocol => ./stub/github.com/houjichao/trpcprotocol

require (
	git.code.oa.com/trpc-go/trpc-config-rainbow v0.1.24
	git.code.oa.com/trpc-go/trpc-filter/debuglog v0.1.5
	git.code.oa.com/trpc-go/trpc-filter/recovery v0.1.3
	git.code.oa.com/trpc-go/trpc-go v0.9.4
	git.code.oa.com/trpc-go/trpc-log-atta v0.1.14
	git.code.oa.com/trpc-go/trpc-metrics-m007 v0.5.1
	git.code.oa.com/trpc-go/trpc-metrics-runtime v0.3.3
	git.code.oa.com/trpc-go/trpc-naming-polaris v0.3.4
	git.woa.com/opentelemetry/opentelemetry-go-ecosystem/instrumentation/oteltrpc v0.3.1
	github.com/golang/mock v1.5.0
	github.com/houjichao/trpcprotocol v0.0.0-00010101000000-000000000000
)
