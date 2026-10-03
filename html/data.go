package html

import "github.com/aguilam/git-activity/data"

type Activities struct {
	CurrentActivity data.ActivityData
	YearActivities []data.ActivityData
}