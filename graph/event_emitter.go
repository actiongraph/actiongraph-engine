package graph

import (
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

type EventEmitter struct {
	ID    string
	Graph *Graph
	Pipes map[string]*Pipe
}

func NewEventEmitter(graph *Graph) *EventEmitter {
	// generate the uuid for the graph
	u := uuid.NewV4()

	eventEmitter := &EventEmitter{
		ID:    u.String(),
		Graph: graph,
		Pipes: make(map[string]*Pipe),
	}
	return eventEmitter
}

func (eventEmitter *EventEmitter) GetPipeByID(id string) *Pipe {
	pipe, exist := eventEmitter.Pipes[id]
	if !exist {
		return nil
	}
	return pipe
}

func (eventEmitter *EventEmitter) AttachPipe(pipe *Pipe) bool {
	if eventEmitter.GetPipeByID(pipe.ID) != nil {
		log.Errorf("Pipe %s already exist in the event emitter %s", pipe.ID, eventEmitter.ID)
		return false
	}
	eventEmitter.Pipes[pipe.ID] = pipe
	log.Infof("New pipe %s attached to the event emitter %s", pipe.ID, eventEmitter.ID)
	return true
}

func (eventEmitter *EventEmitter) EmitEvent(payload map[string]interface{}) {
	log.Debugf("event emitted from %s", eventEmitter.ID)
	for _, pipe := range eventEmitter.Pipes {
		log.Debugf("event emitted to pipe %s", pipe.ID)
		pipe.EmitEvent(eventEmitter, payload)
	}
}
