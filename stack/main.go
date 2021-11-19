package stack

type Stack struct {
	arr       []interface{}
	lastIndex int
}

func NewStack() Stack {
	return Stack{
		arr:       make([]interface{}, 0),
		lastIndex: -1,
	}
}

func (stack *Stack) Peek() interface{} {
	if stack.IsEmpty() {
		panic("Stack is empty!")
	}
	return stack.arr[stack.lastIndex]
}

func (stack *Stack) Push(a interface{}) {
	stack.arr = append(stack.arr, a)
	stack.lastIndex++
}

func (stack *Stack) Pop() interface{} {
	if stack.IsEmpty() {
		panic("Stack is empty!")
	}
	temp := stack.arr[stack.lastIndex]
	stack.lastIndex--
	stack.arr = stack.arr[:len(stack.arr)-1]
	return temp
}

func (stack *Stack) IsEmpty() bool {
	return stack.lastIndex == -1
}
