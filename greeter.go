package main

import (
	"context"
	"github.com/houjichao/trpc-go-demo/base"
	"log"

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

	var maxNum = base.Max(1,10)
	println(maxNum)
	log.Print(maxNum)

	base.ArrayDemo()

	base.MapDemo()

	base.ForDemo()

	base.ConstDemo()

	base.VarDemo()

	base.StrDemo()

	base.DataTypeConvertDemo()

	base.BaseToStr()

	base.StrToBaseDemo()

	return rsp,nil
}
