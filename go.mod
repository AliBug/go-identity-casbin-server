module github.com/alibug/go-identity-casbin-server

go 1.16

require (
	github.com/alibug/go-identity-utils v0.1.12
	github.com/casbin/casbin/v2 v2.31.2
	github.com/casbin/mongodb-adapter/v3 v3.2.0
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
)

replace github.com/casbin/mongodb-adapter/v3 => github.com/alibug/go-casbin-mongodb-adapter/v3 v3.2.3
