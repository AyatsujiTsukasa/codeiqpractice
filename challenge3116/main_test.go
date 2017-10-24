package main

import (
	"testing"
)

func TestConvert(t *testing.T) {
	x := convert("95422357545868773174")
	if x != "9257583174" {
		t.Errorf("got %v\n want %v", x, "9257583174")
	}
}
