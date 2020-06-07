package simplemath

import "testing"

func TestAdd(t *testing.T) {
	r := Add(1, 2)
	if r != 3 {
		t.Errorf("Add(1,2) failed, Got %d, exspected 3.", r)
	}
}

func TestAdd2(t *testing.T) {
	r := Add(1, 2)
	if r != 3 {
		t.Errorf("Add(1,2) failed, Got %d, exspected 3.", r)
	}
}
