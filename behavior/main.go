package main

import (
	"fmt"
	"log"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type SetBehaviorActor struct {
	behavior actor.Behavior
}

type GreetingMsg struct {
	name string
}

func main() {
	// init actor system
	log.Println("Starting actor system")
	fmt.Println("> Stop program by pressing CTRL + C")
	system := actor.NewActorSystem()

	props := actor.PropsFromProducer(func() actor.Actor {
		// create new actor
		actor := &SetBehaviorActor{
			behavior: actor.NewBehavior(),
		}
		// set initial behavior
		actor.behavior.Become(actor.GreetEnglish)
		return actor
	})

	pid := system.Root.Spawn(props)

	// greet every second for eternity
	// language will be changed every time a greeting is issued
	// as there are 3 languages every 3 messages will be the same
	for range time.Tick(time.Second * 1) {
		system.Root.Send(pid, GreetingMsg{"Thomas"})
	}

}

// Default message handler method with delegates message handling
// to initial behavior defined when creating the actor
func (state *SetBehaviorActor) Receive(ctx actor.Context) {
	state.behavior.Receive(ctx)
}

// Greets a user in German and sets behavior to greet in Spanish next time
func (state *SetBehaviorActor) GreetEnglish(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case GreetingMsg:
		fmt.Printf("Hello %v\n", msg.name)
		state.behavior.Become(state.GreetGerman)
	}
}

// Greets a user in German and sets behavior to greet in Spanish next time
func (state *SetBehaviorActor) GreetGerman(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case GreetingMsg:
		fmt.Printf("Hallo %v\n", msg.name)
		state.behavior.Become(state.GreetSpanish)
	}
}

// Greets a user in Spanish and sets behavior to greet in English next time
func (state *SetBehaviorActor) GreetSpanish(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case GreetingMsg:
		fmt.Printf("Hola %v\n", msg.name)
		state.behavior.Become(state.GreetEnglish)
	}
}
