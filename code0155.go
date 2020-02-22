/*
设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) -- 将元素 x 推入栈中。
pop() -- 删除栈顶的元素。
top() -- 获取栈顶元素。
getMin() -- 检索栈中的最小元素。
示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/min-stack
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

type MinStack struct {
	data   []int
	helper []int
}

/** initialize your data structure here. */
func ConstructorMinStack() MinStack {
	return MinStack{
		data:   make([]int, 0),
		helper: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.helper) == 0 {
		this.helper = append(this.helper, x)
	} else {
		if x <= this.helper[len(this.helper)-1] {
			this.helper = append(this.helper, x)
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.data) == 0 {
		panic("no data")
	}

	x := this.data[len(this.data)-1]
	this.data = this.data[0 : len(this.data)-1]
	if x == this.helper[len(this.helper)-1] {
		this.helper = this.helper[0 : len(this.helper)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.data) == 0 {
		panic("no data")
	}

	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.helper) == 0 {
		panic("no data")
	}

	return this.helper[len(this.helper)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
