package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
	iter     *Iterator
	_hasNext bool
	_next    int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter: iter, _hasNext: iter.hasNext(), _next: iter.next()}
}

func (this *PeekingIterator) hasNext() bool {
	return this._hasNext
}

func (this *PeekingIterator) next() int {
	ans := this._next
	this._hasNext = this.iter.hasNext()
	if this._hasNext {
		this._next = this.iter.next()
	}
	return ans
}

func (this *PeekingIterator) peek() int {
	return this._next
}

// leetcode submit region end(Prohibit modification and deletion)

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
}

func (this *Iterator) next() int {
}
