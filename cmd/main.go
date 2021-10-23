package main

import (
	"context"
	"github.com/mayerkv/go-recruitments/recruitment-service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	recruitment_service.RegisterRecruitmentServiceServer(grpcServer, newRecruitmentServiceServer())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

type recruitmentServiceServer struct {
	recruitment_service.UnimplementedRecruitmentServiceServer
}

func newRecruitmentServiceServer() recruitment_service.RecruitmentServiceServer {
	return &recruitmentServiceServer{}
}

func (s *recruitmentServiceServer) PostVacancy(ctx context.Context, request *recruitment_service.PostVacancyRequest) (*recruitment_service.PostVacancyResponse, error) {
	log.Println(request)

	return &recruitment_service.PostVacancyResponse{}, nil
}
