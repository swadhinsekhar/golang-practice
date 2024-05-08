package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func (ll *LinkedList) Add(val int) {
	node := NewListNode(val)
	if ll.Head == nil {
		ll.Head = node
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node
}

func (ll *LinkedList) Remove(val int) {
	if ll.Head == nil {
		return
	}
	if ll.Head.Val == val {
		ll.Head = ll.Head.Next
		return
	}
	current := ll.Head
	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

func RemoveMiddle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		head = nil
		return head
	}
	slow := head
	fast := head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if prev != nil {
		prev.Next = slow.Next
	} else {
		head = slow.Next
	}
	return head
}

/*
Input: head = [1,2,3,4,5]
Output:     = [5,4,3,2,1]
*/
func reverseList(head *ListNode) *ListNode {
	// creating new head in the beginning
	var prev *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev // prev = nil, current = 1 -> 2 -> nil
		prev = current      // prev = 1
		//fmt.Println("prev", prev.Val, "prev.Next", prev.Next)
		current = next
	}
	return prev
}

/*
Input: head = [1,2,3,4,5]
Output: [1,3,5,2,4]
*/
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	odd := head
	even := head.Next
	evenHead := even
	for even != nil && even.Next != nil {
		even.Next, odd.Next = even.Next.Next, even.Next
		odd, even = odd.Next, even.Next
	}
	odd.Next = evenHead
	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%v ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Code
	fmt.Println("Hello, World!")
	ll := LinkedList{}

	inputData := []int{}
	//lint:ignore SA1000 we want to make sure that no two results
	inputData = []int{2, 1, 3, 5, 6, 4, 7} //lint:ignore SA1000
	inputData = []int{1, 2, 3, 4, 5, 6, 7}
	for _, val := range inputData {
		ll.Add(val)
	}
	printList(ll.Head)
	//oddEvenList(ll.Head)
	//printList(ll.Head)

	//ll.Remove(5)
	//printList(ll.Head)

	//RemoveMiddle(ll.Head)
	//printList(ll.Head)

	ll.Head = reverseList(ll.Head)
	printList(ll.Head)
}
