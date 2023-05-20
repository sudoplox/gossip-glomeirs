package main

import (
	"encoding/json"
	"log"
	"crypto/rand"
    "encoding/binary"
	"time"
    "fmt"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type server struct {
	n *maelstrom.Node
}

func main(){
    
	n := maelstrom.NewNode()
    s := &server{
        n:n,
    }
	n.Handle("generate", s.run)

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}

func (s *server) run(msg maelstrom.Message) error {
    var body map[string]interface{}
    if err := json.Unmarshal(msg.Body, &body); err != nil {
        return err
    }

    var randomNum int64
	err := binary.Read(rand.Reader, binary.BigEndian, &randomNum)
	if err != nil {
		return err
	}
    timeNow := time.Now().UnixNano()
    
    response := map[string]interface{}{
        "type": "generate_ok",
        "id": fmt.Sprintf("%v%v", timeNow, randomNum),
    }

    return s.n.Reply(msg, response)
}   