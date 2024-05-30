package test

import (
	"github.com/emirpasic/gods/trees/redblacktree"
	"testing"
)

func TestRedblacktreee(t *testing.T) {
	rbt := redblacktree.NewWithIntComparator()
	rbt.Put(0, 0)
	rbt.Put(100, 100)
	rbt.Put(50, 50)
	//it := rbt.Iterator()
	//it.Begin()
	floor, ff := rbt.Floor(50)
	t.Log(floor.Key, floor.Value, ff)
	ceil, cf := rbt.Ceiling(100)
	t.Log(ceil.Key, ceil.Value, cf)
}
