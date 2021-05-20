package grpcserver

import (
	"context"
	"log"

	"github.com/alibug/go-identity-casbin-server/domain"
	pb "github.com/alibug/go-identity-casbin-server/gen/casbin/proto"
)

type casbinService struct {
	permissionsUC domain.PermissionUseCase
}

// NewCasbinService -
func NewCasbinService(p domain.PermissionUseCase) pb.CasbinServer {
	return &casbinService{permissionsUC: p}
}

func (s *casbinService) HasPermissionForUser(ctx context.Context, r *pb.PermissionRequest) (*pb.BoolReply, error) {
	result, err := s.permissionsUC.HasPermissionForUser(r)
	if err != nil {
		log.Printf("😯 err:%v", err)
		return nil, err
	}
	log.Println("final result: ", result)
	return &pb.BoolReply{Res: result}, nil
}

func (s *casbinService) LoadPolicy(ctx context.Context, r *pb.Empty) (*pb.Empty, error) {
	err := s.permissionsUC.LoadPolicy()
	if err != nil {
		log.Printf("😯 err:%v", err)
		return nil, err
	}
	log.Println("✅  load policy ok!")
	return &pb.Empty{}, nil
}
