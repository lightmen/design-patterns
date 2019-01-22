package composite

import "testing"

func TestComposite(t *testing.T) {
	l1 := NewLeaf("leaf1")
	l2 := NewLeaf("leaf2")

	c := NewComposite("composite")
	c.Add(l1)
	c.Add(l2)

	c.Print("")

	c.Remove(l2)
	c.Print("")
}
