package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/Mushroomator/ProtoActorGo-Examples/remoting/messages"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	system := actor.NewActorSystem()
	options := remote.Configure("127.0.0.1", 8080)
	remoter := remote.NewRemote(system, options)
	// start gRPC server
	remoter.Start()
	// counter1, err := remoter.SpawnNamed("127.0.0.1:8081", "counter1", "counter", time.Second*10)
	// check(err)
	// fmt.Println("Actor gespawned")
	counter1, err := remoter.Spawn("127.0.0.1:8081", "counter", time.Second*10)
	check(err)
	counter2, err := remoter.Spawn("127.0.0.1:8081", "counter", time.Second*10)
	check(err)
	counter3, err := remoter.Spawn("127.0.0.1:8081", "counter", time.Second*10)
	check(err)

	rand.Seed(time.Now().UnixMicro())
	expectedCount1 := rand.Intn(10)
	for i := 0; i < expectedCount1; i++ {
		system.Root.Send(counter1.GetPid(), &messages.Count{})
	}
	fmt.Printf("Counter 1 should count up to %v.\n", expectedCount1)

	expectedCount2 := rand.Intn(20)
	for i := 0; i < expectedCount2; i++ {
		system.Root.Send(counter2.GetPid(), &messages.Count{})
	}
	fmt.Printf("Counter 2 should count up to %v.\n", expectedCount2)

	expectedCount3 := rand.Intn(50)
	for i := 0; i < expectedCount3; i++ {
		system.Root.Send(counter3.GetPid(), &messages.Count{})
	}
	fmt.Printf("Counter 3 should count up to %v.\n", expectedCount3)

	system.Root.Poison(counter1.GetPid())
	system.Root.Poison(counter2.GetPid())
	system.Root.Poison(counter3.GetPid())

	finish := make(chan os.Signal, 1)
	signal.Notify(finish, os.Interrupt, os.Kill)
	<-finish
	fmt.Println("Shutting down node...")
	remoter.Shutdown(false)

}
