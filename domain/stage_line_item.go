package domain

import (
	"errors"
	"time"
)

var (
	ErrStageLineItemAlreadyFinished = errors.New("stage line item already finished")
)

type StageLineItem struct {
	StageId       string
	StartDate     time.Time
	FinishDate    time.Time
	DeadlineDate  time.Time
	ThresholdDate time.Time
}

func (i *StageLineItem) IsStarted() bool {
	return !i.StartDate.IsZero()
}

func (i *StageLineItem) IsFinished() bool {
	return !i.FinishDate.IsZero()
}

func (i *StageLineItem) Finish() error {
	if i.IsFinished() {
		return ErrStageLineItemAlreadyFinished
	}

	i.FinishDate = time.Now()

	return nil
}

func (i *StageLineItem) Recover() {
	i.FinishDate = time.Time{}
}

func (i *StageLineItem) clone() *StageLineItem {
	return &StageLineItem{
		StageId:       i.StageId,
		StartDate:     i.StartDate,
		FinishDate:    i.FinishDate,
		DeadlineDate:  i.DeadlineDate,
		ThresholdDate: i.ThresholdDate,
	}
}

func createNewStageLineItem(setting StageLineSettings) *StageLineItem {
	now := time.Now()

	var deadlineDate time.Time
	if setting.DeadlineDuration > 0 {
		deadlineDate = now.Add(setting.DeadlineDuration)
	}

	var thresholdDate time.Time
	if setting.ThresholdDuration > 0 {
		thresholdDate = now.Add(setting.ThresholdDuration)
	}

	return &StageLineItem{
		StageId:       setting.StageId,
		StartDate:     now,
		FinishDate:    time.Time{},
		DeadlineDate:  deadlineDate,
		ThresholdDate: thresholdDate,
	}
}
