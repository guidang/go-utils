package hexconv

import (
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
