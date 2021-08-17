package rpc

import (
	"context"
	"github.com/asim/go-micro/v3/client"
	"github.com/leeprince/protobuf/grpc/gin_micro"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/8/15 下午12:21
 * @Desc:
 */

func GetUsers(userId int64) (*gin_micro.GetUsersRsp, error) {
	// 创建新的客户端
	microUserClient := gin_micro.NewUserService("micro.user.server", client.DefaultClient)
	// 调用rpc方法
	rsp, err := microUserClient.GetUsers(context.Background(), &gin_micro.GetUsersReq{
		UserId:userId,
	})
	if err != nil {
		return nil, err
	}
	return rsp, err
}