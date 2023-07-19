package test

import "testing"

func TestAdd(t *testing.T) {
	total := Add()
	if total != 2 {
		t.Errorf("Add(1, 2) = %d; want 3", total)
	}
}
