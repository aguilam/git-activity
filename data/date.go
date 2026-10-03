package data

type Day struct {
	Enabled bool
	Colors  []string
	Commits *[]Commit
}

type Week struct {
	Days [7]*Day
}

type MonthLabel struct {
	Name       string
	WeekNumber int
}

type ActivityData struct {
	Months []MonthLabel
	Weeks  []Week
}
