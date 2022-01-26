package service

import (
	"context"

	pb "casso/api/user/service/v1"
	"casso/app/user/service/internal/biz"
	"casso/pkg/errors/normal"
)

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	// 数据校验
	if req.NickName == "" || req.Mobile == "" {
		return &pb.CreateUserReply{}, normal.InvalidParams
	}
	// 调用业务用例
	return s.uc.CreateUser(ctx, &biz.User{Mobile: req.Mobile, Pass: req.Pass, Name: req.NickName, Age: req.Age})
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	// 数据校验
	if req.Id == 0 {
		return &pb.GetUserReply{}, normal.InvalidParams
	}
	// 调用业务用例
	return s.uc.GetUser(ctx, req.Id)
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	// 数据校验
	if req.Id == 0 {
		return &pb.UpdateUserReply{}, normal.InvalidParams
	}
	// 调用业务用例
	return s.uc.UpdateUser(ctx, req)
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	// 数据校验
	if req.Id == 0 {
		return &pb.DeleteUserReply{}, normal.InvalidParams
	}
	// 调用业务用例
	return s.uc.DeleteUser(ctx, req.Id)
}

func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}

func (s *UserService) GetToken(ctx context.Context, req *pb.GetTokenRequest) (*pb.GetTokenReply, error) {
	return &pb.GetTokenReply{Token: ""}, nil
}
