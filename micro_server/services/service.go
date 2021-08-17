package services

import (
	"github.com/asim/go-micro/v3"
	pb "github.com/leeprince/protobuf/grpc/gin_micro"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/8/13 下午11:27
 * @Desc:
 */

func Run()  {
	// 创建服务
	service := micro.NewService(
		micro.Name("user"),
		// micro.Add
	)
	// 初始化服务
	service.Init()
	
	// 注册服务
	pb.RegisterUserServiceHandler(service.Server(), new(UsesrService))
	
	service.Run()
}