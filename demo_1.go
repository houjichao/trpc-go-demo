package main

import (
	"context"

	pb "github.com/houjichao/trpcprotocol"
)

type demo1Impl struct{}

func (s *demo1Impl) Demo1(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	// implement business logic here ...
	// ...
	rsp := &pb.HelloReply{}
	return rsp, nil
}
