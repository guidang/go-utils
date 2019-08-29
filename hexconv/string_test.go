package hexconv

import (
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	in := "粤B88888"
	out := "e7b2a4423838383838"
	if str := String(in); str != out {
		t.Errorf("String(%q) = %v, want %v, %v",
			in, str, out, "incorrect value")
	}
}

func TestBytes(t *testing.T) {
	in := "31303536392e6d6373"
	out := "10569.mcs"
	val, err := Bytes(in)

	if err != nil {
		t.Errorf("Bytes(%q), err = %v",
			in, err)
	} else if !strings.EqualFold(string(val), out) {
		t.Errorf("Bytes(%q) = %v, want %v, %v",
			in, string(val), out, "incorrect value")
	}
}
