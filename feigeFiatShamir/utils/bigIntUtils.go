package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"math/big"
	"strings"
)

func VerifyRelativelyPrime(p, N big.Int) bool {
	zero := big.NewInt(0)
	if p.Cmp(zero) == 0 {
		return false
	}

	for N.Cmp(zero) != 0 {
		rest := big.NewInt(0).Mod(&p, &N)
		p = N
		N = *rest
	}
	// p now holds the greatest common divisor
	// if greatest common divisor is 1 p and n are relatively prime
	one := big.NewInt(1)
	if p.Cmp(one) == 0 {
		return true
	}
	return false
}

// Generate a secure random prime number of a given bit length.
func GenSecureRandPrime(bits int) *big.Int {

	p, err := rand.Prime(rand.Reader, bits)
	if err != nil {
		log.Fatal("Could not generate a random prime number!", err)
	}
	return p
}

func GenSecureRandom(bits int) *big.Int {
	// big.Int interprets as big-endian unsigned int --> so max. possible value is all bits set to 1
	max := new(big.Int)
	max.SetString(strings.Repeat("1", bits), 2)

	randomInt, err := rand.Int(rand.Reader, max)
	if err != nil {
		log.Fatalf("Could not generate random integer between 0 and %v!", max)
	}
	return randomInt
}

func ToBase64Str(bigInt big.Int) string {
	return base64.StdEncoding.EncodeToString(bigInt.Bytes())
}
