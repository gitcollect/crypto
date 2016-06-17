// Created on 17 June 2016
// author: youenn.kahlys.laborde

package log

import (
	"math/big"
	"testing"
)

func TestBruteForce(t *testing.T) {
	g := big.NewInt(2)
	h := big.NewInt(5)
	p := big.NewInt(1019)

	x := BruteForce(g, h, p)
	if x == nil {
		t.Fatal("BruteForce failed")
	}

	r := g.Exp(g, x, p)
	if r.Cmp(h) != 0 {
		t.Error("BruteForce finds", r, "wants", h)
	}
}

func TestShanks(t *testing.T) {
	g := big.NewInt(2)
	h := big.NewInt(5)
	p := big.NewInt(1019)

	x := Shanks(g, h, p)
	if x == nil {
		t.Fatal("Shanks failed")
	}

	r := g.Exp(g, x, p)
	if r.Cmp(h) != 0 {
		t.Error("BruteForce finds", r, "wants", h)
	}
}
