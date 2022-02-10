package ffs

import "math/big"

type BigIntArray []big.Int

type FeigeFiatShamirInitiator struct {
	// number of bits the secret consists of (should be >= 2048 for good security)
	// As this is only for demonstration purposes and not very secure anyway it can be much smaller e.g. 10
	noSecretBits int
	// number of secret values/ verification values (the more values, the securer the system). 10 shold be enough here
	noSecretVals int
	// prime number
	p *big.Int
	// prime number
	q *big.Int
	// N = p * q
	N *big.Int
	// random number (protection against replay attacks)
	// Calculated for each authentication attempt
	r *big.Int
	// sign: either -1 or 1
	// Calculated for each authentication attempt
	s int8
	// Calculated for each authentication attempt
	// y = r * (sᵢ)ᵃᵢ
	y *big.Int
	// Calculated for each authentication attempt
	// X = s * r² mod N
	X *big.Int
	// chosen secret values (those should under no circumstances by given to anybody else/ transmitted etc.)
	// Identical for each authentication process.
	// sᵢ with GCD(sᵢ, N) = 1 for 0 < i <= noSecretVals
	secretVals BigIntArray
	// verification values: used for someone to verify that initiator has the secrets (are trasnmitted on a public channel)
	// vᵢ = sᵢ² mod N for 0 < i <= noSecretVals
	VerifyVals BigIntArray
}

type FeigeFiatShamirResponder struct {
	X             *big.Int
	ChosenSecrets []int8
	N             *big.Int
	// verification values: used for someone to verify that initiator has the secrets (are trasnmitted on a public channel)
	VerifyVals BigIntArray
}
