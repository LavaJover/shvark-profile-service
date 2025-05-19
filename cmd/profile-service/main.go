package main

import (
	"fmt"
	"log"
	"net"

	"github.com/LavaJover/shvark-profile-service/internal/config"
	"github.com/LavaJover/shvark-profile-service/internal/delivery/grpcapi"
	"github.com/LavaJover/shvark-profile-service/internal/infrastructure/postgres"
	profilepb "github.com/LavaJover/shvark-profile-service/proto/gen"
	"google.golang.org/grpc"
)

func main(){
	// Init config
	cfg := config.MustLoad()

	// Init db
	db := postgres.MustInitDB(cfg)

	// Init default profile repository
	profileRepo := &postgres.DefaultProfileRepository{DB: db}

	// Init default profile usecase
	profileUsecase := &postgres.DefaultProfileUsecase{Repo: profileRepo}

	// Creating gRPC server
	grpcServer := grpc.NewServer()
	profileHandler := grpcapi.ProfileHandler{Uc: profileUsecase}

	profilepb.RegisterProfileServiceServer(grpcServer, &profileHandler)

	// Start
	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil{
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("gRPC server started on %s:%s\n", cfg.Host, cfg.Port)
	if err := grpcServer.Serve(lis); err != nil{
		log.Fatalf("failed to serve: %v\n", err)
	}
}