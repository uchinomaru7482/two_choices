package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"two_choices/configs"
	"two_choices/interceptors"
	"two_choices/pkg/cache"
	"two_choices/pkg/firebase"
	"two_choices/pkg/rdb"
	"two_choices/src/interfaces"

	"google.golang.org/grpc"

	_ "github.com/go-sql-driver/mysql"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
)

type server struct{}

func main() {
	fmt.Println("run server")
	ctx := context.Background()
	// 設定情報
	appConfig := &configs.AppConfig{}
	appConfig.LoadAppConfig()
	// DB
	rdbExector := rdb.NewExector(appConfig.RDBConfig)
	// Firebase
	firebaseApp := firebase.NewFirebaseApp(appConfig.FirebaseConfig)
	// NewServer
	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			interceptors.AuthInterceptor(firebaseApp),
		),
	)
	// サービス登録
	interfaces.RegisterServices(ctx, s, appConfig, rdbExector, firebaseApp)
	// redis
	redis := cache.NewCachePool(appConfig.CacheConfig)
	redis.InitPool()
	cache.Redis = redis

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Register reflection service on gRPC server.
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
