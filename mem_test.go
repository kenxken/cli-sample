package main

import (
	"testing"
)

func TestMemStatus(t *testing.T) {
	result := MemStatus()
	if result == 0 {
		t.Fatal("failed test")
	}
}
