package main

import (
	"fmt"
	"math/big"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/Mushroomator/ProtoActorGo-Examples/feigeFiatShamir/actors"
	"github.com/Mushroomator/ProtoActorGo-Examples/feigeFiatShamir/ffs"
)

func main() {
	system := actor.NewActorSystem()

	aliceProps := actor.PropsFromProducer(func() actor.Actor {
		return actors.NewAlice()
	})
	alicePid := system.Root.Spawn(aliceProps)

	victorProps := actor.PropsFromProducer(func() actor.Actor {
		return actors.NewVictor()
	})
	victorPid := system.Root.Spawn(victorProps)

	two := big.NewInt(2)
	three := big.NewInt(3)

	//result := big.NewInt(1)
	fmt.Println(two.Mul(two, three).Mul(two, three).String())

	// tell Alice (Feige-Fiat-Shamir Initiator) to authenticate with Victor (Feige-Fiat-Shamir Responder)
	system.Root.Send(alicePid, ffs.NewInitAuthenticationWithServer(*victorPid))

	fmt.Scanf("%s")
}
