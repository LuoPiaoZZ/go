package grpc1

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"mygrpc/pb"
	"net"
)

func Start() {
	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()

	// 注册实现接口方法
	pb.RegisterTestServerServer(s, &ApiServer{})

	reflection.Register(s)

	log.Printf("listen protobuf server in %s", ":9000")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
