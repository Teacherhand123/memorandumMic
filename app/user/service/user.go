package service

import (
	"context"
	"errors"
	"fmt"
	"micro-memorandum/app/user/reposltory/db/dao"
	"micro-memorandum/app/user/reposltory/db/model"
	"micro-memorandum/idl/pb"
	"micro-memorandum/pkg/e"
	"sync"

	"gorm.io/gorm"
)

type UserSrv struct {
}

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

// 懒汉式单例模式 lazy-loading --> 饿汉式
// GetUserSrv 方法：
// 这是标准的懒汉式单例实现。
// 第一次调用时，通过 UserSrvOnce.Do 执行匿名函数，只会创建一次 UserSrv 实例并赋值给 UserSrvIns。
// 以后再调用就直接返回这个实例。
// 这样可以保证线程安全。
func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func BuildUser(item *model.User) *pb.UserModel {
	return &pb.UserModel{
		Id:        uint32(item.ID),
		UserName:  item.UserName,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.CreatedAt.Unix(),
	}
}

func (u *UserSrv) UserLogin(ctx context.Context, req *pb.UserRequest, resp *pb.UserReponse) (err error) {
	resp.Code = e.Success
	// 查看有没有这个人
	user, err := dao.NewUserDao(ctx).FindUserByUserName(req.UserName)
	// fmt.Println((*user).UserName, err == nil)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// err = errors.New("用户不存在")
		resp.Code = e.Error
		return nil
	} else if err != nil {
		err = errors.New("数据库出错")
		resp.Code = e.Error
		return
	}

	if !user.CheckPassword(req.Password) {
		// fmt.Println("密码 错误")
		// err = errors.New("用户密码错误")
		resp.Code = e.Error
		return nil
	}

	fmt.Println("验证完成 没有问题")

	resp.UserDetail = BuildUser(user)
	return
}

func (u *UserSrv) UserRegister(ctx context.Context, req *pb.UserRequest, resp *pb.UserReponse) (err error) {
	resp.Code = e.Success

	if req.Password != req.PasswordConfirm {
		// err = errors.New("密码不一致")
		return nil
	}
	fmt.Println(req.Password)

	// 查看有没有这个用户
	_, err = dao.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if err == nil {
		// err = errors.New("用户已存在")
		resp.Code = e.Error
		return nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// err = errors.New("数据库出错")
		resp.Code = e.Error
		return
	}

	user := &model.User{
		UserName: req.UserName,
	}

	if err = user.SetPassword(req.Password); err != nil {
		resp.Code = e.Error
		return
	}

	if err = dao.NewUserDao(ctx).CreateUser(user); err != nil {
		resp.Code = e.Error
		return
	}

	return
}
