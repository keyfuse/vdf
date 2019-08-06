// VDF
//
// Copyright 2019 by KeyFuse
//
// GPLv3 License

package sqrt

import (
	"math/big"
	"testing"
)

func TestSqrt(t *testing.T) {
	tv := int64(1000)
	xv := big.NewInt(7)
	p256 := "60464814417085833675395020742168312237934553084050601624605007846337253615407"
	pv, _ := new(big.Int).SetString(p256, 0)

	sqrt := NewVDFSqrt(pv)
	yv := sqrt.Delay(tv, xv)

	yv1, _ := new(big.Int).SetString("37602270514657052337634791495305873950174757777621818362566683935496339681801", 0)
	if yv.Cmp(yv1) != 0 {
		t.Fatalf("yv:%v != yv1:%v", yv, yv1)
	}

	vv := sqrt.Verify(tv, xv, yv)
	if vv != true {
		t.Fatalf("verify.failed:%v", yv)
	}
}
