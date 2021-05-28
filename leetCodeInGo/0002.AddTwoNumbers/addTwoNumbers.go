
 type ListNode struct {
     Val int
      Next *ListNode
  }
  
  
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0;
	dummyHead := &ListNode{}
	cur := dummyHead

	for(l1 !=nil || l2 != nil || carry != 0){
		var v1,v2 int;
		if(l1!=nil){
			v1 = l1.Val;
			l1 = l1.Next;
		}else{
			v1 = 0;
		}

		if(l2 != nil){
			v2 = l2.Val;
			l2 = l2.Next;
		}else{
			v2 = 0;
		}

		cur.Next = &ListNode{Val:(v1+v2+carry)%10}
		cur = cur.Next;
		carry = (v1+v2+carry)/10;
	}
    return dummyHead.Next;
 }