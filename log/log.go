// Created on 17 June 2016
// author: youenn.kahlys.laborde

package log

import (
	"math/big"

	"github.com/kahlys/crypto/math"
)

// BruteForce finds x where g^x = h mod p. If there is no such x, the function
// returns nil.
func BruteForce(g, h, p *big.Int) *big.Int {
	t := big.NewInt(1)
	for i := big.NewInt(1); i.Cmp(p) == -1; i.Add(i, big.NewInt(1)) {
		t.Mul(t, g)
		t.Mod(t, p)
		if t.Cmp(h) == 0 {
			return i
		}
	}
	return nil
}

// Shanks finds x where g = h mod p, with the Baby Steps Giant Steps attack. If
// there is no such x, the function returns nil.
func Shanks(g, h, p *big.Int) *big.Int {
	w := new(big.Int).Sub(p, big.NewInt(1))
	s := math.SqrtCeil(p)

	// gitoi[g^i] = i
	gitoi := make(map[string]*big.Int)
	gitoi["1"] = big.NewInt(0)

	// baby steps
	t := big.NewInt(1)
	for i := big.NewInt(1); i.Cmp(s) == -1; i.Add(i, big.NewInt(1)) {
		t.Mul(t, g)
		t.Mod(t, p)
		gitoi[t.String()] = new(big.Int).Set(i)
	}

	// gs = g^(-s) = g^(w-s)
	ws := new(big.Int).Sub(w, s)
	gs := new(big.Int).Exp(g, ws, p)

	// giant steps h.g^(-ys)
	for i := big.NewInt(0); i.Cmp(s) == -1; i.Add(i, big.NewInt(1)) {
		hgs := new(big.Int).Exp(gs, i, p)
		hgs.Mul(hgs, h)
		hgs.Mod(hgs, p)

		// g^i == h.g^(-ys) => x = sy + i mod w
		if v, ok := gitoi[hgs.String()]; ok {
			r := new(big.Int).Mul(s, i)
			r.Add(r, v)
			r.Mod(r, w)
			return r
		}
	}

	return nil
}
