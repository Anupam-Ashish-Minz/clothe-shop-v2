package sharedtypes

type Graph struct {
	Labels []string
	Data   []int
	Option string
}

type AdminGraphs struct {
	OrderCount    Graph
	RevenueAmount Graph
}
