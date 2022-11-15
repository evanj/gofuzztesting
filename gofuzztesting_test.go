package gofuzztesting

import (
	"encoding/binary"
	"fmt"
	"testing"
)

const panicIfArgAboveLimit = 5000

func buggyFunc(arg int) {
	if arg > panicIfArgAboveLimit {
		panic(fmt.Sprintf("BUG arg=%d > %d", arg, panicIfArgAboveLimit))
	}
}

// This test is not able to find failing test cases quickly on my machine.
func FuzzInt(f *testing.F) {
	fuzzFunc := func(t *testing.T, arg int) {
		buggyFunc(arg)
	}
	f.Fuzz(fuzzFunc)
}

// This fuzz tests finds a failing test quickly.
func FuzzWithCorpus(f *testing.F) {
	// f.Add(panicIfArgAboveLimit - 450) seems to reliably find the bug within 1 min
	// f.Add(panicIfArgAboveLimit - 500) does not
	f.Add(panicIfArgAboveLimit - 450)

	fuzzFunc := func(t *testing.T, arg int) {
		buggyFunc(arg)
	}
	f.Fuzz(fuzzFunc)
}

// This fuzz test finds failing cases very quickly.
func FuzzBytes(f *testing.F) {
	fuzzFunc := func(t *testing.T, arg []byte) {
		var intBytes [8]byte
		copy(intBytes[:], arg)
		intArg := int(binary.LittleEndian.Uint64(intBytes[:]))
		buggyFunc(intArg)
	}
	f.Fuzz(fuzzFunc)
}
