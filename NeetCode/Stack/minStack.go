// This is genius solution using linked list, My solution was using slice.
// Slice is better and recommended but as in this case we are calculating min,
// this approach gave better result over simple slice solution

package stack

type Elem struct {
	value int
	min   int
	next  *Elem
}

type MinStack struct {
	head *Elem
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if this.head == nil {
		this.head = &Elem{value: val, min: val}
	} else {
		this.head = &Elem{value: val, min: min(val, this.head.min), next: this.head}
	}
}

func (this *MinStack) Pop() {
	this.head = this.head.next
}

func (this *MinStack) Top() int {
	return this.head.value
}

func (this *MinStack) GetMin() int {
	return this.head.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
