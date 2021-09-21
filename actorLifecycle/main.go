package main

import (
	"bufio"
	"log"
	"os"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	lcActor "github.com/Mushroomator/ProtoActorGo-Examples/actorLifecycle/lifecycleActor"
)

func main() {
	// create logger
	logger := log.New(os.Stdout, "<Main>\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Create actor system
	system := actor.NewActorSystem()
	// Create "Ping" actor
	props := actor.PropsFromProducer(func() actor.Actor {
		return &lcActor.Actor{}
	})
	pid := system.Root.Spawn(props)
	// wait 5 seconds then cause a panic which will restart the actor actor
	time.AfterFunc(5*time.Second, func() {
		logger.Println("Causing a panic in actor.")
		system.Root.Send(pid, lcActor.PanicMsg{})
	})

	// wait 10 seconds then kill actor
	time.AfterFunc(10*time.Second, func() {
		// kills actor immediately (= no more messages of mailbox will be processed)
		logger.Println("Killing actor immediately.")
		system.Root.Stop(pid)

		// Important distinction! Poison will stop actor after all messages currently in the mailbox are processed!
		//logger.Println("Killing actor after it has processed all messages in its mailbox.")
		// system.Root.Poison(pid)
	})

	// keep program running till enter is hit
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	logger.Println("Exiting program...")
}
