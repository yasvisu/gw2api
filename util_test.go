package gw2api

import "testing"

// For now we assume we only need at most uint32
// On 64-Bit systems uint does default to uint64 though
func TestFlagSet(t *testing.T) {
	var n, i uint
	for i = 0; i < 32; i++ {
		flagSet(&n, i)
	}
	if uint32(n) != ^uint32(0) {
		t.Error("UInt result is not correct")
	}
}

func TestFlagGet(t *testing.T) {
	var n, i uint
	for i = 0; i < 32; i++ {
		flagSet(&n, i)
	}
	for i = 0; i < 32; i++ {
		if !flagGet(n, i) {
			t.Error("Flag should have been set for ", i)
		}
	}
}
