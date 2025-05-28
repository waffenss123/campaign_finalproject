package main

import (
	"campaign-services/config"
	blog "campaign-services/gen/go/blog"
	"campaign-services/repository"
	"campaign-services/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	config.InitDB()
	if config.DB == nil {
		log.Fatal("Failed to connect to database. DB is nil.")
	}

	blogRepo := repository.NewBlogRepository(config.DB)
	blogSvc := services.NewBlogService(blogRepo)

	grpcSrv := grpc.NewServer()
	blog.RegisterBlogServiceServer(grpcSrv, blogSvc)

	listenPort := ":50053"
	lis, err := net.Listen("tcp", listenPort)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", listenPort, err)
	}

	log.Printf("Blog Service running on %s", listenPort)

	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatalf("Failed to running: %v", err)
	}
}
