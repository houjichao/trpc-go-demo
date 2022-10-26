package main

import (
	"context"
	"github.com/houjichao/trpc-go-demo/base"
	"log"

	pb "github.com/houjichao/trpcprotocol"

	_ "github.com/houjichao/trpc-go-demo/base"
)



type demo1Impl struct{}

func (s *demo1Impl) Demo1(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	// implement business logic here ...
	// ...
	rsp := &pb.HelloReply{}

	var maxNum = base.Max(1,10)
	println(maxNum)
	log.Print(maxNum)

	base.ArrayDemo()

	return rsp, nil
}
