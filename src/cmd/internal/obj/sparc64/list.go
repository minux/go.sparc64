// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sparc64

import (
	"cmd/internal/obj"
	"fmt"
)

func init() {
	obj.RegisterRegister(obj.RBaseSPARC64, REG_LAST, Rconv)
	obj.RegisterOpcode(obj.ABaseSPARC64, Anames)
}

func Rconv(r int) string {
	switch {
	case r == RegFP:
		return "RFP"
	case r == RegRSP:
		return "RSP"
	case r == RegZero:
		return "ZR"
	case r == REG_BSP:
		return "BSP"
	case r == REG_BFP:
		return "BFP"
	}
	switch {
	case REG_R0 <= r && r <= REG_R31:
		return fmt.Sprintf("R%d", r-REG_R0)
	case REG_F0 <= r && r <= REG_F31:
		return fmt.Sprintf("F%d", r-REG_F0)
	}
	return fmt.Sprintf("badreg(%d)", r)
}

func DRconv(a int8) string {
	if a >= ClassUnknown && a <= ClassNone {
		return cnames[a]
	}
	return "C_??"
}
