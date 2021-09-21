package lifecycleactor

import (
	"log"
	"os"

	"github.com/AsynkronIT/protoactor-go/actor"
)

// Type for actor
type Actor struct{}

// Message causing a panic
type PanicMsg struct{}

func (state *Actor) Receive(ctx actor.Context) {
	// create logger
	logger := log.New(os.Stdout, "<Actor>\t", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Printf("Received message of type %T\n", ctx.Message())
	switch ctx.Message().(type) {
	// All lifecycles states an actor can be in
	case *actor.Started:
		logger.Printf("[STARTED]\t Started actor %v: Do any initialization here. All custom messages will only start to be processed after this has finsished.\n", ctx.Self().Id)
	case *actor.Stopping:
		logger.Printf("[STOPPING]\t Stopping actor %v: actor is about to shutdown. Do any cleanup (release of ressources etc.) here. "+
			"No more custom messages are taken out of the mailbox at this point. Currently processing messages will still be processed and child actors will receive a stop message.\n", ctx.Self().Id)
	case *actor.Stopped:
		logger.Printf("[STOPPED]\t Stopped actor %v: actor and its children are stopped. The actor is now dead.\n", ctx.Self().Id)
	case *actor.Restarting:
		logger.Printf("[RESTARTING]\t Restarting actor %v: actor is about to restart\n", ctx.Self().Id)
	case *actor.Terminated:
		logger.Printf("[TERMINATED]\t Terminated actor %v: actor is now dead.\n", ctx.Self().Id)
	// Receive message that causes a panic
	case PanicMsg:
		// throw panic --> default supervisor strategy will cause the parent to restart this actor and "Restarting" will be triggered
		logger.Println("Oh no a panic occured in our code!")
		panic("any panic")
	}
}
