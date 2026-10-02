package data

type Day struct {
	colors  []string
	commits []*Commit
}

type Week struct {
	days [7][]*Day
}

type MonthLabel struct {
	name       string
	weekNumber int
}
