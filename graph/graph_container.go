package graph

type GraphContainer struct {
	Graphs []*Graph
}

func NewGraphContainer() *GraphContainer {
	container := &GraphContainer{}
	return container
}

func (container *GraphContainer) AddGraph(graph *Graph) *Graph {
	container.Graphs = append(container.Graphs, graph)
	return graph
}

func (container *GraphContainer) Persist() {

}
