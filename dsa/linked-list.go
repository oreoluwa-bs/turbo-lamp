package dsa

// A linked list is a sequential list of nodes that hold data which point to other nodes also containing data.

// [2] -> [3] -> [4]

// Used in lists, queue, stack, circular lists

// singly and doubly (2x memory than singly)

type LinkedNode struct {
	Value interface{}
	Next  *LinkedNode
}

func NewLinkedList() *LinkedNode {
	l := new(LinkedNode)
	return l
}

func (d *LinkedNode) Insert(value interface{}) {

}
