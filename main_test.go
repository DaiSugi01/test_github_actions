package test

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	total := Add()
	if total != 2 {
		fmt.Println("a")
		t.Errorf("Add(1, 2) = %d; want 3", total)
	}
	if total != 2 {
		fmt.Println("a")
		t.Errorf("Add(1, 2) = %d; want 3", total)
	}
}
