package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/Mushroomator/ProtoActorGo-Examples/remoting/messages"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
)

func main() {
	system := actor.NewActorSystem()
	kind := remote.NewKind("counter", actor.PropsFromProducer(func() actor.Actor {
		return NewCounterActor()
	}))
	options := remote.Configure("127.0.0.1", 8081, remote.WithKinds(kind))
	remoter := remote.NewRemote(system, options)
	// start gRPC server
	remoter.Start()

	finish := make(chan os.Signal, 1)
	signal.Notify(finish, os.Interrupt, os.Kill)
	<-finish
	fmt.Println("Shutting down node...")
	remoter.Shutdown(false)
}

type CounterActor struct {
	counter int
}

func NewCounterActor() *CounterActor {
	return &CounterActor{
		counter: 0,
	}
}

func (state *CounterActor) handleShutdown(c actor.Context) {
	fmt.Printf("Actor %v counted up to %v\n", c.Self().String(), state.counter)
}

func (state *CounterActor) Receive(c actor.Context) {
	switch c.Message().(type) {
	case *actor.Started:
		fmt.Printf("Started actor %v\n", c.Self().String())
	case *messages.Count:
		state.counter++
	case *actor.Stopping:
		state.handleShutdown(c)
	case *remote.EndpointTerminatedEvent:
		state.handleShutdown(c)
	case *actor.Terminated:
		state.handleShutdown(c)
	}
}
