package beauty

import "testing"

func TestCQ(t *testing.T) {
	cq := New(5)
	for i := 0; i < 20; i++ {
		enqueue := cq.Enqueue(i)
		t.Log(enqueue)
		if i&1 > 0 {
			dequeue, b := cq.Dequeue()
			t.Log(dequeue, b)
		}
	}
}
