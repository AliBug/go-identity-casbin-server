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

// LoadPolicy -
/*
func (uc *permissionUsecase) LoadPolicy() error {
	return uc.enforcer.LoadPolicy()
}
*/

// HasPermissionForUser -
func (uc *permissionUsecase) HasPermissionForUserUC(req domain.PermissionRequest) (bool, error) {
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

func (uc *permissionUsecase) AddRoleForUserInDomainUC(req domain.UserRoleInDomainRequest) (bool, error) {
	result, err := uc.enforcer.AddRoleForUserInDomain(req.GetUser(), req.GetRole(), req.GetDomain())
	if err != nil {
		return false, status.Errorf(codes.Internal, err.Error())
	}
	err = uc.enforcer.LoadPolicy()
	if err != nil {
		return false, status.Errorf(codes.Internal, err.Error())
	}

	return result, nil
}

func (uc *permissionUsecase) DeleteRoleForUserInDomainUC(req domain.UserRoleInDomainRequest) (bool, error) {
	result, err := uc.enforcer.DeleteRoleForUserInDomain(req.GetUser(), req.GetRole(), req.GetDomain())
	if err != nil {
		return false, status.Errorf(codes.Internal, err.Error())
	}
	err = uc.enforcer.LoadPolicy()
	if err != nil {
		return false, status.Errorf(codes.Internal, err.Error())
	}
	return result, nil
}

func (uc *permissionUsecase) DeleteRolesForUserInDomainUC(req domain.UserRoleInDomainRequest) (bool, error) {
	result, err := uc.enforcer.DeleteRolesForUserInDomain(req.GetUser(), req.GetDomain())
	if err != nil {
		return false, status.Errorf(codes.Internal, err.Error())
	}
	err = uc.enforcer.LoadPolicy()
	if err != nil {
		return false, status.Errorf(codes.Internal, err.Error())
	}
	return result, nil
}

func (uc *permissionUsecase) GetDomainsForUserUC(req domain.UserRequest) ([]string, error) {
	domains, err := uc.enforcer.GetDomainsForUser(req.GetUser())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return domains, nil
}

/*
  rpc GetDomainsForUserUC(UserRequest) returns (StringsReply) {}
  rpc GetRolesForUserInDomain(UserRoleInDomainRequest) returns (StringsReply) {}
  rpc GetRolesInDomainsForUser(UserRequest) returns (RolesInDomains) {}
*/
func (uc *permissionUsecase) GetRolesForUserInDomainUC(req domain.UserRoleInDomainRequest) []string {
	return uc.enforcer.GetRolesForUserInDomain(req.GetUser(), req.GetDomain())
}

func (uc *permissionUsecase) GetRolesInDomainsForUserUC(req domain.UserRequest) (map[string][]string, error) {
	domains, err := uc.enforcer.GetDomainsForUser(req.GetUser())
	if err != nil {
		return nil, err
	}
	result := make(map[string][]string)
	for _, domain := range domains {
		roles := uc.enforcer.GetRolesForUserInDomain(req.GetUser(), domain)
		result[domain] = roles
	}
	return result, nil
}

func (uc *permissionUsecase) GetPolicies() [][]string {
	return uc.enforcer.GetNamedPolicy("p")
}
