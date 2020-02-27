package main

import (
	"strings"
	"testing"
)

func TestSorted(t *testing.T) {
	arr := []string{"b", "ą", "ć", "c", "a", "ż", "ś"}

	expected := []string{"a", "ą", "b", "c", "ć", "ś", "ż"}
	sorted(arr, "POLISH_CI")

	if strings.Join(arr, ",") != strings.Join(expected, ",") {
		t.Errorf("Expected %v, but got %v", expected, arr)
	}

}
