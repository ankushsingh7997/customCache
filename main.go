package main

import "fmt"

const SIZE = 5

type Node struct {
	val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Cache struct {
	Queue Queue
	Hash  Hash
}
type Hash map[string]*Node

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}
	head.Right = tail
	tail.Left = head
	return Queue{Head: head, Tail: tail}

}

func (c *Cache) Check(str string) {
	node := &Node{}
	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{val: str}
	}

	c.Add(node)
	c.Hash[str] = node

}

func (c *Cache) Remove(n *Node) *Node {
	left := n.Left
	right := n.Right

	right.Left = left
	left.Right = right
	c.Queue.Length -= 1
	delete(c.Hash, n.val)
	return n

}
func (c *Cache) Add(n *Node) {
	temp := c.Queue.Head.Right
	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = temp
	temp.Left = n
	c.Queue.Length++

	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.Left)

	}

}

func (c *Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Println()
	fmt.Printf("%d - [", q.Length)

	for i := 0; i < q.Length; i++ {

		fmt.Printf("{%s}", node.val)

		node = node.Right
	}
	fmt.Printf("]")

}

func main() {
	fmt.Println("Start Cache")
	cache := NewCache()
	for _, word := range []string{"name", "class", "test", "testing", "ankush", "vishank"} {
		cache.Check(word)
		cache.Display()
	}

}
