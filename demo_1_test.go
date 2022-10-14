package main

import (
	"context"
	"reflect"
	"testing"

	_ "git.code.oa.com/trpc-go/trpc-go/http"

	"github.com/golang/mock/gomock"

	pb "github.com/houjichao/trpcprotocol"
)

//go:generate go mod tidy

//go:generate mockgen -destination=stub/github.com/houjichao/trpcprotocol/demo_mock.go -package=trpcprotocol -self_package=github.com/houjichao/trpcprotocol --source=stub/github.com/houjichao/trpcprotocol/demo.trpc.go

func Test_demo1Impl_Demo1(t *testing.T) {
	// 开始写mock逻辑
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	demo1Service := pb.NewMockDemo1Service(ctrl)
	var inorderClient []*gomock.Call
	// 预期行为
	m := demo1Service.EXPECT().Demo1(gomock.Any(), gomock.Any()).AnyTimes()
	m.DoAndReturn(func(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
		s := &demo1Impl{}
		return s.Demo1(ctx, req)
	})
	gomock.InOrder(inorderClient...)

	// 开始写单元测试逻辑
	type args struct {
		ctx context.Context
		req *pb.HelloRequest
		rsp *pb.HelloReply
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rsp *pb.HelloReply
			var err error
			if rsp, err = demo1Service.Demo1(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("demo1Impl.Demo1() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rsp, tt.args.rsp) {
				t.Errorf("demo1Impl.Demo1() rsp got = %v, want %v", rsp, tt.args.rsp)
			}
		})
	}
}
