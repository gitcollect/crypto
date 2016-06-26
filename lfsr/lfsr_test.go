// created on 26 june 2016
// author: youenn.kahlys.laborde

package lfsr

import (
	"math/big"
	"testing"
)

func TestLFSR(t *testing.T) {
	p, _ := new(big.Int).SetString("2d", 16)   // p = x16 + x14 + x13 + x11 + 1
	r, _ := new(big.Int).SetString("ace1", 16) // nonzero start
	l := NewLFSR(r,p,16)

	l.Inc()
	step := l.Register.Text(16)
	if step != "5670" {
		t.Error("LFSR first cycle is", step, "wants 5670")
	}
}
