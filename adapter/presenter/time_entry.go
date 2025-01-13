package presenter

import (
	"fmt"
	"os"
	"sort"
	"time"
	"time-tracker/domain/dto"
)

type TimeEntryPresenter struct{}

func NewTimeEntryPresenter() *TimeEntryPresenter {
	return &TimeEntryPresenter{}
}

func (p *TimeEntryPresenter) ShowError(err error) {
	fmt.Printf("❌ Error: %v\n", err)
	os.Exit(1)
}

func (p *TimeEntryPresenter) ShowTimeEntries(entries []dto.TimeEntryWithProjectAndTaskName, from, to *time.Time) {
	if from != nil && to != nil {
		fmt.Printf("Period: from %s to %s\n",
			from.Format("2006-01-02 15:04"),
			to.Format("2006-01-02 15:04"))
	} else if from != nil {
		fmt.Printf("Period: from %s\n", from.Format("2006-01-02 15:04"))
	} else if to != nil {
		fmt.Printf("Period: to %s\n", to.Format("2006-01-02 15:04"))
	} else {
		fmt.Println("Period: all")
	}

	if len(entries) == 0 {
		fmt.Print("No time entries found.")
		return
	}

	// 開始日時でソート
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].StartedAt.Before(entries[j].StartedAt)
	})

	for _, entry := range entries {
		fmt.Printf("[%s] %s: %s ~ %s (%v)\n",
			entry.ProjectName, entry.TaskName,
			entry.StartedAt.Format("2006-01-02 15:04"),
			entry.EndedAt.Format("2006-01-02 15:04"),
			entry.Duration)
	}
}
