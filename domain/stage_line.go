package domain

import (
	"errors"
)

var (
	ErrStageLineSettingsIsEmpty = errors.New("stage line settings is empty")
)

type StageLine struct {
	StageId  string
	Settings []StageLineSettings
	History  map[string]*StageLineItem
}

func CreateNewStageLine(settings []StageLineSettings) (*StageLine, error) {
	if len(settings) < 1 {
		return nil, ErrStageLineSettingsIsEmpty
	}

	firstStageSettings := settings[0]

	return &StageLine{
		StageId:  firstStageSettings.StageId,
		Settings: settings,
		History: map[string]*StageLineItem{
			firstStageSettings.StageId: createNewStageLineItem(firstStageSettings),
		},
	}, nil
}

func (sl *StageLine) TransferItems(settings []StageLineSettings) (*StageLine, error) {
	var newSettings = settings
	if len(newSettings) == 0 {
		newSettings = sl.Settings
	}

	if len(newSettings) < 1 {
		return nil, ErrStageLineSettingsIsEmpty
	}

	var itemsForNewSettings map[string]*StageLineItem

	for _, s := range settings {
		if i, ok := sl.History[s.StageId]; ok {
			newItem := i.clone()

			if newItem.IsStarted() && !newItem.IsFinished() {
				if s.DeadlineDuration > 0 {
					newItem.DeadlineDate = newItem.StartDate.Add(s.DeadlineDuration)
				}
				if s.ThresholdDuration > 0 {
					newItem.ThresholdDate = newItem.StartDate.Add(s.ThresholdDuration)
				}
			}

			itemsForNewSettings[newItem.StageId] = newItem
		}
	}

	stageId := getCurrentStageFrom(newSettings, itemsForNewSettings)

	sl.Finish()

	return &StageLine{
		StageId:  stageId,
		Settings: newSettings,
		History:  itemsForNewSettings,
	}, nil
}

func (sl *StageLine) HasNextStage() bool {
	for _, s := range sl.Settings {
		if _, ok := sl.History[s.StageId]; !ok {
			return true
		}
	}

	return false
}

func (sl *StageLine) GetNextStageSettings() (StageLineSettings, bool) {
	for _, s := range sl.Settings {
		if _, ok := sl.History[s.StageId]; !ok {
			return s, ok
		}
	}

	return StageLineSettings{}, false
}

func (sl *StageLine) ToNextStage() {
	if item, ok := sl.History[sl.StageId]; ok && !item.IsFinished() {
		item.Finish()
	}

	if nextStageSettings, ok := sl.GetNextStageSettings(); ok {
		nextItem := createNewStageLineItem(nextStageSettings)
		sl.History[nextItem.StageId] = nextItem
		sl.StageId = nextItem.StageId
	}
}

func (sl *StageLine) Finish() {
	if item, ok := sl.History[sl.StageId]; ok && !item.IsFinished() {
		item.Finish()
	}
}

func (sl *StageLine) Recover() {
	if currentItem, ok := sl.History[sl.StageId]; ok {
		currentItem.Recover()
	}
}

func getCurrentStageFrom(settings []StageLineSettings, items map[string]*StageLineItem) string {
	// первый начатый и не завершенный
	if item := getFirstStartedAndNotFinished(settings, items); item != nil {
		return item.StageId
	}

	// последний если все завершены
	if ok := allFinished(settings, items); ok {
		return settings[len(settings)-1].StageId
	}

	// первый не завершенный
	if item := getFirstNotFinished(settings, items); item != nil {
		return item.StageId
	}

	// последний завершенный
	if item := getLastFinished(settings, items); item != nil {
		return item.StageId
	}

	// первый без истории
	for _, s := range settings {
		if _, ok := items[s.StageId]; !ok {
			return s.StageId
		}
	}

	return ""
}

func getFirstStartedAndNotFinished(settings []StageLineSettings, items map[string]*StageLineItem) *StageLineItem {
	for _, s := range settings {
		if item, ok := items[s.StageId]; ok && item.IsStarted() && !item.IsFinished() {
			return item
		}
	}

	return nil
}

func allFinished(settings []StageLineSettings, items map[string]*StageLineItem) bool {
	for _, s := range settings {
		if item, ok := items[s.StageId]; !ok || !item.IsFinished() {
			return false
		}
	}

	return true
}

func getFirstNotFinished(settings []StageLineSettings, items map[string]*StageLineItem) *StageLineItem {
	for _, s := range settings {
		if item, ok := items[s.StageId]; ok && !item.IsFinished() {
			return item
		}
	}

	return nil
}

func getLastFinished(settings []StageLineSettings, items map[string]*StageLineItem) *StageLineItem {
	var lastFinished *StageLineItem

	for _, s := range settings {
		if item, ok := items[s.StageId]; ok && lastFinished.IsFinished() {
			lastFinished = item
		}
	}

	return lastFinished
}
