package main

import (
	"context"
	"github.com/mayerkv/go-recruitmens/grpc-service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	grpc_service.RegisterRecruitmentServiceServer(grpcServer, newRecruitmentServiceServer())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

type recruitmentServiceServer struct {
	grpc_service.UnimplementedRecruitmentServiceServer
}

func newRecruitmentServiceServer() grpc_service.RecruitmentServiceServer {
	return &recruitmentServiceServer{}
}

func (s *recruitmentServiceServer) PostVacancy(ctx context.Context, request *grpc_service.PostVacancyRequest) (*grpc_service.PostVacancyResponse, error) {
	log.Println(request.Message)

	return &grpc_service.PostVacancyResponse{
		Message: os.Getenv("HOSTNAME"),
	}, nil
}
