package main

import "fmt"

// Define the Node struct for a linked list
type Node struct {
	data int
	next *Node
}

// Function to print the linked list
func printList(node *Node) {
	for node != nil {
		fmt.Printf("%d -> ", node.data)
		node = node.next
	}
	fmt.Println("nil")
}

// Function to merge two sorted linked lists
func mergeLists(l1 *Node, l2 *Node) *Node {
	// Create a dummy node to act as the starting point
	dummy := &Node{}
	current := dummy

	// Traverse both lists and merge them in sorted order
	for l1 != nil && l2 != nil {
		if l1.data <= l2.data {
			current.next = l1
			l1 = l1.next
		} else {
			current.next = l2
			l2 = l2.next
		}
		current = current.next
	}

	// If one of the lists has remaining elements, append them
	if l1 != nil {
		current.next = l1
	} else {
		current.next = l2
	}

	// Return the next node from the dummy node as the merged list's head
	return dummy.next
}

// Function to create a new node
func newNode(data int) *Node {
	return &Node{data: data, next: nil}
}

func main() {
	// Create the first sorted linked list: 1 -> 3 -> 5
	l1 := newNode(1)
	l1.next = newNode(3)
	l1.next.next = newNode(5)

	// Create the second sorted linked list: 2 -> 4 -> 6
	l2 := newNode(2)
	l2.next = newNode(4)
	l2.next.next = newNode(6)

	fmt.Println("List 1:")
	printList(l1)

	fmt.Println("List 2:")
	printList(l2)

	// Merge the two lists
	mergedList := mergeLists(l1, l2)

	// Print the merged linked list
	fmt.Println("Merged List:")
	printList(mergedList)
}
