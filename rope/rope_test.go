package rope

import (
	"testing"
)

func TestBasics(t *testing.T) {
	r := Rope[int]{}

	r.Insert(1, 0)
	r.Insert(2, 1)
	r.Insert(0, 0)
	r.Insert(3, 3)

	if got := r.Len(); got != 4 {
		t.Errorf("expected r.Len() to return 4 but got %d", got)
	}

	for i :=0; i < 4; i++ {
		if got := r.Get(i); got != i {
			t.Errorf("expected r.Get(%d) to return %d but got %d", i, i, got)
		}
	}
}

func TestConcat(t *testing.T) {
	r1 := Rope[int]{}
	r1.Insert(0, 0)
	r1.Insert(1, 1)
	r1.Insert(2, 2)
	r1.Insert(3, 4)

	r2 := Rope[int]{}
	r2.Insert(4, 0)
	r2.Insert(5, 1)
	r2.Insert(6, 2)
	r2.Insert(7, 3)

	r1.Concat(&r2);
	
	if got := r1.Len(); got != 8 {
		t.Errorf("expected r.Len() to return 8 but got %d", got)
	}

	for i := 0; i < 8; i++ {
		if got := r1.Get(i); got != i {
			t.Errorf("expected r.Get(%d) to return %d but got %d", i, i, got)
		}
	}
}