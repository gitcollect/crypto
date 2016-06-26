// created on 26 june 2016
// author: youenn.kahlys.laborde

package lfsr

import (
	"math/big"
)

type LFSR struct {
	Register *big.Int
	Taps     *big.Int
	Size     uint
}

// NewLFSR creates a Linead-feedback shift register, when the register and the
// feedback polynome are represented by bits of numbers.
func (l *LFSR) NewLFSR(reg, taps *big.Int) {
	l.Register = reg
	l.Taps = taps
}

// Inc cycles the register and returns the output bit.
func (l *LFSR) Inc() uint {
	tmp := new(big.Int).Xor(l.Register, l.Taps)
	bit := parity(tmp)
	l.Register.Rsh(l.Register, 1)
	retro := new(big.Int).Lsh(big.NewInt(int64(bit)), l.Size-1)
	l.Register.Xor(l.Register, retro)
	return bit
}

func parity(n *big.Int) uint {
	t := new(big.Int).Set(n)
	t.Xor(t, new(big.Int).Rsh(t, 32))
	t.Xor(t, new(big.Int).Rsh(t, 16))
	t.Xor(t, new(big.Int).Rsh(t, 8))
	t.Xor(t, new(big.Int).Rsh(t, 4))
	t.Xor(t, new(big.Int).Rsh(t, 2))
	t.Xor(t, new(big.Int).Rsh(t, 1))
	return t.Bit(0)
}

func (l *LFSR) String() string {
	var rr, rp string
	rr = l.Register.Text(2)
	for i := uint(len(rr)); i < l.Size; i++ {
		rr = "0" + rr
	}
	rp += l.Taps.Text(2)
	for i := uint(len(rp)); i < l.Size; i++ {
		rp = "0" + rp
	}
	return rr + "\n" + rp
}
