// Created on 18 June 2016
// author youenn.kahlys.laborde

package factorization

import (
	"math/big"

	"github.com/kahlys/crypto/math"
)

// BruteForce returns a primal number in n decomposition.
func BruteForce(n *big.Int) *big.Int {
	s := math.SqrtCeil(n)
	i := big.NewInt(2)
	for i.Cmp(s) == -1 {
		m := new(big.Int)
		m.Mod(n, i)
		if m.Cmp(big.NewInt(0)) == 0 {
			return i
		}
		i.Add(i, big.NewInt(1))
	}
	return big.NewInt(1)
}

// BruteForce returns a primal number in n decomposition. It uses the pseudo-random
// sequence x^2 +1 mod n.
func RhoPollard(n *big.Int) *big.Int {
	t, x := big.NewInt(1), big.NewInt(1)
	y := randSequence(x, n)

	yx := new(big.Int).Sub(y, x)
	d := new(big.Int).GCD(nil, nil, yx, n)

	if d.Cmp(big.NewInt(1)) != 0 {
		return d
	}

	x.Set(y)
	y = randSequence(x, n)
	t.Mul(t, big.NewInt(2))

	for {
		i := big.NewInt(1)
		for i.Cmp(t) < 1 {
			yx.Sub(y, x)
			d.GCD(nil, nil, yx, n)

			if d.Cmp(big.NewInt(1)) != 0 {
				return d
			}
			if d.Cmp(big.NewInt(1)) == 0 && i.Cmp(t) == -1 {
				y = randSequence(y, n)
			}
			i.Add(i, big.NewInt(1))
		}
		x.Set(y)
		y = randSequence(x, n)
		t.Mul(t, big.NewInt(2))
	}
}

// y = x^2 +1 mod n
func randSequence(x, n *big.Int) *big.Int {
	r := new(big.Int)
	r.Exp(x, big.NewInt(2), n)
	r.Add(r, big.NewInt(1))
	r.Mod(r, n)
	return r
}
