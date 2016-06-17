// Created on 17 June 2016
// author: youenn.kahlys.laborde

package log

import (
	"math/big"
)

// BruteForce finds x where g^x = h mod p. If there is no such x, the function
// returns nil.
func BruteForce(g, h, p *big.Int) *big.Int {
	t := big.NewInt(1)
	i := big.NewInt(1)
	for i.Cmp(p) == -1 {
		t.Mul(t, g)
		t.Mod(t, p)
		if t.Cmp(h) == 0 {
			return i
		}
		i.Add(i, big.NewInt(1))
	}
	return nil
}

// Shanks finds x where g = h mod p, with the Baby Steps Giant Steps attack. If
// there is no such x, the function returns nil.
func Shanks(g, h, p *big.Int) *big.Int {
	w := new(big.Int)
	w.Sub(p, big.NewInt(1))
	s := ceilSqrt(p)

	// gitoi[g^i] = i
	gitoi := make(map[string]*big.Int)
	gitoi["1"] = big.NewInt(0)

	// baby steps
	t := big.NewInt(1)
	i := big.NewInt(1)
	for i.Cmp(s) == -1 {
		t.Mul(t, g)
		t.Mod(t, p)

		gitoi[t.String()] = new(big.Int)
		gitoi[t.String()].Set(i)

		i.Add(i, big.NewInt(1))
	}

	// gs = g^(-s) = g^(w-s)
	ws := new(big.Int)
	ws.Sub(w, s)
	gs := new(big.Int)
	gs.Exp(g, ws, p)

	// giant steps h.g^(-ys)
	i = big.NewInt(0)
	for i.Cmp(s) == -1 {
		hgs := new(big.Int)
		hgs.Exp(gs, i, p)
		hgs.Mul(hgs, h)
		hgs.Mod(hgs, p)

		// g^i == h.g^(-ys) => x = sy + i mod w
		if v, ok := gitoi[hgs.String()]; ok {
			r := new(big.Int)
			r.Mul(s, i)
			r.Add(r, v)
			r.Mod(r, w)
			return r
		}
	}

	return nil
}

func ceilSqrt(n *big.Int) *big.Int {
	x := floorSqrt(n)
	x.Add(x, big.NewInt(1))
	return x
}

// floorSqrt is a removed function from older version of GO.
func floorSqrt(n *big.Int) (x *big.Int) {
	switch n.Sign() {
	case -1:
		panic(-1)
	case 0:
		return big.NewInt(0)
	}

	var px, nx big.Int
	x = big.NewInt(0)
	x.SetBit(x, n.BitLen()/2+1, 1)
	for {
		nx.Rsh(nx.Add(x, nx.Div(n, x)), 1)
		if nx.Cmp(x) == 0 || nx.Cmp(&px) == 0 {
			break
		}
		px.Set(x)
		x.Set(&nx)
	}
	return
}
