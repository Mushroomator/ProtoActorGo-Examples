package actors

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/Mushroomator/ProtoActorGo-Examples/feigeFiatShamir/ffs"
)

type Victor struct {
	ffsr ffs.FeigeFiatShamirResponder
}

// Create Victor (V stands for Verfication). Victor is the Feige-Fiat-Shamir Responder and verifies that Alice knows/ does not know a secret.
func NewVictor() *Victor {
	return &Victor{}
}

func (state *Victor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		state.init()
	case ffs.AuthenticationRequest:
		resp := state.ffsr.OnAuthenticateRequest(msg.VerifyVals, msg.X, msg.N)
		ctx.Request(ctx.Sender(), resp)
	case ffs.AuthenticationVerifyRequest:
		resp := state.ffsr.OnAuthenticationVerify(msg.Y)
		ctx.Request(ctx.Sender(), resp)
	}
}

func (v *Victor) init() {
	v.ffsr = ffs.FeigeFiatShamirResponder{}
}
