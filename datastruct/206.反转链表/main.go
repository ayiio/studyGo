package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode = nil
	var cur *ListNode = head
	for cur != nil {
		tmp := cur.Next // 缓存下节点
		cur.Next = pre  // 当前指向前一个
		// 向后移动
		pre = cur // 更新前一个
		cur = tmp // 更新当前节点
	}
	return pre
}

func main() {
	var list = &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{2, nil}}}}
	tmplist := list
	for list != nil {
		fmt.Print(list.Val, " ")
		list = list.Next
	}
	fmt.Println("\nreverse")
	newlist := reverseList(tmplist)
	for newlist != nil {
		fmt.Print(newlist.Val, " ")
		newlist = newlist.Next
	}
}
