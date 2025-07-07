package http

import (
	"micro-memorandum/app/gateway/rpc"
	"micro-memorandum/idl/pb"
	"micro-memorandum/pkg/ctl"
	"micro-memorandum/pkg/jwt"
	"micro-memorandum/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegisterHandler(ctx *gin.Context) {
	var req pb.UserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserRegisterHandler-ShouldBind"))
		return
	}

	userResp, err := rpc.UserRegister(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserRegisterHandler-UserRegister-RPC"))
		return
	}

	ctx.JSON(http.StatusOK, ctl.RespSucess(ctx, userResp))
}

func UserLoginHandler(ctx *gin.Context) {
	var req pb.UserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserRegisterHandler-ShouldBind"))
		return
	}

	userResp, err := rpc.UserLogin(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserRegisterHandler-UserRegister-RPC"))
		return
	}
	token, err := jwt.GenerateToken(uint(userResp.UserDetail.Id))
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserRegisterHandler-ShouldBind"))
		return
	}

	res := &types.TokenData{
		User:  userResp,
		Token: token,
	}

	ctx.JSON(http.StatusOK, ctl.RespSucess(ctx, res))
}
