package http

import (
	"micro-memorandum/app/gateway/rpc"
	"micro-memorandum/idl/pb"
	"micro-memorandum/pkg/ctl"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-ShouldBind"))
		return
	}

	user, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-GetUserInfo-jwt"))
		return
	}

	req.Uid = uint64(user.Id)
	taskResp, err := rpc.TaskCreate(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-CreateTask"))
		return
	}

	ctx.JSON(http.StatusOK, ctl.RespSucess(ctx, taskResp))
}

func UpdateTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-ShouldBind"))
		return
	}

	user, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-GetUserInfo"))
		return
	}

	req.Uid = uint64(user.Id)
	taskResp, err := rpc.UpdateTask(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-UpdateTask"))
		return
	}

	ctx.JSON(http.StatusOK, ctl.RespSucess(ctx, taskResp))
}

func DeleteTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-ShouldBind"))
		return
	}

	user, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-GetUserInfo"))
		return
	}

	req.Uid = uint64(user.Id)
	taskResp, err := rpc.DeleteTask(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-DeleteTask"))
		return
	}

	ctx.JSON(http.StatusOK, ctl.RespSucess(ctx, taskResp))
}

func ListTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-ShouldBind"))
		return
	}

	user, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-GetUserInfo"))
		return
	}

	req.Uid = uint64(user.Id)
	taskResp, err := rpc.GetTasksList(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-ListTask"))
		return
	}

	ctx.JSON(http.StatusOK, ctl.RespSucess(ctx, taskResp))
}

func GetTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-ShouldBind"))
		return
	}

	user, err := ctl.GetUserInfo(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-GetUserInfo"))
		return
	}

	req.Uid = uint64(user.Id)
	taskResp, err := rpc.GetTask(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "CreateTaskHandler-GetTask"))
		return
	}

	ctx.JSON(http.StatusOK, ctl.RespSucess(ctx, taskResp))
}
