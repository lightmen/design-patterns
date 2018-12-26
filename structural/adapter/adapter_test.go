package adapter

import (
	"testing"
)

var expect = "SpecificRequest"

func TestAdapter(t *testing.T) {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)

	result := target.Request()
	if result != expect {
		t.Fatalf("expect: %s, actual: %s", expect, result)
	}else {
		t.Logf("%s\n", result)
	}
}
