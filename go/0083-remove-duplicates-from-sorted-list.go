func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	n := head
	for n.Next != nil {
		if n.Val == n.Next.Val {
			n.Next = n.Next.Next
		} else {
			n = n.Next
		}
	}
	return head
}
