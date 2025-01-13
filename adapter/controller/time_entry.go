package controller

import (
	"fmt"
	"time"
	"time-tracker/adapter/presenter"
	"time-tracker/usecase"
)

type TimeEntryController struct {
	usecase   *usecase.TimeEntryUsecase
	presenter *presenter.TimeEntryPresenter
}

func NewTimeEntryController(u *usecase.TimeEntryUsecase, p *presenter.TimeEntryPresenter) *TimeEntryController {
	return &TimeEntryController{
		usecase:   u,
		presenter: p,
	}
}

func (c *TimeEntryController) ListTimeEntries(from, to string) {
	var fromDate, toDate *time.Time

	if from != "" {
		parsedFrom, err := time.Parse("2006-01-02", from)
		if err != nil {
			c.presenter.ShowError(fmt.Errorf("❌ Invalid format for --from. Use YYYY-MM-DD."))
			return
		}
		fromDate = &parsedFrom
	}
	if to != "" {
		parsedTo, err := time.Parse("2006-01-02", to)
		if err != nil {
			c.presenter.ShowError(fmt.Errorf("❌ Invalid format for --to. Use YYYY-MM-DD."))
			return
		}
		toDate = &parsedTo
	}

	// 両方未指定なら今日の日付範囲を設定
	now := time.Now()
	todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	todayEnd := todayStart.Add(24 * time.Hour)
	if fromDate == nil && toDate == nil {
		fromDate = &todayStart
		toDate = &todayEnd
	}

	timeEntries, err := c.usecase.ListTimeEntries(fromDate, toDate)
	if err != nil {
		c.presenter.ShowError(err)
		return
	}
	c.presenter.ShowTimeEntries(timeEntries, fromDate, toDate)
}
