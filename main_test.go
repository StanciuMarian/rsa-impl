package main

import (
	"math/big"
	"testing"
)

type test struct {
	input     *big.Int
	expectedD *big.Int
	expectedR uint
}

func Test_Decomposition(t *testing.T) {
	testValues := []test{
		// 2 ** r * d + 1
		{big.NewInt(9), big.NewInt(1), 3},  // 2 ** 3 * 1 + 1
		{big.NewInt(11), big.NewInt(5), 1}, // 2 ** 1 * 5 + 1
		{big.NewInt(21), big.NewInt(5), 2}, // 2 ** 2 * 5 + 1
	}
	big.NewInt(8)

	for _, testValue := range testValues {
		d, r := getFactors(testValue.input)
		if d.Cmp(testValue.expectedD) != 0 {
			t.Errorf("got %s for d value for %s, expected %s", d.String(), testValue.input.String(), testValue.expectedD.String())
		}
		if r != testValue.expectedR {
			t.Errorf("got %d for r value for %s, expected %d", r, testValue.input, testValue.expectedR)
		}
	}

}
