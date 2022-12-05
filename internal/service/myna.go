package service

import (
	"context"

	pb "api-echo/api/apiecho/v1"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-kratos/kratos/v2/transport"
)

type MynaService struct {
	pb.UnimplementedMynaServer
	log *log.Helper
}

func NewMynaService(logger log.Logger) *MynaService {
	return &MynaService{log: log.NewHelper(logger)}
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

func (s *MynaService) FillData(ctx context.Context, req *pb.FillDataRequest) (*pb.FillDataReply, error) {
	s.log.Debug(req.GetCount(), req.GetContent())
	set := map[int]string{}
	for i := range make([]any, int(req.GetCount())) {
		s.log.Debug(i)
		set[i] = req.GetContent()
	}
	return &pb.FillDataReply{}, nil
}

func (s *MynaService) GetData(ctx context.Context, req *pb.GetDataRequest) (*pb.GetDataReply, error) {
	s.log.Debug(req.GetCount(), req.GetContent())
	set := make([]string, req.GetCount())
	for i := range set {
		s.log.Debug(i)
		set[i] = req.GetContent()
	}
	return &pb.GetDataReply{
		Contents: set,
	}, nil

}

func (s *MynaService) Status(ctx context.Context, req *pb.StatusRequest) (*pb.StatusReply, error) {
	return &pb.StatusReply{}, nil
}
