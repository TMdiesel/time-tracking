package usecase

import (
	"time"
	"time-tracker/domain/dto"
	"time-tracker/domain/interfaces"
)

type TimeEntryUsecase struct {
	repo interfaces.ITimeEntryRepository
}

func NewTimeEntryUsecase(repo interfaces.ITimeEntryRepository) *TimeEntryUsecase {
	return &TimeEntryUsecase{repo: repo}
}

func (u *TimeEntryUsecase) ListTimeEntries(from, to *time.Time) ([]dto.TimeEntryWithProjectAndTaskName, error) {
	timeEntries, err := u.repo.ListTimeEntriesWithProjectAndTaskName(from, to)
	if err != nil {
		return nil, err
	}
	return timeEntries, err
}
