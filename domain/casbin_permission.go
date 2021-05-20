package domain

// PermissionRequest ...
type PermissionRequest interface {
	GetUser() string
	GetPermissions() []string
}

// PermissionUseCase -
type PermissionUseCase interface {
	HasPermissionForUser(PermissionRequest) (bool, error)
	LoadPolicy() error
}
