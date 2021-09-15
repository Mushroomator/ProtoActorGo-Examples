package main

import (
	"fmt"
	"log"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type Hello struct{ Who string }
type HelloActor struct{}

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
	// All lifecycles an actor can have
	case *actor.Started:
		fmt.Printf("> Started named actor %v: Do any initialization here. All custom messages will only start to be processed after this has finsished.\n", context.Self().Id)
		fmt.Println("> Started: done.")
	case *actor.Stopping:
		fmt.Printf("> Stopping named actor %v: actor is about to shutdown. Do any cleanup (release of ressources etc.) here\n", context.Self().Id)
	case *actor.Stopped:
		fmt.Printf("Stopped named actor %v: actor and its children are stopped. After this no more messages will be processed\n", context.Self().Id)
	case *actor.Restarting:
		fmt.Printf("> Restarting named actor %v: actor is about to restart", context.Self().Id)
	case *actor.Terminated:
		fmt.Printf("> Terminated named actor %v: actor is now dead.", context.Self().Id)

	// React to custom message Hello
	case *Hello:
		fmt.Printf("Hello %v\n", msg.Who)
	// All other messages do not trigger any functionality within this actor
	default:
		fmt.Println("Received unknown message.")
	}
}
