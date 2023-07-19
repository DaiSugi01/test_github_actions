package test

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	total := Add()
	for i := 0; i < 100000; i++ {
	}
	if total != 2 {
		fmt.Println("a")
		t.Errorf("Add(1, 2) = %d; want 3", total)
	}
}
