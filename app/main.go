package main

import (
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/casbin/casbin/v2"

	"github.com/alibug/go-identity-utils/config"

	pb "github.com/alibug/go-identity-casbin-server/gen/casbin/proto"
	_grpcDelivery "github.com/alibug/go-identity-casbin-server/permission/delivery/grpc"

	// _permissionUseCase "github.com/alibug/go-identity-casbin-server/permission/usecase"
	_casbinAdapter "github.com/casbin/mongodb-adapter/v3"
)

func main() {
	log.Println("启动")
	mongoURL := config.ReadMongoConfig("mongo")

	duration := config.ReadCustomIntConfig("mongo.duration", false)

	timeDuration := time.Duration(duration) * time.Second

	adapter, err := _casbinAdapter.NewAdapter(mongoURL, timeDuration)

	if err != nil {
		log.Fatalf("Init mongo adapter fail: %s", err)
	}

	casbinModel := config.ReadCasbinFilePath("casbin")
	enforcer, err := casbin.NewEnforcer(casbinModel, adapter)

	if err != nil {
		log.Fatalf("Init Casbin enforcer fail: %s", err)
	}

	// permissionUseCase := _permissionUseCase.NewPermissionUsecase(enforcer)

	casbinService := _grpcDelivery.NewCasbinServer(enforcer)

	server := grpc.NewServer()

	pb.RegisterCasbinServer(server, casbinService)

	port := config.ReadCustomStringConfig("grpc.port")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("serve fail: %v", err)
	}
}
