// VDF
//
// Copyright 2019 by KeyFuse
//
// GPLv3 License

package sqrt

import (
	"math/big"
)

// VDFSqrt --
type VDFSqrt struct {
	p *big.Int
}

// NewVDFSqrt -- creates new VDF.
func NewVDFSqrt(p *big.Int) *VDFSqrt {
	return &VDFSqrt{
		p: p,
	}
}

// Euler's criterion is a formula for determining whether an integer is a quadratic residue modulo a prime.
// A^((p-1)/2) % p, if it equals 1 then A is a quadratic residue module a prime number
func (v *VDFSqrt) quadraticResidue(x *big.Int) bool {
	// p-1/2
	xx := new(big.Int).Set(x)
	t := new(big.Int).Exp(xx, new(big.Int).Div(new(big.Int).Sub(v.p, big.NewInt(1)), big.NewInt(2)), v.p)
	return t.Cmp(big.NewInt(1)) == 0
}

// Delay -- eval the delay.
func (v *VDFSqrt) Delay(t int64, x *big.Int) *big.Int {
	var r *big.Int
	xx := new(big.Int).Set(x)

	for i := int64(0); i < t; i++ {
		if !v.quadraticResidue(xx) {
			xx = xx.Neg(xx).Mod(xx, v.p)
		}
		r = xx.ModSqrt(xx, v.p)
	}
	return r
}

// Verify -- verify the result y.
func (v *VDFSqrt) Verify(t int64, x *big.Int, y *big.Int) bool {
	var r *big.Int
	yy := new(big.Int).Set(y)

	if !v.quadraticResidue(x) {
		x = big.NewInt(0).Mod(big.NewInt(0).Neg(x), v.p)
	}
	for i := int64(0); i < t; i++ {
		r = yy.Exp(yy, big.NewInt(2), v.p)
	}
	return r.Cmp(x) == 0
}
