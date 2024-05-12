package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
2130. Maximum Twin Sum of a Linked List
Given the head of a singly linked list, return the maximum possible twin sum of the list.
Input: head = [5,4,2,1]
Output: 6
Explanation:
	Nodes 0 and 1 are the twins of nodes 3 and 2, respectively. All have twin sum = 6.
	There are no other nodes with twins in the linked list.
	Thus, the maximum twin sum of the linked list is 6.
*/

func pairSum(head *ListNode) int {
	max := 0
	// solve using two pointers
	fast := head
	slow := head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	nex := slow.Next
	slow.Next = nil
	temp := reverseList(nex)
	for head != nil && temp != nil {
		max = Max(max, head.Val+temp.Val)
		head = head.Next
		temp = temp.Next
	}
	return max
}

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
func printLinkedList(head *ListNode) {
	for node := head; node != nil; node = node.Next {
		fmt.Printf("%v ", node.Val)
	}
	fmt.Println()
}

func main() {
	ll := &LinkedList{}
	inputData := []int{5, 4, 2, 1}
	for _, val := range inputData {
		ll.Add(val)
	}
	printLinkedList(ll.Head)
	fmt.Println("pairSum: ", pairSum(ll.Head))

	//inputData = []int{4, 2, 2, 3}
	//for _, val := range inputData {
	//	ll.Add(val)
	//}
	//printLinkedList(ll.Head)
	//fmt.Println("pairSum: ", pairSum(ll.Head))
}
