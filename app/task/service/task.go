package service

import (
	"context"
	"encoding/json"
	"errors"
	"micro-memorandum/app/task/reposltory/db/dao"
	"micro-memorandum/app/task/reposltory/db/model"
	"micro-memorandum/app/task/reposltory/mq"
	"micro-memorandum/idl/pb"
	"micro-memorandum/pkg/e"
	"sync"

	"gorm.io/gorm"
)

type TaskSrv struct {
}

var TaskSrvIns *TaskSrv
var TaskSrvOnce sync.Once

// 懒汉式单例模式 lazy-loading --> 饿汉式
// GetTaskSrv 方法：
// 这是标准的懒汉式单例实现。
// 第一次调用时，通过 TaskSrvOnce.Do 执行匿名函数，只会创建一次 TaskSrv 实例并赋值给 TaskSrvIns。
// 以后再调用就直接返回这个实例。
// 这样可以保证线程安全。
func GetTaskSrv() *TaskSrv {
	TaskSrvOnce.Do(func() {
		TaskSrvIns = &TaskSrv{}
	})
	return TaskSrvIns
}

// create task 送到mq， 通过mq入库
func (t *TaskSrv) CreateTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	resp.Code = e.Success
	body, _ := json.Marshal(req)
	err = mq.SendMessage2MQ(body)
	if err != nil {
		resp.Code = e.Error
		return
	}

	return
}

func TaskMQ2DB(ctx context.Context, req *pb.TaskRequest) error {
	m := &model.Task{
		Uid:       uint(req.Uid),
		Title:     req.Title,
		Status:    int(req.Status),
		Content:   req.Content,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}

	return dao.NewTaskDao(ctx).CreateTask(m)
}

func (t *TaskSrv) GetTasksList(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskListResponse) (err error) {
	resp.Code = e.Success
	if req.Limit == 0 {
		req.Limit = 10
	}

	r, count, err := dao.NewTaskDao(ctx).LimitTaskByUserId(req.Uid, int(req.Start), int(req.Limit))
	if err != nil {
		resp.Code = e.Error
		return nil
	}

	var taskRes []*pb.TaskModel
	for _, item := range r {
		taskRes = append(taskRes, BuildTask(item))
	}

	resp.TaskList = taskRes
	resp.Count = uint32(count)
	return
}

func (t *TaskSrv) GetTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	resp.Code = e.Success
	r, err := dao.NewTaskDao(ctx).GetTaskByTaskIdAndUserId(req.Id, req.Uid)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.Code = e.Error
		return nil
	} else if err != nil {
		err = errors.New("数据库出错")
		resp.Code = e.Error
		return
	}

	resp.TaskDetail = BuildTask(r)
	return
}

func (t *TaskSrv) UpdateTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	resp.Code = e.Success
	r, err := dao.NewTaskDao(ctx).UpdateTask(req)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.Code = e.Error
		return
	} else if err != nil {
		err = errors.New("数据库出错")
		resp.Code = e.Error
		return
	}

	resp.TaskDetail = BuildTask(r)
	return
}

func (t *TaskSrv) DeleteTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	resp.Code = e.Success
	err = dao.NewTaskDao(ctx).DeleteTask(req.Id, req.Uid)
	r := &model.Task{}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.Code = e.Error
		return nil
	} else if err != nil {
		err = errors.New("数据库出错")
		resp.Code = e.Error
		return
	}

	resp.TaskDetail = BuildTask(r)
	return
}

func BuildTask(item *model.Task) *pb.TaskModel {
	return &pb.TaskModel{
		Id:         uint64(item.ID),
		Uid:        uint64(item.Uid),
		Title:      item.Title,
		Content:    item.Content,
		StartTime:  item.StartTime,
		EndTime:    item.EndTime,
		Status:     int64(item.Status),
		CreateTime: item.CreatedAt.Unix(),
		UpdateTime: item.UpdatedAt.Unix(),
	}
}
