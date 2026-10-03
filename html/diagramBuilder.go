package html

import (
	"time"

	"github.com/aguilam/git-activity/data"
	"github.com/aguilam/git-activity/svgbuilder"
)

func diagramBuilder(from time.Time, to time.Time, datedCommits map[string][]data.Commit) data.ActivityData {
	current := from
	var lastMonth int
	var lastWeekStart time.Time
	weekCounter := -1
	var activity data.ActivityData
	for current.Before(to) {
		currentWeekDay := int(current.Weekday())
		currentMonth := int(current.Month())

		date := current.Format("2006-01-02")
		commits := datedCommits[date]

		weekStart := current.AddDate(0, 0, -currentWeekDay)
		if lastWeekStart.IsZero() || !weekStart.Equal(lastWeekStart) {
			lastWeekStart = weekStart
			weekCounter++
			activity.Weeks = append(activity.Weeks, data.Week{})
		}

		nextWeekMonth := int(weekStart.AddDate(0, 0, 7).Month())
		if currentMonth != lastMonth && nextWeekMonth == currentMonth {
			lastMonth = currentMonth
			activity.Months = append(activity.Months, data.MonthLabel{Name: current.Format("Jan"),WeekNumber: weekCounter})
		}

		colors := svgbuilder.GetCommitColor(commits)

		activity.Weeks[weekCounter].Days[currentWeekDay] = &data.Day{Enabled: true,Colors: colors, Commits: &commits}

		current = current.AddDate(0, 0, 1)
	}
	return activity
}