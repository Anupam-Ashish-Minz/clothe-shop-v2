package sharedtypes

type Graph struct {
	Labels []string
	Data   []int
}

type AdminGraphs struct {
	OrderCount Graph
}
