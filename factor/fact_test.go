// Created on 18 June 2016
// author youenn.kahlys.laborde

package factorization

import (
	"math/big"
	"testing"
)

func TestBruteForce(t *testing.T) {
	n := big.NewInt(123456789)
	x := BruteForce(n)
	m := new(big.Int).Mod(n, x)
	if m.Cmp(big.NewInt(0)) != 0 {
		t.Error("brute-force failed")
	}
}

func TestRhoPollard(t *testing.T) {
	n := big.NewInt(187)
	x := RhoPollard(n)
	m := new(big.Int).Mod(n, x)
	if m.Cmp(big.NewInt(0)) != 0 {
		t.Error("rho-pollard failed")
	}

}
