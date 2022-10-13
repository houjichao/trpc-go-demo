package main

import (
	"context"

	pb "github.com/houjichao/trpcprotocol"
)

type greeterImpl struct{}

func (s *greeterImpl) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	// implement business logic here ...
	// ...
	rsp := &pb.HelloReply{}
	return rsp, nil
}
