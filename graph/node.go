package graph

import (
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

type Node struct {
	ID         string
	Graph      *Graph
	InnerPipes map[string]*Pipe
}

// NewNode .. create a new node for a graph
func NewNode(graph *Graph) *Node {
	u := uuid.NewV4()
	node := &Node{
		ID:         u.String(),
		Graph:      graph,
		InnerPipes: make(map[string]*Pipe),
	}
	return node
}

func (node *Node) GetInnerPipeByID(id string) *Pipe {
	pipe, exist := node.InnerPipes[id]
	if !exist {
		return nil
	}
	return pipe
}

func (node *Node) AttachInnerPipe(pipe *Pipe) bool {
	if node.GetInnerPipeByID(pipe.ID) != nil {
		return false
	}
	pipe.AttachOutputNode(node)
	node.InnerPipes[pipe.ID] = pipe
	log.Infof("new inner pipe %s attached to node %s", pipe.ID, node.ID)
	return true
}

func (node *Node) ReceiveEvent(sourcePipe *Pipe, payload map[string]interface{}) {
	log.Debugf("received event from pipe %s in node %s", sourcePipe.ID, node.ID)
}

func (node *Node) Persist() {

}
