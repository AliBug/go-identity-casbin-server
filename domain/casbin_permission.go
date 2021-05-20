package domain

// PermissionRequest ...
type PermissionRequest interface {
	GetUser() string
	GetPermissions() []string
}

// UserDomainRoleRequest ...
type UserDomainRoleRequest interface {
	GetUser() string
	GetDomain() string
	GetRole() string
}

// PermissionUseCase -
type PermissionUseCase interface {
	HasPermissionForUser(PermissionRequest) (bool, error)
	LoadPolicy() error

	AddRoleForUserInDomain(UserDomainRoleRequest) (bool, error)
}
