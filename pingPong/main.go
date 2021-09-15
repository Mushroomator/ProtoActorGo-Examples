package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type PingPongMsg string
type StartPing struct {
	Pid *actor.PID
}
type PingPongActor struct {
	Name     string
	Msg      PingPongMsg
	delayInS time.Duration
}

func main() {
	// Create actor system
	system := actor.NewActorSystem()
	// Create "Ping" actor
	propsPing := actor.PropsFromProducer(func() actor.Actor {
		return &PingPongActor{
			Name:     "Ping Actor",
			Msg:      "PING!",
			delayInS: 1,
		}
	})
	pidPing := system.Root.Spawn(propsPing)

	// create "Pong" actor
	propsPong := actor.PropsFromProducer(func() actor.Actor {
		return &PingPongActor{
			Name:     "Pong Actor",
			Msg:      "PONG!",
			delayInS: 1,
		}
	})
	pidPong := system.Root.Spawn(propsPong)

	// Kick off the ping-pong match!
	system.Root.Send(pidPing, &StartPing{Pid: pidPong})

	// keep program running till enter is hit
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println("> Exiting program...")
}

// Handler whenever a message is received/ fetched from mailbox
func (state *PingPongActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		fmt.Println("Started PING actor.")
	case *StartPing:
		fmt.Printf("%v: Starting PING-PONG messaging...\n", state.Name)
		ctx.Request(msg.Pid, state.Msg)
	case PingPongMsg:
		fmt.Printf("%v: %v\n", state.Name, msg)
		pid := ctx.Sender()
		time.AfterFunc(state.delayInS*time.Second, func() {
			// IMPORTANT! Use Request() otherwise sender will not be send properly
			ctx.Request(pid, state.Msg)
		})
	default:
		fmt.Println("Received unknown message.")
	}
}
