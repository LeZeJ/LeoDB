package skiplist

type Node struct{
	data interface{} //存储任意对象
	next []*Node //存储每一个level的下一个节点，数组索引代表level
}

//创建一个新的节点
func newNode(data interface{},height int)*Node{
	newnode:=new(Node) //地址
	newnode.data=data
	newnode.next=make([]*Node,height,height)

	return newnode
}

//获得第level层的下一个节点
func (node *Node)getNext(level int)*Node{
	return node.next[level]
}
