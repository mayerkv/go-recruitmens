package main

import (
	"context"
	"fmt"
	"github.com/mayerkv/go-recruitmens/grpc-service"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:9090", opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := grpc_service.NewRecruitmentServiceClient(conn)

	ctx := context.Background()

	vacancy, err := client.PostVacancy(ctx, &grpc_service.PostVacancyRequest{
		PositionId: "123",
		CustomerId: "321",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", vacancy.VacancyId)

	getVacancy, err := client.GetVacancy(ctx, &grpc_service.GetVacancyRequest{VacancyId: vacancy.VacancyId})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", getVacancy.Vacancy)

	vacancies, err := client.ShowVacancies(ctx, &grpc_service.ShowVacanciesRequest{
		CustomerId:     "321",
		Page:           3,
		Size:           10,
		OrderBy:        grpc_service.Vacancy_CREATED_AT,
		OrderDirection: grpc_service.OrderDirection_ASC,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", vacancies.List)
	fmt.Printf("%#v\n", vacancies.Count)

	searchVacancies, err := client.SearchVacancies(ctx, &grpc_service.SearchVacanciesRequest{
		Page:           1,
		Size:           10,
		OrderBy:        grpc_service.Vacancy_CREATED_AT,
		OrderDirection: grpc_service.OrderDirection_ASC,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", searchVacancies.List)
	fmt.Printf("%#v\n", searchVacancies.Count)

	_, err = client.ApproveVacancy(ctx, &grpc_service.ApproveVacancyRequest{VacancyId: vacancy.VacancyId})
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.TakeInWorkVacancy(ctx, &grpc_service.TakeInWorkVacancyRequest{VacancyId: vacancy.VacancyId})
	if err != nil {
		log.Fatal(err)
	}

	considerCandidateResponse, err := client.ConsiderCandidate(ctx, &grpc_service.ConsiderCandidateRequest{
		VacancyId:     vacancy.VacancyId,
		CandidateId:   "123",
		ResponsibleId: "321",
		Settings: []*grpc_service.StageLineSettings{
			{
				StageId:              "1",
				DeadlineDurationSec:  2400,
				ThresholdDurationSec: 3600,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", considerCandidateResponse.RecruitmentId)
}
