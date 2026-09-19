package svgbuilder

func GetCommitColor(commitsCount int) string {
	switch {
	case commitsCount >= 4:
		return "#56d364"
	case commitsCount == 3:
		return "#2ea043"
	case commitsCount == 2:
		return "#196c2e"
	case commitsCount == 1:
		return "#033a16"
	default:
		return "#151b23"
	}
}