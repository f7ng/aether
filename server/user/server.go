package user

import (
	"context"
	"fmt"
	pb "github.com/f7ng/aether/protobuf"
)

type Server struct {
	pb.UnimplementedUserServer
}

func (s *Server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{Name: fmt.Sprintf("%v", in.GetId())}, nil
}
