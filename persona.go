package persona

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

type Persona struct {
	Name   string `bson:"name" json:"name"`
	Sex    string `bson:"sex" json:"sex"`
	Zodiac string `bson:"zodiac" json:"zodiac"`
}

func randomAddress() string {
	bytes := make([]byte, 20)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("0x%x", bytes)
}

func getDna(address string) [32]big.Int {
	var dna = [32]big.Int{}
	// do this via big.int for compatibility with ethers
	d := big.NewInt(0)
	b, _ := d.SetString(strings.Replace(address, "0x", "", 1), 16)
	keccak := crypto.Keccak256Hash(b.Bytes())
	str := strings.Replace(keccak.String(), "0x", "", 1)
	var count = 0

	for i := 0; i < len(str)-1; i = i + 2 {
		sub := string(str[i]) + string(str[i+1])
		n := big.NewInt(0)
		s, _ := n.SetString(sub, 16)
		dna[count].Set(s)
		count++
	}

	return dna
}

func dnaToPersona(dna [32]big.Int) Persona {
	var sum = big.NewInt(0)
	var even = big.NewInt(0)
	var odd = big.NewInt(0)
	var sex = "female"

	for i := 0; i < len(dna); i++ {
		sum.Add(sum, &dna[i])
		if i%2 == 0 {
			odd.Add(odd, &dna[i])
		} else {
			even.Add(even, &dna[i])
		}
	}

	mod := big.NewInt(0)
	mod.Mod(sum, big.NewInt(2))
	if mod.Cmp(big.NewInt(1)) == 0 {
		sex = "male"
	}

	mod.Mod(sum, big.NewInt(12))
	zodiac := zodiacs[mod.Int64()]

	var given string
	if sex == "female" {
		index := math.Floor(float64(even.Int64()) / 8)
		given = Names.Female[int64(index)]
	} else {
		index := math.Floor(float64(odd.Int64()) / 8)
		given = Names.Male[int64(index)]
	}

	index := math.Floor(float64(sum.Int64()) / 2)
	var family = Names.Family[int64(index)]

	return Persona{
		Name:   given + " " + family,
		Zodiac: zodiac,
		Sex:    sex,
	}
}

func New(address string) Persona {
	if address == "" {
		address = randomAddress()
	}
	dna := getDna(address)
	p := dnaToPersona(dna)
	return p
}
