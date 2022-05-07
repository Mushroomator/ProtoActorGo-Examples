package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/asynkron/protoactor-go/actor"
)

type Continue struct{}
type Stop struct{}

type UnboundedNonDetActor struct {
	counter int
}

func main() {
	// create actor system
	system := actor.NewActorSystem()
	// spawn an actor
	actorProps := actor.PropsFromProducer(func() actor.Actor { return &UnboundedNonDetActor{} })
	pidUndActor := system.Root.Spawn(actorProps)

	sendMessage := func(msg interface{}) {
		system.Root.Send(pidUndActor, msg)
	}
	// send continue and stop message concurrently
	go sendMessage(Continue{})
	go sendMessage(Stop{})

	bufio.NewScanner(os.Stdin).Scan()
}

func (state *UnboundedNonDetActor) Receive(ctx actor.Context) {
	switch ctx.Message().(type) {
	case *actor.Started:
		fmt.Println("Actor has started...")
	case Continue:
		// increase counter
		state.counter++
		// sends itself a continue message
		ctx.Send(ctx.Self(), Continue{})
	case Stop:
		fmt.Printf("Actor has counted up to %v.\n", state.counter)
		// kills itself after this message has been processed
		ctx.Poison(ctx.Self())
	case *actor.Stopped:
		fmt.Println("Actor shutdown")
	}
}
