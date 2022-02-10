package ffs

import (
	"log"
	"math/big"
)

// Create new instance of FeigeFiatShamirResponder. This is the server instance authenticating
// clients on
func NewFeigeFiatShamirResponder() FeigeFiatShamirResponder {
	return FeigeFiatShamirResponder{}
}

// Called when a Feige-Fiat-Shamir Initiator wants to authenticate itself against a responder.
// The initiator will send k ∈ N verification values, an X value and N.
// The responder now needs to generate k elements aᵢ with 0 < i <= k where each one is randomly either a 0 or a 1.
func (ffsr *FeigeFiatShamirResponder) OnAuthenticateRequest(verifyVals BigIntArray, X, N *big.Int) AuthenticationResponse {
	ffsr.VerifyVals = verifyVals
	ffsr.N = N
	ffsr.X = X
	ffsr.ChosenSecrets = make([]int8, len(verifyVals))

	for i := 0; i < len(verifyVals); i++ {
		//randomVal, err := rand.Int(rand.Reader, big.NewInt(1))
		//if err != nil {
		//	log.Fatal("Could not generate random number \u2208 {0, 1}")
		//}
		//ffsr.ChosenSecrets[i] = int8(randomVal.Int64())
		ffsr.ChosenSecrets[i] = int8(1)
	}
	return NewAuthenticationResponse(ffsr.ChosenSecrets)
}

// Decides wether authentication is successful or not.
// Authentication is successful if y² = +-X * (vᵢ)ᵃᵢ mod N is true. Otherwise authentication failed.
func (ffsr *FeigeFiatShamirResponder) OnAuthenticationVerify(y *big.Int) AuthenticationVerifyResponse {
	isAuthenticated := false
	// y²
	ySquared := y.Mul(y, y).Mod(y, ffsr.N)

	product := big.NewInt(1)
	for _, v := range ffsr.VerifyVals {
		product.Mul(product, &v)
	}
	// +-X * (vᵢ)ᵃᵢ mod N
	minusX := big.NewInt(1).Mul(ffsr.X, big.NewInt(int64(-1)))
	plus := big.NewInt(1).Mod(big.NewInt(1).Mul(product, ffsr.X), ffsr.N)
	minus := big.NewInt(1).Mod(big.NewInt(1).Mul(product, minusX), ffsr.N)

	// access is granted if y² = +-X * (vᵢ)ᵃᵢ mod N,	 ∀i with 0 < i <= noSecretVals
	if ySquared.Cmp(plus) == 0 || ySquared.Cmp(minus) == 0 {
		isAuthenticated = true
		log.Printf("%.40s does match %.40s or %.40s! - Access Granted!\n", ySquared, plus, minus)
	} else {
		log.Printf("%.40s does not match %.40s or %.40s! - Access Denied!\n", ySquared, plus, minus)
	}
	return NewAuthenticationVerifyResponse(isAuthenticated)
}
