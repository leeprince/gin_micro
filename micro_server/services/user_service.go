package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/leeprince/protobuf/grpc/gin_micro"
	"micro_server/dao"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/8/13 下午11:41
 * @Desc:
 */

type UsesrService struct {
	
}

func (s *UsesrService) GetUsers(ctx context.Context, req *gin_micro.GetUsersReq, rsp *gin_micro.GetUsersRsp) error  {
	userId := req.UserId
	
	userDao := dao.NewUserDao()
	user, err := userDao.GetUserByUserId(userId)
	if err != nil {
		return err
	}
	jsonBytes,_ := json.Marshal(user)
	
	rsp.Code = 0
	rsp.Message = "successful."
	rsp.Data = jsonBytes
	
	fmt.Println("GetUsers.rsp:", rsp)
	return nil
}


 