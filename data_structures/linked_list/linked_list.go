package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func New(val int, next *Node) *Node {
	return &Node{
		val:  val,
		next: next,
	}
}

func getMiddle(head *Node) int {
	length := 0
	dummy := head

	for dummy != nil {
		length++
		dummy = dummy.next
	}

	for i := 0; i < length/2; i++ {
		head = head.next
	}

	return head.val
}

func isCycle(head *Node) bool {
	slow := head
	fast := head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if fast == slow {
			return true
		}
	}

	return false
}

func getMiddleFastSlow(head *Node) int {
	slow := head
	fast := head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow.val
}

func middleNode(head *Node) *Node {
	slow := head
	fast := head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}

func deleteDuplicates(head *Node) *Node {
	current := head

	for current != nil && current.next != nil {
		if current.val == current.next.val {
			current.next = current.next.next
		} else {
			current = current.next
		}
	}

	return head
}

func reverseLinkedList(head *Node) *Node {
	var prev *Node
	curr := head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	return prev
}

func printList(head *Node) {
	for head != nil {
		if head.next == nil {
			fmt.Printf("%v\n", head.val)
		} else {
			fmt.Printf("%v -> ", head.val)
		}

		head = head.next
	}
}

func main() {
	head := New(1, nil)
	second := New(2, nil)
	third := New(3, nil)
	fourth := New(4, nil)
	five := New(5, nil)
	six := New(5, nil)

	head.next = second
	second.next = third
	third.next = fourth
	fourth.next = five
	five.next = six
	//third.next = head

	fmt.Println(getMiddle(head))
	fmt.Println(getMiddleFastSlow(head))
	fmt.Println(isCycle(head))
	fmt.Println(middleNode(head))

	headD := New(1, New(1, New(2, nil)))
	fmt.Println("Delete Duplicates")
	printList(deleteDuplicates(headD))

	headDD := New(1, New(1, New(2, New(3, New(3, nil)))))
	printList(deleteDuplicates(headDD))

	fmt.Println("Reverse LL")
	printList(reverseLinkedList(head))
}
