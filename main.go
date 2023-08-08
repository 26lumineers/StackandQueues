package main

import "fmt"

//! Stack
type Stack struct{
	item []int
}
// Push
func (s *Stack) Push(value int) {
	s.item = append(s.item,value)
}
// Pop
// Returns theh removed value
func (s *Stack) Pop() int{
	l := len(s.item)-1
	toRemove := s.item[l]
	fmt.Println("poping :: ",toRemove)
	s.item = s.item[:l]
	return toRemove
}

//!Queues
type Queue struct{
	item []int
}

//Enqueue
func(q *Queue) Enqueue(value int){
	q.item = append(q.item, value)
}
//Dequeue
func(q *Queue) Dequeue() int{
	toRemove := q.item[0]
	fmt.Printf("%d was dequeued\n",toRemove)
	q.item = q.item[1:]
	return toRemove
}
func main(){
	//! Implementing Stack
	// myStack := Stack{}
	// myStack.Push(5)
	// myStack.Push(10)
	// myStack.Push(20)
	// fmt.Println(myStack)
	// fmt.Println(len(myStack.item))
	// myStack.Pop()
	// fmt.Println(myStack)

	//! Implenting Queues
	myQueue := Queue{}
	myQueue.Enqueue(10)
	myQueue.Enqueue(20)
	myQueue.Enqueue(30)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)

}