package domain

// PermissionRequest ...
type PermissionRequest interface {
	GetUser() string
	GetPermissions() []string
}

// UserRoleInDomainRequest ...
type UserRoleInDomainRequest interface {
	GetUser() string
	GetDomain() string
	GetRole() string
}

// UserRequest ...
type UserRequest interface {
	GetUser() string
}

// DomaiRequest ...
type DomaiRequest interface {
	GetDomain() string
}

// FilteredPolicyRequest -
type FilteredPolicyRequest interface {
	GetPType() string
	GetFieldIndex() int32
	GetFieldValues() []string
}

// PermissionUseCase -
type PermissionUseCase interface {
	HasPermissionForUserUC(PermissionRequest) (bool, error)
	AddRoleForUserInDomainUC(UserRoleInDomainRequest) (bool, error)
	DeleteRoleForUserInDomainUC(UserRoleInDomainRequest) (bool, error)
	DeleteRolesForUserInDomainUC(UserRoleInDomainRequest) (bool, error)
	GetDomainsForUserUC(UserRequest) ([]string, error)
	GetRolesForUserInDomainUC(UserRoleInDomainRequest) []string
	GetRolesInDomainsForUserUC(UserRequest) (map[string][]string, error)

	// for roles manager
	GetPoliciesUC() [][]string
	// GetPoliciesInDomainUC(DomaiRequest) [][]string
	GetNamedFilteredPolicyUC(FilteredPolicyRequest) [][]string
}

/*
  rpc GetDomainsForUserUC(UserRequest) returns (StringsReply) {}
  rpc GetRolesForUserInDomain(UserRoleInDomainRequest) returns (StringsReply) {}
  rpc GetRolesInDomainsForUser(UserRequest) returns (RolesInDomains) {}
*/
