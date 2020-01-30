package main

import (
	"fmt"
)

type List struct {
	len  int
	head *Item
	last *Item
}

type Item struct {
	value int
	next  *Item
}

func insert(data int, input *List) error {
	var nextItem *Item
	newItem := &Item{data, nextItem}
	if input.len == 0 {
		input.head = newItem
		input.len = input.len + 1
		input.last = newItem
	} else {
		input.len = input.len + 1
		input.last.next = newItem
		input.last = newItem
	}
	return nil
}

func displayValues(input *List) error {
	headPtr := input.head
	for input.head != nil {
		data := input.head
		fmt.Println(data.value)
		input.head = data.next
	}
	input.head = headPtr
	return nil
}

func insertSpecificPosition(refValue, value int, input *List) error {
	headPtr := input.head
	for input.head != nil {
		if input.head.value == refValue {
			newItem := &Item{value: value, next: input.head.next}
			input.len = input.len + 1
			input.head.next = newItem
			input.head = newItem.next
		} else {
			input.head = input.head.next
		}
	}
	input.head = headPtr
	return nil
}

func reverseLinkedList(input *List) error {
	var next, prev *Item
	for input.head != nil {
		next = input.head.next
		input.head.next = prev
		prev = input.head
		input.head = next
	}
	input.head = prev
	return nil
}

func main() {
	// Driver program to save the list structure.
	data := []int{10, 20, 30, 40}
	input := &List{}
	for _, each := range data {
		err := insert(each, input)
		if err != nil {
			return
		}
	}
	err := insertSpecificPosition(40, 60, input)
	if err != nil {
		return
	}
	err = reverseLinkedList(input)
	err = displayValues(input)
	if err != nil {
		return
	}
}
