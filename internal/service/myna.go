package service

import (
	"context"

	pb "api-echo/api/apiecho/v1"
)

type MynaService struct {
	pb.UnimplementedMynaServer
}

func NewMynaService() *MynaService {
	return &MynaService{}
}

func (s *MynaService) Header(ctx context.Context, req *pb.HeaderRequest) (*pb.HeaderReply, error) {
	return &pb.HeaderReply{}, nil
}
