package _4_实现一个最小栈

type MinStack struct {
	stack1 []int
	stack2 []int
}

func Constructor() MinStack {
	return MinStack{
		stack1: []int{},
		stack2: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack1 = append(this.stack1, val)
	if len(this.stack2) > 0 {
		if this.stack2[len(this.stack2)-1] > val {
			this.stack2 = append(this.stack2, val)
		} else {
			this.stack2 = append(this.stack2, this.stack2[len(this.stack2)-1])
		}
	} else if len(this.stack2) == 0 {
		this.stack2 = append(this.stack2, val)
	}
}

func (this *MinStack) Pop() {
	this.stack1 = this.stack1[:len(this.stack1)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
}

func (this *MinStack) Top() int {
	return this.stack1[len(this.stack1)-1]
}

func (this *MinStack) GetMin() int {
	return this.stack2[len(this.stack2)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
