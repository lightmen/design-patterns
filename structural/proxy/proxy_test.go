package proxy

import (
	"testing"
	"fmt"
)

func TestProxy(t *testing.T) {
	real := &RealSubject{}

	var sub Subject
	sub = &Proxy{
		real: real,
	}

	req := sub.Request()
	fmt.Println(req)
}
