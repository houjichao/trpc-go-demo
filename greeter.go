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
	rsp.Msg = req.Msg
	return rsp, nil
}

func (s *greeterImpl) Demo1(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	// implement business logic here ...
	// ...
	rsp := &pb.HelloReply{}
	rsp.Msg = "Hello, I am tRPC-Go server."
	return rsp,nil
}
