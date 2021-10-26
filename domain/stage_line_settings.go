package domain

import "time"

type StageLineSettings struct {
	StageId           string
	DeadlineDuration  time.Duration
	ThresholdDuration time.Duration
}
