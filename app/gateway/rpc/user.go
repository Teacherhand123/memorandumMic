package rpc

import (
	"context"
	"micro-memorandum/idl/pb"
	"micro-memorandum/pkg/e"
)

func UserLogin(ctx context.Context, req *pb.UserRequest) (resp *pb.UserReponse, err error) {
	resp, err = UserService.UserLogin(ctx, req)
	if err != nil || resp.Code != e.Success {
		resp.Code = e.Error
		return
	}
	return
}

func UserRegister(ctx context.Context, req *pb.UserRequest) (resp *pb.UserReponse, err error) {
	resp, err = UserService.UserRegister(ctx, req)
	if err != nil || resp.Code != e.Success {
		resp.Code = e.Error
		return
	}
	return
}
