package cappedCache

import "testing"

func TestCappedCache(t *testing.T) {
	c := New(4)
	c.Insert(0)
	c.Insert(1)
	c.Insert(2)
	c.Insert(3)
	c.Insert(4)

	for i, val := range c.List() {
		if val.(int) != i+1 {
			t.Fatalf("unexpected value, expected %v and received %v", i+1, val)
		}
	}
}
