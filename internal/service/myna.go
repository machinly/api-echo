package service

import (
	"context"

	pb "api-echo/api/apiecho/v1"

	"github.com/go-kratos/kratos/v2/transport"
)

type MynaService struct {
	pb.UnimplementedMynaServer
}

func NewMynaService() *MynaService {
	return &MynaService{}
}

func (s *MynaService) Header(ctx context.Context, req *pb.HeaderRequest) (*pb.HeaderReply, error) {
	headerSet := map[string]string{}
	if tr, ok := transport.FromServerContext(ctx); ok {
		keys := tr.RequestHeader().Keys()
		for _, key := range keys {
			headerSet[key] = tr.RequestHeader().Get(key)
		}
	}
	return &pb.HeaderReply{
		Headers: headerSet,
	}, nil
}
