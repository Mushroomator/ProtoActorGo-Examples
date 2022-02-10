package ffs

import (
	"math/big"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type InitAuthenticationWithServer struct {
	Server actor.PID
}

func NewInitAuthenticationWithServer(server actor.PID) InitAuthenticationWithServer {
	return InitAuthenticationWithServer{
		Server: server,
	}
}

type AuthenticationRequest struct {
	VerifyVals BigIntArray
	X          *big.Int
	N          *big.Int
}

func NewAuthenticationRequest(verifyVals BigIntArray, x, n *big.Int) AuthenticationRequest {
	return AuthenticationRequest{
		VerifyVals: verifyVals,
		X:          x,
		N:          n,
	}
}

type AuthenticationResponse struct {
	ChosenSecrets []int8
}

func NewAuthenticationResponse(chosenSecrets []int8) AuthenticationResponse {
	return AuthenticationResponse{
		ChosenSecrets: chosenSecrets,
	}
}

type AuthenticationVerifyRequest struct {
	Y *big.Int
}

func NewAuthenticationVerifyRequest(y *big.Int) AuthenticationVerifyRequest {
	return AuthenticationVerifyRequest{
		Y: y,
	}
}

type AuthenticationVerifyResponse struct {
	IsAccessGranted bool
}

func NewAuthenticationVerifyResponse(isAccessGranted bool) AuthenticationVerifyResponse {
	return AuthenticationVerifyResponse{
		IsAccessGranted: isAccessGranted,
	}
}
