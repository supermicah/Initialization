package priority

type ListNode struct {
	Val  int
	Next *ListNode
}

func toNode(node []int) *ListNode {
	if node == nil {
		return nil
	}
	n := &ListNode{Val: node[0]}
	if len(node) > 1 {
		n.Next = toNode(node[1:])
	}

	return n
}

type Queue []*ListNode

// Len 长度
func (q Queue) Len() int {
	return len(q)
}

// Less 比较
func (q Queue) Less(i, j int) bool {
	return q[i].Val < q[j].Val
}

// Swap 交换
func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

// Push 添加
func (q Queue) Push(d interface{}) {
	q = append(q, d.(*ListNode))
}

// Pop 获取
func (q Queue) Pop() interface{} {
	old := q
	n := len(q)
	x := old[n-1]
	q = old[:n-1]
	return x
}

func Run() {
}
