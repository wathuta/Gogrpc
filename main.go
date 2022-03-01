package main

import (
	"log"
	"net"

	core "github.com/wathuta/Core"
	"github.com/wathuta/handlers"
	"github.com/wathuta/pb"
	"google.golang.org/grpc"
)

func main() {
	userService := core.NewService()

	usersevicehandler := handlers.NewUserServiceHandler(userService)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, usersevicehandler)
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
