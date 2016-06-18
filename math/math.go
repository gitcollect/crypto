// Created on 18 June 2016
// author youenn.kahlys.laborde

package math

import (
    "math/big"
)

// SqrtFloor is the copy of a removed function from older version of GO.
func SqrtFloor(n *big.Int) (x *big.Int) {
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

func SqrtCeil(n *big.Int) *big.Int {
    x := SqrtFloor(n)
    x.Add(x, big.NewInt(1))
    return x
}
