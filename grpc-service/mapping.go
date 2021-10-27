package grpc_service

import (
	"github.com/mayerkv/go-recruitmens/domain"
	"time"
)

func mapVacancyOrderBy(order Vacancy_Order) string {
	if order == Vacancy_CREATED_AT {
		return "createdAt"
	}

	return ""
}

func mapRecruitmentOrderBy(order Recruitment_Order) string {
	if order == Recruitment_CREATED_AT {
		return "createdAt"
	}

	return ""
}

func mapOrderDirection(direction OrderDirection) domain.Predicate {
	if direction == OrderDirection_DESC {
		return domain.PredicateDesc
	}

	return domain.PredicateAsc
}

func mapVacancies(vacancies []domain.Vacancy) []*Vacancy {
	var res []*Vacancy
	for _, item := range vacancies {
		res = append(res, mapVacancy(&item))
	}

	return res
}

func mapVacancy(vacancy *domain.Vacancy) *Vacancy {
	if vacancy == nil {
		return nil
	}

	return &Vacancy{
		Id:         vacancy.Id,
		PositionId: vacancy.PositionId,
		CustomerId: vacancy.CustomerId,
		CreatedAt:  vacancy.CreatedAt.Format(time.RFC3339),
		Status:     mapVacancyStatus(vacancy.Status),
	}
}

func mapVacancyStatus(status domain.VacancyStatus) Vacancy_Status {
	switch status {
	case domain.VacancyStatusPending:
		return Vacancy_PENDING
	case domain.VacancyStatusAgreed:
		return Vacancy_AGREED
	case domain.VacancyStatusRejected:
		return Vacancy_REJECTED
	case domain.VacancyStatusInWork:
		return Vacancy_IN_WORK
	case domain.VacancyStatusClosed:
		return Vacancy_CLOSED
	}

	return 0
}

func mapSettings(settings []*StageLineSettings) []domain.StageLineSettings {
	var res []domain.StageLineSettings
	for _, item := range settings {
		res = append(res, mapSetting(item))
	}

	return res
}

func mapSetting(item *StageLineSettings) domain.StageLineSettings {
	return domain.StageLineSettings{
		StageId:           item.StageId,
		DeadlineDuration:  time.Second * time.Duration(item.DeadlineDurationSec),
		ThresholdDuration: time.Second * time.Duration(item.ThresholdDurationSec),
	}
}

func mapRecruitment(recruitment *domain.Recruitment) *Recruitment {
	return &Recruitment{
		Id:            recruitment.Id,
		CandidateId:   recruitment.CandidateId,
		ResponsibleId: recruitment.ResponsibleId,
		CreatedAt:     recruitment.CreatedAt.Format(time.RFC3339),
		StageLine:     mapStageLine(recruitment.StageLine),
		Vacancy:       mapVacancy(recruitment.Vacancy),
		RefuseReason:  mapRefuseReason(recruitment.RefuseReason),
	}
}

func mapRefuseReason(reason *domain.RefuseReason) *RefuseReason {
	if reason == nil {
		return nil
	}

	return &RefuseReason{
		ReasonId: reason.ReasonId,
		Comment:  reason.Comment,
	}
}

func mapStageLine(line *domain.StageLine) *StageLine {
	if line == nil {
		return nil
	}

	return &StageLine{
		StageId:  line.StageId,
		Settings: mapStageLineSettings(line.Settings),
		History:  mapStageLineHistory(line.History),
	}
}

func mapStageLineHistory(history map[string]*domain.StageLineItem) map[string]*StageLineItem {
	var res map[string]*StageLineItem

	for id, item := range history {
		res[id] = mapStageLineItem(item)
	}

	return res
}

func mapStageLineItem(item *domain.StageLineItem) *StageLineItem {
	if item == nil {
		return nil
	}

	var startDate string
	if !item.StartDate.IsZero() {
		startDate = item.StartDate.Format(time.RFC3339)
	}

	var finishDate string
	if !item.FinishDate.IsZero() {
		finishDate = item.FinishDate.Format(time.RFC3339)
	}

	var deadlineDate string
	if !item.DeadlineDate.IsZero() {
		deadlineDate = item.DeadlineDate.Format(time.RFC3339)
	}

	var thresholdDate string
	if !item.FinishDate.IsZero() {
		thresholdDate = item.FinishDate.Format(time.RFC3339)
	}

	return &StageLineItem{
		StageId:       item.StageId,
		StartDate:     startDate,
		FinishDate:    finishDate,
		DeadlineDate:  deadlineDate,
		ThresholdDate: thresholdDate,
	}
}

func mapStageLineSettings(settings []domain.StageLineSettings) []*StageLineSettings {
	var res []*StageLineSettings

	for _, item := range settings {
		res = append(res, mapStageLineSetting(item))
	}

	return res
}

func mapStageLineSetting(item domain.StageLineSettings) *StageLineSettings {
	return &StageLineSettings{
		StageId:              item.StageId,
		DeadlineDurationSec:  int32(item.DeadlineDuration.Seconds()),
		ThresholdDurationSec: int32(item.ThresholdDuration.Seconds()),
	}
}

func mapRecruitments(items []domain.Recruitment) []*Recruitment {
	var res []*Recruitment
	for _, item := range items {
		res = append(res, mapRecruitment(&item))
	}

	return res
}
