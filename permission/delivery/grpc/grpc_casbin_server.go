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

func (s *casbinService) HasPermissionForUser(ctx context.Context, req *pb.PermissionRequest) (*pb.BoolReply, error) {
	result, err := s.permissionsUC.HasPermissionForUserUC(req)
	if err != nil {
		log.Printf("😯 err:%v", err)
		return nil, err
	}
	log.Println("final result: ", result)
	return &pb.BoolReply{Res: result}, nil
}

func (s *casbinService) AddRoleForUserInDomain(ctx context.Context, req *pb.UserRoleInDomainRequest) (*pb.BoolReply, error) {
	result, err := s.permissionsUC.AddRoleForUserInDomainUC(req)
	if err != nil {
		return nil, err
	}
	return &pb.BoolReply{Res: result}, nil
}

func (s *casbinService) DeleteRoleForUserInDomain(ctx context.Context, req *pb.UserRoleInDomainRequest) (*pb.BoolReply, error) {
	result, err := s.permissionsUC.DeleteRoleForUserInDomainUC(req)
	if err != nil {
		return nil, err
	}
	return &pb.BoolReply{Res: result}, nil
}

func (s *casbinService) DeleteRolesForUserInDomain(ctx context.Context, req *pb.UserRoleInDomainRequest) (*pb.BoolReply, error) {
	result, err := s.permissionsUC.DeleteRolesForUserInDomainUC(req)
	if err != nil {
		return nil, err
	}
	return &pb.BoolReply{Res: result}, nil
}

func (s *casbinService) GetDomainsForUser(ctx context.Context, req *pb.UserRequest) (*pb.StringArrayReply, error) {
	domains, err := s.permissionsUC.GetDomainsForUserUC(req)
	if err != nil {
		return nil, err
	}
	return &pb.StringArrayReply{Data: domains}, nil
}

func (s *casbinService) GetRolesForUserInDomain(ctx context.Context, req *pb.UserRoleInDomainRequest) (*pb.StringArrayReply, error) {
	roles := s.permissionsUC.GetRolesForUserInDomainUC(req)
	return &pb.StringArrayReply{Data: roles}, nil
}

func (s *casbinService) GetRolesInDomainsForUser(ctx context.Context, req *pb.UserRequest) (*pb.MapStringArrayReply, error) {
	results, err := s.permissionsUC.GetRolesInDomainsForUserUC(req)
	if err != nil {
		return nil, err
	}
	var list = make(map[string]*pb.StringArrayReply)
	for key, v := range results {
		roles := &pb.StringArrayReply{
			Data: v,
		}
		list[key] = roles
	}
	return &pb.MapStringArrayReply{Data: list}, nil
}

func (s *casbinService) GetPolicies(ctx context.Context, req *pb.EmptyRequest) (*pb.StringArray2DReply, error) {
	results := s.permissionsUC.GetPolicies()
	return wrapPlainPolicy(results), nil
}

func wrapPlainPolicy(policy [][]string) *pb.StringArray2DReply {
	if len(policy) == 0 {
		return &pb.StringArray2DReply{}
	}

	policyReply := &pb.StringArray2DReply{}
	policyReply.Data = make([]*pb.StringArrayReply, len(policy))
	for e := range policy {
		policyReply.Data[e] = &pb.StringArrayReply{Data: policy[e]}
	}
	return policyReply
}
