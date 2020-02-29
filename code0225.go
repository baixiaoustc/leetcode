/*
使用队列实现栈的下列操作：

push(x) -- 元素 x 入栈
pop() -- 移除栈顶元素
top() -- 获取栈顶元素
empty() -- 返回栈是否为空
注意:

你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合法的。
你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-stack-using-queues
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

type MyStack struct {
	a *LinkedQueue
	b *LinkedQueue
}

/** Initialize your data structure here. */
func ConstructorMyStack() MyStack {
	return MyStack{
		a: NewLinkedQueue(),
		b: NewLinkedQueue(),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	if this.a.Size() != 0 {
		this.a.Add(x)
		return
	}
	if this.b.Size() != 0 {
		this.b.Add(x)
		return
	}
	//默认选a
	this.a.Add(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	for this.a.Size() > 0 {
		x, _ := this.a.Peek()
		this.a.Remove()
		if this.a.Size() == 0 {
			return x.(int)
		} else {
			this.b.Add(x)
		}
	}

	for this.b.Size() > 0 {
		x, _ := this.b.Peek()
		this.b.Remove()
		if this.b.Size() == 0 {
			return x.(int)
		} else {
			this.a.Add(x)
		}
	}

	return -1
}

/** Get the top element. */
func (this *MyStack) Top() int {
	for this.a.Size() > 0 {
		x, _ := this.a.Peek()
		this.a.Remove()
		this.b.Add(x)
		if this.a.Size() == 0 {
			return x.(int)
		}
	}

	for this.b.Size() > 0 {
		x, _ := this.b.Peek()
		this.b.Remove()
		this.a.Add(x)
		if this.b.Size() == 0 {
			return x.(int)
		}
	}

	return -1
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.a.Size()+this.b.Size() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
