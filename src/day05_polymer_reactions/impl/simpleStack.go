package impl

type StringStack struct {
	values []string
	length int
}

func myPush(stack *StringStack, val string) {
	stack.values = append(stack.values, val)
	stack.length++
}

func myPop(stack *StringStack) string {
	lastVal := ""

	if stack.length > 0 {
		lastVal = stack.values[len(stack.values)-1]
		stack.values = stack.values[:len(stack.values)-1]
		stack.length--
	}

	return lastVal
}

func myPeek(stack *StringStack) string {
	lastVal := ""

	if stack.length > 0 {
		lastVal = stack.values[len(stack.values)-1]
	}

	return lastVal
}
