package main

import (
	"github.com/mayerkv/go-recruitmens/domain"
	"github.com/mayerkv/go-recruitmens/grpc-service"
	"github.com/mayerkv/go-recruitmens/repository"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	vacancyRepository := repository.NewInMemoryVacancyRepository()
	recruitmentRepository := repository.NewInMemoryRecruitmentRepository()
	recruitmentService := domain.NewRecruitmentService(recruitmentRepository, vacancyRepository)
	vacancyService := domain.NewVacancyService(vacancyRepository)
	srv := grpc_service.NewRecruitmentServiceServerImpl(recruitmentService, vacancyService)

	if err := runGrpcServer(srv); err != nil {
		log.Fatal(err)
	}
}

func runGrpcServer(srv grpc_service.RecruitmentServiceServer) error {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	grpc_service.RegisterRecruitmentServiceServer(grpcServer, srv)

	return grpcServer.Serve(lis)
}
