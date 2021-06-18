package ineed

import (
	"testing"
)

func Test_StringToINT64(t *testing.T) {
	s := "9223372036854775807"
	var bigint int64

	if !StringtoINT64(s, &bigint) {
		t.Fail()
	}

	t.Logf("\nstring %v ", bigint)
}

func Test_StringToBYTE(t *testing.T) {
	s := "922337203"
	s = "9"

	newByte, ok := StringtoBYTE(s)
	t.Logf("\nstring %v %v ", newByte, ok)
}
