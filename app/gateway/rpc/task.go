package rpc

import (
	"context"
	"micro-memorandum/idl/pb"
	"micro-memorandum/pkg/e"
)

func TaskCreate(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	resp, err = TaskService.CreateTask(ctx, req)
	if err != nil || resp.Code != e.Success {
		resp.Code = e.Error
		return resp, nil
	}

	return
}

func GetTasksList(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskListResponse, err error) {
	resp, err = TaskService.GetTasksList(ctx, req)
	if err != nil || resp.Code != e.Success {
		resp.Code = e.Error
		return resp, nil
	}

	return
}

func GetTask(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	resp, err = TaskService.GetTask(ctx, req)
	if err != nil || resp.Code != e.Success {
		resp.Code = e.Error
		return resp, nil
	}

	return
}

func UpdateTask(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	resp, err = TaskService.UpdateTask(ctx, req)
	if err != nil || resp.Code != e.Success {
		resp.Code = e.Error
		return resp, nil
	}

	return
}

func DeleteTask(ctx context.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	resp, err = TaskService.DeleteTask(ctx, req)
	if err != nil || resp.Code != e.Success {
		resp.Code = e.Error
		return resp, nil
	}

	return
}
