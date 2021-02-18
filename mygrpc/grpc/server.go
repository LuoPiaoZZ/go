package grpc1

import (
	"context"
	"fmt"
	"mygrpc/pb"
)

type ApiServer struct {
}

func (s *ApiServer) GetAddRes(ctx context.Context,in *pb.AddReq)(*pb.ReplyResp,error){
	return &pb.ReplyResp{RespStr: fmt.Sprintf("add:%d",in.A+in.B)},nil
}

func (s *ApiServer)GetUserRes(ctx context.Context,in *pb.UserReq)(*pb.ReplyResp,error)  {
	return &pb.ReplyResp{RespStr: fmt.Sprintf("user:%s,age:%d",in.User,in.Age)},nil
}
