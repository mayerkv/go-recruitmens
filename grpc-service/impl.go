package grpc_service

import (
	"context"
	"github.com/mayerkv/go-recruitmens/domain"
)

type RecruitmentServiceServerImpl struct {
	recruitmentService *domain.RecruitmentService
	vacancyService     *domain.VacancyService
}

func NewRecruitmentServiceServerImpl(recruitmentService *domain.RecruitmentService, vacancyService *domain.VacancyService) *RecruitmentServiceServerImpl {
	return &RecruitmentServiceServerImpl{recruitmentService: recruitmentService, vacancyService: vacancyService}
}

func (s *RecruitmentServiceServerImpl) PostVacancy(ctx context.Context, request *PostVacancyRequest) (*PostVacancyResponse, error) {
	vacancy, err := s.vacancyService.PostVacancy(request.PositionId, request.CustomerId)
	if err != nil {
		return nil, err
	}

	return &PostVacancyResponse{
		VacancyId: vacancy.Id,
	}, nil
}

func (s *RecruitmentServiceServerImpl) ShowVacancies(ctx context.Context, request *ShowVacanciesRequest) (*ShowVacanciesResponse, error) {
	vacancyPage, err := s.vacancyService.ShowVacancies(request.CustomerId, domain.Pageable{
		Page:           int(request.Page),
		Size:           int(request.Size),
		OrderBy:        mapVacancyOrderBy(request.OrderBy),
		OrderPredicate: mapOrderDirection(request.OrderDirection),
	})
	if err != nil {
		return nil, err
	}

	return &ShowVacanciesResponse{
		List:  mapVacancies(vacancyPage.Items),
		Count: int32(vacancyPage.TotalCount),
	}, nil
}

func (s *RecruitmentServiceServerImpl) SearchVacancies(ctx context.Context, request *SearchVacanciesRequest) (*SearchVacanciesResponse, error) {
	vacancyPage, err := s.vacancyService.SearchVacancies(domain.Pageable{
		Page:           int(request.Page),
		Size:           int(request.Size),
		OrderBy:        mapVacancyOrderBy(request.OrderBy),
		OrderPredicate: mapOrderDirection(request.OrderDirection),
	})
	if err != nil {
		return nil, err
	}

	return &SearchVacanciesResponse{
		List:  mapVacancies(vacancyPage.Items),
		Count: int32(vacancyPage.TotalCount),
	}, nil
}

func (s *RecruitmentServiceServerImpl) ChangeVacancyPosition(ctx context.Context, request *ChangeVacancyPositionRequest) (*Empty, error) {
	return NewEmpty(), s.vacancyService.ChangeVacancyPosition(request.VacancyId, request.PositionId)
}

func (s *RecruitmentServiceServerImpl) ApproveVacancy(ctx context.Context, request *ApproveVacancyRequest) (*Empty, error) {
	return NewEmpty(), s.vacancyService.ApproveVacancy(request.VacancyId)
}

func (s *RecruitmentServiceServerImpl) CloseVacancy(ctx context.Context, request *CloseVacancyRequest) (*Empty, error) {
	return NewEmpty(), s.vacancyService.CloseVacancy(request.VacancyId)
}

func (s *RecruitmentServiceServerImpl) RejectVacancy(ctx context.Context, request *RejectVacancyRequest) (*Empty, error) {
	return NewEmpty(), s.vacancyService.RejectVacancy(request.VacancyId)
}

func (s *RecruitmentServiceServerImpl) TakeInWorkVacancy(ctx context.Context, request *TakeInWorkVacancyRequest) (*Empty, error) {
	return NewEmpty(), s.vacancyService.TakeInWorkVacancy(request.VacancyId)
}

func (s *RecruitmentServiceServerImpl) GetVacancy(ctx context.Context, request *GetVacancyRequest) (*GetVacancyResponse, error) {
	vacancy, err := s.vacancyService.GetVacancy(request.VacancyId)
	if err != nil {
		return nil, err
	}

	return &GetVacancyResponse{Vacancy: mapVacancy(vacancy)}, nil
}

func (s *RecruitmentServiceServerImpl) ConsiderCandidate(ctx context.Context, request *ConsiderCandidateRequest) (*ConsiderCandidateResponse, error) {
	recruitment, err := s.recruitmentService.ConsiderCandidate(request.VacancyId, request.CandidateId, request.ResponsibleId, mapSettings(request.Settings))
	if err != nil {
		return nil, err
	}

	return &ConsiderCandidateResponse{
		RecruitmentId: recruitment.Id,
	}, nil
}

func (s *RecruitmentServiceServerImpl) ConsiderCandidateAnotherVacancy(ctx context.Context, request *ConsiderCandidateAnotherVacancyRequest) (*ConsiderCandidateAnotherVacancyResponse, error) {
	recruitment, err := s.recruitmentService.ConsiderCandidateAnotherVacancy(request.RecruitmentId, request.VacancyId, mapSettings(request.Settings))
	if err != nil {
		return nil, err
	}

	return &ConsiderCandidateAnotherVacancyResponse{
		RecruitmentId: recruitment.Id,
	}, nil
}

func (s *RecruitmentServiceServerImpl) AcceptRecruitmentStage(ctx context.Context, request *AcceptRecruitmentStageRequest) (*Empty, error) {
	return NewEmpty(), s.recruitmentService.AcceptRecruitmentStage(request.RecruitmentId, request.RequestedStageId)
}

func (s *RecruitmentServiceServerImpl) DenyRecruitment(ctx context.Context, request *DenyRecruitmentRequest) (*Empty, error) {
	return NewEmpty(), s.recruitmentService.DenyRecruitment(request.RecruitmentId, request.Reason.ReasonId, request.Reason.Comment)
}

func (s *RecruitmentServiceServerImpl) GetRecruitment(ctx context.Context, request *GetRecruitmentRequest) (*GetRecruitmentResponse, error) {
	recruitment, err := s.recruitmentService.GetRecruitment(request.RecruitmentId)
	if err != nil {
		return nil, err
	}

	return &GetRecruitmentResponse{
		Recruitment: mapRecruitment(recruitment),
	}, nil
}

func (s *RecruitmentServiceServerImpl) ShowRecruitments(ctx context.Context, request *ShowRecruitmentRequest) (*ShowRecruitmentsResponse, error) {
	recruitmentPage, err := s.recruitmentService.ShowRecruitments(request.ResponsibleId, domain.Pageable{
		Page:           int(request.Page),
		Size:           int(request.Size),
		OrderBy:        mapRecruitmentOrderBy(request.OrderBy),
		OrderPredicate: mapOrderDirection(request.OrderDirection),
	})
	if err != nil {
		return nil, err
	}

	return &ShowRecruitmentsResponse{
		List:  mapRecruitments(recruitmentPage.Items),
		Count: int32(recruitmentPage.TotalCount),
	}, nil
}

func (s *RecruitmentServiceServerImpl) mustEmbedUnimplementedRecruitmentServiceServer() {
	panic("implement me")
}
