package domain

// PermissionRequest ...
type PermissionRequest interface {
	GetUser() string
	GetPermissions() []string
}

// PermissionUseCase -
type PermissionUseCase interface {
	HasPermissionForUserInDomain(PermissionRequest) (bool, error)
}
