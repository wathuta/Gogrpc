package handlers

import (
	"context"

	"github.com/wathuta/myservice"
	pb "github.com/wathuta/pb"
)

type userServiceHandler struct {
	userService myservice.Service
}

func NewUserServiceHandler(userService myservice.Service) *userServiceHandler {
	return &userServiceHandler{
		userService: userService,
	}
}
func (u *userServiceHandler) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	resp := &pb.GetUsersResponse{}
	resultMap, err := u.userService.GetUsers(req.GetIds())
	if err != nil {
		return resp, err
	}
	for _, user := range resultMap {

		resp.Users = append(resp.Users, marshaller(&user))
	}
	return resp, nil
}
func marshaller(u *myservice.User) *pb.User {
	var ret *pb.User
	ret.Id = u.UserID
	ret.Name = u.Name
	return ret
}

func (u *userServiceHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	resp := &pb.GetUserResponse{}
	// resultMap, err := u.userService.GetUser(req.Id)
	// if err != nil {
	// 	return resp, err
	// }
	// for _, user := range resultMap {

	// 	resp.Users = append(resp.Users, marshaller(&user))
	// }
	return resp, nil
}
