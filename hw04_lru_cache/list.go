package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Prev  *ListItem
	Next  *ListItem
}

type list struct {
	head   *ListItem
	tail   *ListItem
	length int
}

func (l list) Len() int {
	return l.length
}

func (l list) Front() *ListItem {
	return l.head
}

func (l list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	node := &ListItem{
		Value: v,
		Next:  l.head,
		Prev:  nil,
	}

	if l.length == 0 {
		l.tail = node
	} else {
		l.head.Prev = node
	}
	l.head = node

	l.length++

	return node
}

func (l *list) PushBack(v interface{}) *ListItem {
	node := &ListItem{
		Value: v,
		Next:  nil,
		Prev:  l.tail,
	}

	if l.length == 0 {
		l.head = node
	} else {
		l.tail.Next = node
	}
	l.tail = node

	l.length++

	return node
}

func (l *list) Remove(i *ListItem) {
	switch {
	case l.length == 1:
		l.head = nil
		l.tail = nil
	case i == l.head:
		l.head = l.head.Next
		l.head.Prev = nil
	case i == l.tail:
		l.tail = l.tail.Prev
		l.tail.Next = nil
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}

	l.length--
}

func (l *list) MoveToFront(i *ListItem) {
	switch {
	case i == l.head:
		return
	case i == l.tail:
		l.tail = l.tail.Prev
		l.tail.Next = nil
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}

	i.Next = l.head
	i.Prev = nil
	l.head.Prev = i
	l.head = i
}

func NewList() List {
	return &list{}
}
