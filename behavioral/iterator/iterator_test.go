package iterator

import "testing"

func TestIterator(t *testing.T) {
	la := NewListAggregate(1, 9)

	printIterator(la)
}
