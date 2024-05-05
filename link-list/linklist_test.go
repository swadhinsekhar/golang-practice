package main

import "testing"

func TestReverseList(t *testing.T) {
	ll := LinkedList{}
	inputData := []int{1, 2, 3, 4, 5}
	for _, val := range inputData {
		ll.Add(val)
	}

	expectedData := []int{5, 4, 3, 2, 1}
	expectedHead := ll.Head
	for i := 0; i < len(expectedData); i++ {
		expectedHead.Val = expectedData[i]
		expectedHead = expectedHead.Next
	}

	reversedHead := reverseList(ll.Head)

	for i := 0; i < len(expectedData); i++ {
		if reversedHead.Val != expectedData[i] {
			t.Errorf("Expected %d, but got %d", expectedData[i], reversedHead.Val)
		}
		reversedHead = reversedHead.Next
	}
}
