package usecase

import (
	"log"

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

func (uc *permissionUsecase) HasPermissionForUser(req domain.PermissionRequest) (bool, error) {
	// log.Println("req user: ", req.GetUser())
	// log.Println("req permis: ", req.GetPermissions())

	roles, err := uc.enforcer.GetImplicitRolesForUser(req.GetUser(), req.GetPermissions()[0])

	if err != nil {
		// log.Printf("roles err: %v", err)
		return false, status.Errorf(codes.Internal, err.Error())
	}

	// log.Println("roles: ", roles)

	// hasPermission := false

	for _, role := range roles {
		result := uc.enforcer.HasPermissionForUser(role, req.GetPermissions()...)
		if result {
			// hasPermission = true
			log.Println("got !")
			return true, nil
		}
	}
	log.Println("👮 fail!")
	return false, nil
}
