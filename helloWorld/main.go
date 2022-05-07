package main

import (
	"fmt"
	"log"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type Hello struct{ Who string }
type HelloActor struct {
	name string
}

func main() {
	fmt.Printf("> Stop program by hitting [ENTER]\n\n")
	actorSystem := actor.NewActorSystem()
	props := actor.PropsFromProducer(func() actor.Actor {
		return &HelloActor{}
	})
	pid, err := actorSystem.Root.SpawnNamed(props, "Hello World Actor")
	if err != nil {
		log.Fatal(err.Error())
	}
	actorSystem.Root.Send(pid, &Hello{Who: "Thomas"})

	// halt program till enter is pressed
	fmt.Scanln()
}

func (state *HelloActor) Receive(context actor.Context) {
	// Print message type
	fmt.Printf("Type of message: %T\n", context.Message())
	// Go type switch to distinguish between messages
	switch msg := context.Message().(type) {
	// React to custom message Hello
	case *Hello:
		fmt.Printf("Hello %v\n", msg.Who)
	// All other messages do not trigger any functionality within this actor
	default:
		fmt.Println("Received unknown message.")
	}
}
