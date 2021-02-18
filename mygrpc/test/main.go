package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mygrpc/pb"
)
//调用服务测试：respStr:"add:-8"  ---- respStr:"user:lp,age:22"
func main()  {
	conn,err:=grpc.Dial(":9000",grpc.WithInsecure())
	if err!=nil{
		fmt.Println("err:",err.Error())
	}
	defer conn.Close()

	c:=pb.NewTestServerClient(conn)
	//调用方法
	m1,err:=c.GetAddRes(context.Background(),&pb.AddReq{
		A: 1,
		B:-9,
	})
	if err!=nil{
		fmt.Println("err:",err.Error())
	}
	m2,err:=c.GetUserRes(context.Background(),&pb.UserReq{
		User: "lp",
		Age: 22,
	})
	if err!=nil{
		fmt.Println("err:",err.Error())
	}
	fmt.Println(m1,"----",m2)
}
