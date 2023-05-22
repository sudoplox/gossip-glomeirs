package main

import (
	"encoding/json"
	"log"
	"sync"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type server struct {
	n *maelstrom.Node

	idsMutex sync.RWMutex
	ids []float64 

	topologMutex sync.RWMutex
	currTopology map[string][]string
}

func main(){
    
	n := maelstrom.NewNode()
    s := &server{
        n:n,
    }
	n.Handle("broadcast", s.broadcastHandler)
	n.Handle("read", s.readHandler)
	n.Handle("topology", s.topologyHandler)

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}

func (s *server) broadcastHandler(msg maelstrom.Message) error {
    var body map[string]interface{}
    if err := json.Unmarshal(msg.Body, &body); err != nil {
        return err
    }
	s.idsMutex.Lock()
	s.ids = append(s.ids, body["message"].(float64))
	s.idsMutex.Unlock()

    response := map[string]interface{}{
        "type": "broadcast_ok",
    }

    return s.n.Reply(msg, response)
} 
func (s *server) readHandler(msg maelstrom.Message) error {
    var body map[string]interface{}
    if err := json.Unmarshal(msg.Body, &body); err != nil {
        return err
    }

	s.idsMutex.RLock()
	ids := s.ids
	s.idsMutex.RUnlock()

    response := map[string]interface{}{
        "type": "read_ok",
		"messages" : ids,
    }

    return s.n.Reply(msg, response)
}

type topologyMsg struct {
	Topology map[string][]string `json:"topology"`
}

func (s *server) topologyHandler(msg maelstrom.Message) error {
    var t topologyMsg
    if err := json.Unmarshal(msg.Body, &t); err != nil {
        return err
    }

	s.topologMutex.Lock()
	s.currTopology = t.Topology
	s.topologMutex.Unlock()


    response := map[string]interface{}{
        "type": "topology_ok",
    }

    return s.n.Reply(msg, response)
}   