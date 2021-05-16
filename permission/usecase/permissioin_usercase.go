package usecase

import (
	"github.com/alibug/go-identity-casbin-server/domain"
	"github.com/casbin/casbin/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type permissionUsecase struct {
	enforcer *casbin.Enforcer
}

// NewPermissionUsecase -
func NewPermissionUsecase(enforcer *casbin.Enforcer) domain.PermissionUseCase {
	return &permissionUsecase{enforcer: enforcer}
}

/*
func (r *permissionUsecase) HasPermissionForUser(user string, permissions ...string) bool {
	return r.enforcer.HasPermissionForUser(user, permissions...)
}
*/

func (uc *permissionUsecase) HasPermissionForUserInDomain(req domain.PermissionRequest) (bool, error) {
	if len(req.GetPermissions()) < 3 {
		return false, status.Errorf(codes.InvalidArgument, "Req permissions must 3")
	}

	roles, err := uc.enforcer.GetImplicitRolesForUser(req.GetUser(), req.GetPermissions()[0])
	if err != nil {
		return false, status.Errorf(codes.Internal, err.Error())
	}

	hasPermission := false

	for _, role := range roles {
		result := uc.enforcer.HasPermissionForUser(role, req.GetPermissions()...)
		if result {
			hasPermission = true
			break
		}
	}
	return hasPermission, nil
}
