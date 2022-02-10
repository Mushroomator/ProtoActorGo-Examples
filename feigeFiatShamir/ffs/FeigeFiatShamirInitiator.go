package ffs

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/Mushroomator/ProtoActorGo-Examples/feigeFiatShamir/utils"
)

func (m BigIntArray) String() string {
	var sb strings.Builder
	sb.WriteString("\n")
	line := strings.Repeat("-", 5+40+40+40+4*2+5*1) + "\n"
	sb.WriteString(line)
	sb.WriteString("| i     | sᵢ in base10 (first 40 digits)           | sᵢ in base64 (first 40 characters)       | sᵢ in binary (first 40 bits)             |\n")
	sb.WriteString(line)
	for i, val := range m {
		sb.WriteString(fmt.Sprintf("| %5d | %40.40s | %40.40s | %40.40s |", i, val.String(), utils.ToBase64Str(val), val.Text(2)))
		sb.WriteString("\n")
	}
	return sb.String()
}

// Create new Feige-Fiat-Shamir Initiator.
// noSecretBits must be >= 5
// noSecretVals must be >= 1
func NewFeigeFiatShamirInitiator(noSecretBits, noSecretVals int) FeigeFiatShamirInitiator {
	if noSecretBits < 5 {
		panic("Number of bits for the secret must at least be 5!")
	}
	if noSecretVals < 1 {
		panic("Number of secret values must at least be 1!")
	}
	return FeigeFiatShamirInitiator{
		noSecretBits: noSecretBits,
		noSecretVals: noSecretVals,
	}
}

func (ffsi *FeigeFiatShamirInitiator) Prepare() {
	ffsi.p, ffsi.q, ffsi.N = ffsi.CreatePrimes()
	// generate random secret values sᵢ: sᵢ with GCD(sᵢ, N) = 1 for 0 < i <= noSecretVals
	ffsi.secretVals = ffsi.GenerateSecretValues(ffsi.noSecretVals, ffsi.N)
	// calculate verification values based on secret values vᵢ: vᵢ = sᵢ² mod N for 0 < i <= noSecretVals
	ffsi.VerifyVals = ffsi.CalculateVerifyValues(ffsi.secretVals, ffsi.N)

	// print everything to console
	ffsi.printPrepartionValues(false)
}

func (ffsi *FeigeFiatShamirInitiator) printPrepartionValues(publicOnly bool) {
	publicSection, privateSection := "", ""

	publicSection = fmt.Sprintf(`
+-----------------------------------------------------------
| Public section
+-----------------------------------------------------------
p = %.40s			(Base64: %.40s	Binary: %.40s)
q = %.40s			(Base64: %.40s	Binary: %.40s)
N = p * q = %.40s		(Base64: %.40s	Binary: %.40s)
Verification values vᵢ
%s
	`,
		ffsi.p.String(), utils.ToBase64Str(*ffsi.p), ffsi.p.Text(2), // p
		ffsi.q.String(), utils.ToBase64Str(*ffsi.q), ffsi.q.Text(2), // q
		ffsi.N.String(), utils.ToBase64Str(*ffsi.N), ffsi.N.Text(2),
		ffsi.VerifyVals.String()) // N

	if !publicOnly {
		privateSection = fmt.Sprintf(`
+-----------------------------------------------------------
| Private section
+-----------------------------------------------------------
Secret values sᵢ
%s
		`, ffsi.secretVals.String())
	}
	var sb strings.Builder
	sb.WriteString(publicSection)
	sb.WriteString(privateSection)
	fmt.Println(sb.String())
}

// Creates two different prime numbers p and q and return them along with their product N
func (ffsi *FeigeFiatShamirInitiator) CreatePrimes() (p, q, N *big.Int) {
	log.Printf("Generating two secret prime numbers p and q...\n")
	// generate to secure prime numbers with given amount of bits
	p = utils.GenSecureRandPrime(ffsi.noSecretBits)
	q = utils.GenSecureRandPrime(ffsi.noSecretBits)

	// make sure p and q are different prime numbers (very unlikely in a big enough key space)
	for p.Cmp(q) == 0 {
		q = utils.GenSecureRandPrime(ffsi.noSecretBits)
	}
	log.Printf("Calculating product N of two secret prime numbers p and q...\n")
	N = big.NewInt(1).Mul(p, q)

	return p, q, N
}

