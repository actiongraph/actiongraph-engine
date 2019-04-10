package app

import (
	"github.com/actiongraph/actiongraph-engine/graph"
	log "github.com/sirupsen/logrus"
)

type App struct {
	GraphContainers []*graph.GraphContainer
}

// NewApp .. get a new app
func NewApp() *App {
	app := &App{}
	return app
}

func (app *App) AddGraphContainer(container *graph.GraphContainer) *graph.GraphContainer {
	app.GraphContainers = append(app.GraphContainers, container)
	return container
}

// Start .. start the app
func (app *App) Start() {
	// Test
	// Create graph
	graphContainer := app.AddGraphContainer(graph.NewGraphContainer())
	g := graphContainer.AddGraph(graph.NewGraph(graphContainer, "test-graph"))

	// Create emitter
	emitter := graph.NewEventEmitter(g)
	g.AddEventEmitter(emitter)

	// Create node
	n := graph.NewNode(g)
	g.AddNode(n)

	// Attach pipe
	p := graph.NewPipe(n, emitter)
	n.AttachInnerPipe(p)
	emitter.AttachPipe(p)

	// Emit data from the EventEmitter
	emitter.EmitEvent(map[string]interface{}{
		"int_value": 5,
	})

	log.Info("app started")
}
