package domain

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrCannotDenyRecruitment     = errors.New("cannot deny recruitment")
	ErrCannotCloseRecruitment    = errors.New("cannot close recruitment")
	ErrCannotAcceptStage         = errors.New("cannot accept stage")
	ErrInvalidRequestedStage     = errors.New("requested stage does not correspond to the current")
	ErrCannotTransferRecruitment = errors.New("cannot transfer recruitment")
)

type RecruitmentStatus int

const (
	RecruitmentStatusInWork RecruitmentStatus = iota
	RecruitmentStatusClosed
	RecruitmentStatusDenied
	RecruitmentStatusTransferred
)

type Recruitment struct {
	Id            string
	Status        RecruitmentStatus
	CandidateId   string
	ResponsibleId string
	StageLine     *StageLine
	Vacancy       *Vacancy
	CreatedAt     time.Time
	RefuseReason  *RefuseReason
}

func (r *Recruitment) CanModify() bool {
	return r.Status == RecruitmentStatusInWork
}

func (r *Recruitment) Deny(reasonId string, comment string) error {
	if !r.canSwitchStatus(RecruitmentStatusDenied) {
		return ErrCannotDenyRecruitment
	}

	r.RefuseReason = &RefuseReason{
		ReasonId: reasonId,
		Comment:  comment,
	}

	r.StageLine.Finish()

	return nil
}

func (r *Recruitment) success() error {
	if !r.canSwitchStatus(RecruitmentStatusClosed) {
		return ErrCannotCloseRecruitment
	}

	r.Status = RecruitmentStatusClosed
	r.StageLine.Finish()

	return nil
}

func (r *Recruitment) AcceptCurrentStage(requestedStageId string) error {
	if r.Status != RecruitmentStatusInWork {
		return ErrCannotAcceptStage
	}

	if r.StageLine.StageId != requestedStageId {
		return ErrInvalidRequestedStage
	}

	if r.StageLine.HasNextStage() {
		r.StageLine.ToNextStage()
	} else {
		return r.success()
	}

	return nil
}

func (r *Recruitment) TransferTo(vacancy *Vacancy, settings []StageLineSettings) (*Recruitment, error) {
	if !r.canTransfer() || !vacancy.canTransferRecruitment() {
		return nil, ErrCannotTransferRecruitment
	}

	stageLine, err := r.StageLine.TransferItems(settings)
	if err != nil {
		return nil, err
	}

	r.Status = RecruitmentStatusTransferred

	return &Recruitment{
		Id:            uuid.NewString(),
		Status:        RecruitmentStatusInWork,
		CandidateId:   r.CandidateId,
		ResponsibleId: r.ResponsibleId,
		StageLine:     stageLine,
		Vacancy:       vacancy,
		CreatedAt:     time.Now(),
		RefuseReason:  nil,
	}, nil

}

func (r *Recruitment) canTransfer() bool {
	return r.Status == RecruitmentStatusInWork
}

func (r *Recruitment) canSwitchStatus(s RecruitmentStatus) bool {
	statusMap := map[RecruitmentStatus][]RecruitmentStatus{
		RecruitmentStatusInWork:      {RecruitmentStatusClosed, RecruitmentStatusDenied, RecruitmentStatusInWork, RecruitmentStatusTransferred},
		RecruitmentStatusClosed:      {},
		RecruitmentStatusDenied:      {},
		RecruitmentStatusTransferred: {},
	}

	allowedStatuses, ok := statusMap[r.Status]
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

type RefuseReason struct {
	ReasonId string
	Comment  string
}
