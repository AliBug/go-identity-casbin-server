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
	_permissionUseCase "github.com/alibug/go-identity-casbin-server/permission/usecase"
	_casbinAdapter "github.com/casbin/mongodb-adapter/v3"
)

func main() {
	log.Println("启动")
	mongoURL := config.ReadMongoConfig("mongo")

	timeDuration := 100 * time.Second

	adapter, err := _casbinAdapter.NewAdapter(mongoURL, timeDuration)

	if err != nil {
		log.Fatalf("Init mongo adapter fail: %s", err)
	}

	casbinModel := config.ReadCasbinFilePath("casbin")
	enforcer, err := casbin.NewEnforcer(casbinModel, adapter)

	if err != nil {
		log.Fatalf("Init Casbin enforcer fail: %s", err)
	}

	permissionUseCase := _permissionUseCase.NewPermissionUsecase(enforcer)

	casbinService := _grpcDelivery.NewCasbinService(permissionUseCase)

	server := grpc.NewServer()
	// pb.RegisterSearchServiceServer(server, &SearchService{})
	pb.RegisterCasbinServer(server, casbinService)

	lis, err := net.Listen("tcp", ":"+"9001")
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}
