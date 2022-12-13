package main

import (
	"fmt"
	"strconv"
)

func main() {
	testStackStructure()
}

/*

	Implement the **stack** data structure. It has cells to contain data. For example, integers `1`, `2`, `3`, `4`, and so on. The cells are *indexed* from the bottom (index 0) to the top (index n). Let’s assume **n=3** for this exercise, so we have **4** places. A new stack contains **0** in all cells. A new value is put in the highest cell, which is empty (containing 0) on top (of the stack): this is called **push**. To get a value from the stack, take the value in the highest cell which is not 0: this is called **pop**. We can understand why a stack is called a *Last In First Out (LIFO)* structure.
 	Define a new type `Stack` for this data structure. Make 2 methods `Push` and `Pop`.
	Make a `String()` method (for *debugging* purposes) that shows the content of the stack as: `[0:i] [1:j] [2:k] [3:l]`. Take the underlying data structure, a **struct** containing an `index`, an array `data` of *int*, and the `ix` contains the first free position.

	Generalize the implementation by making the number of elements 4 a constant `LIMIT`.

*/

func testStackStructure() {
	var stack Stack
	stack.Push(18)
	stack.Push(26)
	stack.Push(35)
	stack.Push(47)
	stack.Pop()
	stack.Push(59)
	stack.Pop()
	stack.Push(66)
	fmt.Printf("%v\n", stack)
}

const LIMIT = 4

type Stack struct {
	data [LIMIT]int
	ix   int
}

func (s *Stack) Push(x int) {
	if s.ix == LIMIT {
		fmt.Println("Stack is already full")
		return // stack overflow
	}
	s.data[s.ix] = x
	s.ix++ // increment index of stack element present
}

func (s *Stack) Pop() int {
	if s.ix > 0 { // stack not empty
		s.ix-- // decrement index of stack element present
		element := s.data[s.ix]
		s.data[s.ix] = 0
		return element
	}
	return -1 // stack is empty
}

func (s Stack) String() string {
	str := ""
	for i := 0; i < s.ix; i++ {
		str += "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}
