package list

type Node[Element any] struct {
	element Element
	next    *Node[Element]
	prev    *Node[Element]
}

func NewNode[Element any](element Element) *Node[Element] {
	return &Node[Element]{
		element: element,
		next:    nil,
		prev:    nil,
	}
}

func (node *Node[Element]) Element() Element {
	return node.element
}

func (node *Node[Element]) Next() *Node[Element] {
	return node.next
}

func (node *Node[Element]) Prev() *Node[Element] {
	return node.prev
}

type List[Element any] struct {
	head *Node[Element]
	tail *Node[Element]
	size int
}

func NewList[Element any]() *List[Element] {
	return &List[Element]{}
}

func (list *List[Element]) Size() int {
	return list.size
}

func (list *List[Element]) First() *Node[Element] {
	return list.head
}

func (list *List[Element]) Last() *Node[Element] {
	return list.tail
}

func (list *List[Element]) AddFront(element Element) {
	node := NewNode(element)
	if list.head == nil {
		list.tail = node
	} else {
		list.head.next = node
		node.prev = list.head
	}
	list.head = node
	list.size++
}

func (list *List[Element]) AddBack(element Element) {
	node := NewNode(element)
	if list.tail == nil {
		list.head = node
	} else {
		list.tail.prev = node
		node.next = list.tail
	}
	list.tail = node
	list.size++
}

func (list *List[Element]) AddBefore(before *Node[Element], element Element) {
	if before == list.head {
		list.AddFront(element)
		return
	}
	node := NewNode(element)
	node.next = before.next
	node.prev = before
	before.next.prev = node
	before.next = node
	list.size++
}

func (list *List[Element]) AddAfter(after *Node[Element], element Element) {
	if after == list.tail {
		list.AddBack(element)
		return
	}
	node := NewNode(element)
	node.prev = after.prev
	node.next = after
	after.prev.next = node
	after.prev = node
	list.size++
}

func (list *List[Element]) RemoveFront() (Element, bool) {
	if list.size == 0 {
		var zero Element
		return zero, false
	}
	element := list.First().Element()
	list.head = list.head.prev
	list.head.next = nil
	list.size--
	return element, true
}

func (list *List[Element]) RemoveBack() (Element, bool) {
	if list.size == 0 {
		var zero Element
		return zero, false
	}
	element := list.Last().Element()
	list.tail = list.tail.next
	list.tail.prev = nil
	list.size--
	return element, true
}

func (list *List[Element]) RemoveNode(node *Node[Element]) Element {
	next := node.next
	prev := node.prev
	next.prev = prev
	prev.next = next
	list.size--
	return node.Element()
}

// To reverse a LinkedList we just detach each node and assemble a new list
// by going from the head to the tail and building the new list from the tail.
func (list *List[Element]) Reverse() {
	if list.Size() == 0 {
		return
	}
	newTail := list.head
	newHead := newTail
	iter := list.head.prev
	newTail.prev = nil
	for iter != nil {
		newHead.next = iter
		iter = iter.prev
		newHead.next.prev = newHead
		newHead = newHead.next
		newHead.next = nil
	}
	list.head = newHead
	list.tail = newTail
}

func (list *List[Element]) Data() []Element {
	data := make([]Element, 0, list.size)
	for iter := list.head; iter != nil; iter = iter.prev {
		data = append(data, iter.Element())
	}
	return data
}

func (list *List[Element]) DataReverse() []Element {
	data := make([]Element, 0, list.size)
	for iter := list.tail; iter != nil; iter = iter.next {
		data = append(data, iter.Element())
	}
	return data
}
