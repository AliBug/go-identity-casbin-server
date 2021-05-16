package grpcserver

import (
	"context"

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

func (s *casbinService) HasPermissionForUserInDomain(ctx context.Context, r *pb.PermissionRequest) (*pb.BoolReply, error) {
	result, err := s.permissionsUC.HasPermissionForUserInDomain(r)
	if err != nil {
		return nil, err
	}
	return &pb.BoolReply{Res: result}, nil
}