func (ffsi *FeigeFiatShamirInitiator) GenerateSecretValues(noSecretVals int, N *big.Int) (secretVals BigIntArray) {
	log.Printf("Generating %d secret values...\n", noSecretVals)
	// allocate memory for verification values
	secretVals = make(BigIntArray, noSecretVals)
	// keep track of all random (secret) values -> do not use same (secret) random value twice (is very unlikely in big enough key space anyway)
	randInts := make(map[*big.Int]interface{})

	for i, secretCandidate := range secretVals {
		// only use unprecedented random values that are relatively prime to N
		for _, ok := randInts[&secretCandidate]; !utils.VerifyRelativelyPrime(secretCandidate, *N) && !ok; {
			secretCandidate = *utils.GenSecureRandom(ffsi.noSecretBits)
		}
		// a suitable candidate has been found!
		secretVals[i] = secretCandidate
	}
	return secretVals
}

func (ffsi *FeigeFiatShamirInitiator) CalculateVerifyValues(secretVals BigIntArray, n *big.Int) (verifyVals BigIntArray) {
	log.Printf("Calculating %d verification values from secret values...\n", len(secretVals))
	// create array for verification values
	verifyVals = make([]big.Int, len(secretVals))

	// calculate verification values: vᵢ = sᵢ² mod N for all 0 < i <= k
	for i, secret := range secretVals {
		squared := secret.Mul(&secret, &secret)
		verifyVals[i].Mod(squared, n)
	}
	return verifyVals
}

//send verification values, X and N to authentication responder (these are public and even when eavesdropped do not compromise security of protocol)
func (ffsi *FeigeFiatShamirInitiator) InitiateAuthentication() AuthenticationRequest {
	ffsi.s = GetSign()
	ffsi.r = utils.GenSecureRandom(ffsi.noSecretBits)
	ffsi.X = CalculateX(ffsi.s, ffsi.r, ffsi.N)

	ffsi.PrintAuthenticationStep1Values(false)

	return NewAuthenticationRequest(ffsi.VerifyVals, ffsi.X, ffsi.N)
}

func (ffsi *FeigeFiatShamirInitiator) PrintAuthenticationStep1Values(publicOnly bool) {
	publicSection, privateSection := "", ""
	publicSection = fmt.Sprintf(`
+-----------------------------------------------------------
| Authentication
+-----------------------------------------------------------
PUBLIC SECTION:
X = %40.40s			(Base64: %.40s	Binary: %.40s)
	`, ffsi.X.String(), utils.ToBase64Str(*ffsi.X), ffsi.X.Text(2)) // X

	if !publicOnly {
		privateSection = fmt.Sprintf(`
PRIVATE SECTION:
s = %d
r = %40.40s			(Base64: %.40s	Binary: %.40s)		
`,
			ffsi.s, //s
			ffsi.r, utils.ToBase64Str(*ffsi.r), ffsi.r.Text(2))
	}

	var sb strings.Builder
	sb.WriteString(publicSection)
	sb.WriteString(privateSection)
	log.Println(sb.String())
}

func CalculateX(s int8, r, N *big.Int) (X *big.Int) {
	log.Printf("Calculating X for authentication...\n")
	bigS := big.NewInt(int64(s))
	rSquared := big.NewInt(1).Mul(r, r)
	X = rSquared.Mod(rSquared, N).Mul(rSquared, bigS)
	//X = bigS.Mul(bigS, rSquared).Mod(bigS, N)
	return X
}

func GetSign() int8 {
	log.Printf("Get sign for authentication...\n")
	randomVal, err := rand.Int(rand.Reader, big.NewInt(1))
	if err != nil {
		log.Fatal("Could not generate random number \u2208 {0, 1}")
	}

	intVal := randomVal.Int64()
	if intVal == 1 {
		return 1
	}
	return -1
}

func (ffsi *FeigeFiatShamirInitiator) OnAuthenticationResponse(as []int8) AuthenticationVerifyRequest {

	product := big.NewInt(1)
	for i, a := range as {
		if a != 0 {
			product.Mul(product, &ffsi.secretVals[i])
		}
	}

	// y = r * (sᵢ)ᵃᵢ mod N
	y := ffsi.r.Mul(ffsi.r, product).Mod(ffsi.r, ffsi.N)
	log.Printf("\nPUBLIC SECTION:\ny = %.40s		(Base64 %.40s	Binary: %.40s)\n", y.String(), utils.ToBase64Str(*y), y.Text(2))
	return NewAuthenticationVerifyRequest(y)
}

func (ffsi *FeigeFiatShamirInitiator) OnAuthenticationVerifyResponse(isAccessGranted bool) {
	if isAccessGranted {
		log.Println("Access was granted!")
		log.Println("Doing some work now...")
	} else {
		log.Println("Access was denied!")
		log.Println("Backing off...")
	}
}
