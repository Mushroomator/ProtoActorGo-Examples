package actors

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/Mushroomator/ProtoActorGo-Examples/feigeFiatShamir/ffs"
)

type Alice struct {
	ffsi ffs.FeigeFiatShamirInitiator
}

func NewAlice() *Alice {
	return &Alice{}
}

func (state *Alice) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		state.init()
	case ffs.InitAuthenticationWithServer:
		req := state.ffsi.InitiateAuthentication()
		ctx.Request(&msg.Server, req)
	case ffs.AuthenticationResponse:
		verifyReq := state.ffsi.OnAuthenticationResponse(msg.ChosenSecrets)
		ctx.Request(ctx.Sender(), verifyReq)
	case ffs.AuthenticationVerifyResponse:
		state.ffsi.OnAuthenticationVerifyResponse(msg.IsAccessGranted)
	}
}

// Initialize Alice (A). Alice is the initiator of the conversation using the Feige-Fiat-Shamir protocol.
// To initialize Alice needs to...
//  - choose two different prime numbers p and q
//  - calculate N = p * q and publish N (and keep p and q PRIVATE!)
//  - choose k secret values sᵢ ∀i with 0 < i <= noSecretVals
func (alice *Alice) init() {
	alice.ffsi = ffs.NewFeigeFiatShamirInitiator(
		128, // must be >= 5
		10,  // must be >= 1
	)
	alice.ffsi.Prepare()
}
