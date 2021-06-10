package grpcserver

import (
	"context"
	"log"

	pb "github.com/alibug/go-identity-casbin-server/gen/casbin/proto"
	"github.com/casbin/casbin/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type casbinServer struct {
	enforcer *casbin.Enforcer
}

// NewCasbinServer -
func NewCasbinServer(e *casbin.Enforcer) pb.CasbinServer {
	return &casbinServer{enforcer: e}
}

// 必须修改!
/*
func (s *casbinServer) HasPermissionForUser(ctx context.Context, req *pb.PermissionRequest) (*pb.BoolReply, error) {
	roles, err := s.enforcer.GetImplicitRolesForUser(req.GetUser(), req.GetPermissions()[0])

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	for _, role := range roles {
		result := s.enforcer.HasPermissionForUser(role, req.GetPermissions()...)
		if result {
			return &pb.BoolReply{Res: true}, nil
		}
	}

	return &pb.BoolReply{Res: false}, nil
}
*/

func (s *casbinServer) HasPermissionForUser(ctx context.Context, req *pb.PermissionRequest) (*pb.BoolReply, error) {
	// 1、组成参数的数组
	parameters := append([]string{req.GetUser()}, req.GetPermissions()...)

	// 2、转换类型
	rvals := make([]interface{}, len(parameters))

	for i := range rvals {
		rvals[i] = parameters[i]
	}

	log.Println("rvals: ", rvals)

	// 2、计算 enforce结果
	result, err := s.enforcer.Enforce(rvals...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.BoolReply{Res: result}, nil
}

func (s *casbinServer) AddRoleForUserInDomain(ctx context.Context, req *pb.UserRoleInDomainRequest) (*pb.BoolReply, error) {
	result, err := s.enforcer.AddRoleForUserInDomain(req.GetUser(), req.GetRole(), req.GetDomain())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	err = s.enforcer.LoadPolicy()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.BoolReply{Res: result}, nil
}

func (s *casbinServer) DeleteRoleForUserInDomain(ctx context.Context, req *pb.UserRoleInDomainRequest) (*pb.BoolReply, error) {
	result, err := s.enforcer.DeleteRoleForUserInDomain(req.GetUser(), req.GetRole(), req.GetDomain())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	err = s.enforcer.LoadPolicy()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.BoolReply{Res: result}, nil
}

func (s *casbinServer) DeleteRolesForUserInDomain(ctx context.Context, req *pb.UserRoleInDomainRequest) (*pb.BoolReply, error) {
	result, err := s.enforcer.DeleteRolesForUserInDomain(req.GetUser(), req.GetDomain())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	err = s.enforcer.LoadPolicy()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.BoolReply{Res: result}, nil
}

func (s *casbinServer) GetDomainsForUser(ctx context.Context, req *pb.UserRoleInDomainRequest) (*pb.ArrayReply, error) {
	domains, err := s.enforcer.GetDomainsForUser(req.GetUser())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.ArrayReply{Data: domains}, nil
}

func (s *casbinServer) GetRolesForUserInDomain(ctx context.Context, req *pb.UserRoleInDomainRequest) (*pb.ArrayReply, error) {
	roles := s.enforcer.GetRolesForUserInDomain(req.GetUser(), req.GetDomain())
	return &pb.ArrayReply{Data: roles}, nil
}

func (s *casbinServer) GetRolesInDomainsForUser(ctx context.Context, req *pb.UserRoleInDomainRequest) (*pb.MapArrayReply, error) {
	domains, err := s.enforcer.GetDomainsForUser(req.GetUser())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	dict := make(map[string]*pb.ArrayReply, len(domains))
	for _, domain := range domains {

		roles := s.enforcer.GetRolesForUserInDomain(req.GetUser(), domain)
		dict[domain] = &pb.ArrayReply{
			Data: roles,
		}
	}

	return &pb.MapArrayReply{Data: dict}, nil
}

func (s *casbinServer) GetNamedPolicy(ctx context.Context, req *pb.PolicyRequest) (*pb.Array2DReply, error) {
	return wrapPlainPolicy(s.enforcer.GetNamedPolicy(req.PType)), nil
}

func (s *casbinServer) GetFilteredNamedPolicy(ctx context.Context, req *pb.FilteredPolicyRequest) (*pb.Array2DReply, error) {
	return wrapPlainPolicy(s.enforcer.GetFilteredNamedPolicy(req.GetPType(), int(req.GetFieldIndex()), req.GetFieldValues()...)), nil
}

func (s *casbinServer) AddNamedPolicy(ctx context.Context, req *pb.PolicyRequest) (*pb.BoolReply, error) {
	result, err := s.enforcer.AddNamedPolicy(req.GetPType(), req.GetParams())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	// ⚠️ 凡是修改型的 操作， LoadPolicy 不可少
	err = s.enforcer.LoadPolicy()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.BoolReply{Res: result}, nil
}

func (s *casbinServer) AddPolicy(ctx context.Context, req *pb.PolicyRequest) (*pb.BoolReply, error) {
	result, err := s.enforcer.AddPolicy(req.GetParams())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	// ⚠️ 凡是修改型的 操作， LoadPolicy 不可少
	err = s.enforcer.LoadPolicy()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.BoolReply{Res: result}, nil
}

func (s *casbinServer) RemoveNamedPolicy(ctx context.Context, req *pb.PolicyRequest) (*pb.BoolReply, error) {
	result, err := s.enforcer.RemoveNamedPolicy(req.GetPType(), req.GetParams())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	err = s.enforcer.LoadPolicy()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.BoolReply{Res: result}, nil
}

func (s *casbinServer) RemovePolicy(ctx context.Context, req *pb.PolicyRequest) (*pb.BoolReply, error) {
	result, err := s.enforcer.RemovePolicy(req.GetParams())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	err = s.enforcer.LoadPolicy()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.BoolReply{Res: result}, nil
}

func (s *casbinServer) DeleteUser(ctx context.Context, req *pb.UserRoleInDomainRequest) (*pb.BoolReply, error) {
	result, err := s.enforcer.DeleteUser(req.GetUser())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	err = s.enforcer.LoadPolicy()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.BoolReply{Res: result}, nil
}

func wrapPlainPolicy(policy [][]string) *pb.Array2DReply {
	if len(policy) == 0 {
		return &pb.Array2DReply{}
	}

	policyReply := &pb.Array2DReply{}
	policyReply.Data = make([]*pb.ArrayReply, len(policy))
	for e := range policy {
		policyReply.Data[e] = &pb.ArrayReply{Data: policy[e]}
	}
	return policyReply
}
