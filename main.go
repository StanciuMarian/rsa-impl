package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var smallPrimes = []uint8{
	3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53,
}

type KeySize int

const (
	Key1024 = 1024
	Key2048 = 2048
	Key3072 = 3072
)

func main() {

Start:
	candidate := getPrimeCandidate(Key3072)
	fmt.Println("new candidate len == %d", len(candidate.String()))
	if candidate.Bit(0) == 0 {
		fmt.Println("discarded one")
		goto Start
	}

	for _, smallPrime := range smallPrimes {
		mod := big.NewInt(0)
		if len(mod.Mod(candidate, big.NewInt(int64(smallPrime))).Bits()) == 0 {
			fmt.Println("discarded one")
			goto Start
		}
	}

	if !passesRabinMiller(candidate, 50) {
		goto Start
	}

	fmt.Printf("Random prime number generated: len(prime) = %d, value = %s", len(candidate.String()), candidate.String())
}

func getPrimeCandidate(keyS KeySize) *big.Int {
	twoBase := big.NewInt(2)
	keySize := big.NewInt(int64(keyS))
	almostKeySize := big.NewInt(int64((keyS) - 1))

	min := big.NewInt(0)
	min.Exp(twoBase, almostKeySize, nil)
	min.Add(min, big.NewInt(1))

	max := big.NewInt(0)
	max.Exp(twoBase, keySize, nil)
	max.Sub(max, big.NewInt(-1))

	random := pickRandomInRange(min, max)
	return random.Add(random, max)
}

func pickRandomInRange(to, from *big.Int) *big.Int {
	maxRand := big.NewInt(0)
	maxRand.Sub(from, to)
	random, err := rand.Int(rand.Reader, maxRand)
	if err != nil {
		panic(err)
	}

	return random
}

var two = big.NewInt(2)

func passesRabinMiller(candidate *big.Int, noRounds int) bool {
	d, r := getFactors(candidate)
	candidateSubOne := big.NewInt(1).Sub(candidate, big.NewInt(1))

	for i := 0; i < noRounds; i++ {
		fmt.Printf("round %d\n", i)
		a := pickRandomInRange(two, big.NewInt(1).Sub(candidate, two))
		x := big.NewInt(1)
		x.Exp(a, d, candidate)
		if (len(x.Bits()) == 1 && x.Bit(0) == 1) || x.Cmp(candidateSubOne) == 0 {
			continue
		}
		fmt.Println("`going into the second part of the round")
		for j := 1; uint(j) < r; j++ {

			x = x.Mul(x, two).Mod(x, candidate)
			if x.Cmp(candidateSubOne) == 0 {
				continue
			}
		}
		return false
	}
	return true
}

func getFactors(newInt *big.Int) (*big.Int, uint) {
	if newInt.Bit(0) == 0 {
		panic("even number")
	}
	subOne := big.NewInt(0)
	subOne.SetBit(newInt, 0, 0)
	r := subOne.TrailingZeroBits()
	d := big.NewInt(1)
	return d.Rsh(subOne, r), r
}
