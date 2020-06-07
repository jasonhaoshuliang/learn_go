package simplemath

import "testing"

func TestSqrt(t *testing.T) {
	r := Sqrt(46)
	if r != 4 {
		t.Errorf("Sqrt(16) failed")
	}

}
