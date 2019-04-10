package graph

import (
	"fmt"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

type Graph struct {
	ID            string
	Container     *GraphContainer
	Name          string
	Nodes         map[string]*Node
	EventEmitters map[string]*EventEmitter
}

// NewGraph .. get a new graph struct
func NewGraph(container *GraphContainer, name string) *Graph {
	// generate the uuid for the graph
	u := uuid.NewV4()

	graph := &Graph{
		ID:            u.String(),
		Container:     container,
		Name:          name,
		Nodes:         make(map[string]*Node),
		EventEmitters: make(map[string]*EventEmitter),
	}
	return graph
}

func (graph *Graph) GetNameWithID() string {
	return fmt.Sprintf("%s[%s]", graph.Name, graph.ID)
}

func (graph *Graph) GetNodeByID(id string) *Node {
	node, exist := graph.Nodes[id]
	if !exist {
		return nil
	}
	return node
}

func (graph *Graph) GetEventEmitterByID(id string) *EventEmitter {
	eventEmitter, exist := graph.EventEmitters[id]
	if !exist {
		return nil
	}
	return eventEmitter
}

func (graph *Graph) AddNode(node *Node) bool {
	if graph.GetNodeByID(node.ID) != nil {
		log.Errorf("Node %s already exist in the graph %s", node.ID, graph.GetNameWithID())
		return false
	}
	graph.Nodes[node.ID] = node
	log.Infof("New node %s added to the graph %s", node.ID, graph.GetNameWithID())
	return true
}

func (graph *Graph) AddEventEmitter(eventEmitter *EventEmitter) bool {
	if graph.GetEventEmitterByID(eventEmitter.ID) != nil {
		log.Errorf("emitter %s already exist in the graph %s", eventEmitter.ID, graph.GetNameWithID())
		return false
	}
	graph.EventEmitters[eventEmitter.ID] = eventEmitter
	log.Infof("New emitter %s added to the graph %s", eventEmitter.ID, graph.GetNameWithID())
	return true
}

func (graph *Graph) Persist() {

}
