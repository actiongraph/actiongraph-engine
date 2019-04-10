package graph

import (
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

type Pipe struct {
	ID                 string
	SourceNode         *Node
	SourceEventEmitter *EventEmitter
	OutputNodes        map[string]*Node
}

func NewPipe(sourceNode *Node, sourceEmitter *EventEmitter) *Pipe {
	u := uuid.NewV4()

	pipe := &Pipe{
		ID:                 u.String(),
		SourceNode:         sourceNode,
		SourceEventEmitter: sourceEmitter,
		OutputNodes:        make(map[string]*Node),
	}
	return pipe
}

func (pipe *Pipe) HasNodeSource() bool {
	if pipe.SourceNode == nil {
		return false
	}
	return true
}

func (pipe *Pipe) HasEmitterSource() bool {
	if pipe.SourceEventEmitter == nil {
		return false
	}
	return true
}

func (pipe *Pipe) GetOutputNodeByID(id string) *Node {
	node, exist := pipe.OutputNodes[id]
	if !exist {
		return nil
	}
	return node
}

func (pipe *Pipe) EmitEvent(source *EventEmitter, payload map[string]interface{}) {
	log.Debugf("received event from %s", source.ID)
	for _, node := range pipe.OutputNodes {
		node.ReceiveEvent(pipe, payload)
	}
}

func (pipe *Pipe) AttachOutputNode(node *Node) bool {
	if pipe.GetOutputNodeByID(node.ID) != nil {
		log.Errorf("Output node %s already exist in the pipe %s", node.ID, pipe.ID)
		return false
	}
	pipe.OutputNodes[node.ID] = node
	log.Infof("New output node %s added to the pipe %s", node.ID, pipe.ID)
	return true
}

func (pipe *Pipe) Persist() {

}
