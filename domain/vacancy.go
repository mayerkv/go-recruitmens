package domain

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrCannotApproveVacancy        = errors.New("cannot approve vacancy")
	ErrCannotCloseVacancy          = errors.New("cannot close vacancy")
	ErrCannotRejectVacancy         = errors.New("cannot reject vacancy")
	ErrCannotTakeInWorkVacancy     = errors.New("cannot take in work vacancy")
	ErrVacancyNotInWork            = errors.New("vacancy not in work")
	ErrCannotChangeVacancyPosition = errors.New("cannot change vacancy position")
)

type VacancyStatus int

const (
	VacancyStatusPending VacancyStatus = iota
	VacancyStatusAgreed
	VacancyStatusRejected
	VacancyStatusInWork
	VacancyStatusClosed
)

type Vacancy struct {
	Id         string
	PositionId string
	CustomerId string
	CreatedAt  time.Time
	Status     VacancyStatus
}

func CreateVacancy(positionId string, customerId string) *Vacancy {
	return &Vacancy{
		Id:         uuid.NewString(),
		PositionId: positionId,
		CustomerId: customerId,
		CreatedAt:  time.Now(),
		Status:     VacancyStatusPending,
	}
}

func (v *Vacancy) ChangePosition(positionId string) error {
	if v.Status != VacancyStatusPending {
		return ErrCannotChangeVacancyPosition
	}

	v.PositionId = positionId

	return nil
}

func (v *Vacancy) Approve() error {
	if !v.canSwitchTo(VacancyStatusAgreed) {
		return ErrCannotApproveVacancy
	}

	v.Status = VacancyStatusAgreed

	return nil
}

func (v *Vacancy) Close() error {
	if !v.canSwitchTo(VacancyStatusClosed) {
		return ErrCannotCloseVacancy
	}

	v.Status = VacancyStatusClosed

	return nil
}

func (v *Vacancy) Reject() error {
	if !v.canSwitchTo(VacancyStatusRejected) {
		return ErrCannotRejectVacancy
	}

	v.Status = VacancyStatusRejected

	return nil
}

func (v *Vacancy) TakeInWork() error {
	if !v.canSwitchTo(VacancyStatusInWork) {
		return ErrCannotTakeInWorkVacancy
	}

	v.Status = VacancyStatusInWork

	return nil
}

func (v *Vacancy) canSwitchTo(s VacancyStatus) bool {
	statusMap := map[VacancyStatus][]VacancyStatus{
		VacancyStatusPending:  {VacancyStatusAgreed, VacancyStatusRejected},
		VacancyStatusAgreed:   {VacancyStatusInWork, VacancyStatusClosed},
		VacancyStatusRejected: {},
		VacancyStatusInWork:   {VacancyStatusClosed},
		VacancyStatusClosed:   {},
	}

	allowedStatuses, ok := statusMap[v.Status]
	if !ok {
		return false
	}

	for _, as := range allowedStatuses {
		if as == s {
			return true
		}
	}

	return false
}

func (v *Vacancy) ConsiderCandidate(candidateId string, responsibleId string, settings []StageLineSettings) (*Recruitment, error) {
	if v.Status != VacancyStatusInWork {
		return nil, ErrVacancyNotInWork
	}

	stageLine, err := CreateNewStageLine(settings)
	if err != nil {
		return nil, err
	}

	return &Recruitment{
		Id:            uuid.NewString(),
		Status:        RecruitmentStatusInWork,
		CandidateId:   candidateId,
		ResponsibleId: responsibleId,
		StageLine:     stageLine,
		Vacancy:       v,
		CreatedAt:     time.Now(),
	}, nil
}

func (v *Vacancy) canTransferRecruitment() bool {
	return v.Status == VacancyStatusInWork
}
