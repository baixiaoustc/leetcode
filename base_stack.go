package leetcode

type StackInterface struct {
	data []interface{}
}

func (stack *StackInterface) Push(value interface{}) {
	stack.data = append(stack.data, value)
}

func (stack *StackInterface) Pop() interface{} {
	if len(stack.data) == 0 {
		panic("stack empty")
	}

	value := stack.data[len(stack.data)-1]
	stack.data = stack.data[0 : len(stack.data)-1]
	return value
}

func (stack *StackInterface) Peak() interface{} {
	if len(stack.data) == 0 {
		panic("stack empty")
	}

	value := stack.data[len(stack.data)-1]
	return value
}

func (stack *StackInterface) Size() int {
	return len(stack.data)
}

func NewStackInterface() *StackInterface {
	return &StackInterface{
		data: make([]interface{}, 0),
	}
}
