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
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	first *ListItem
	last  *ListItem
	len   int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.first
}

func (l *list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	newListItem := ListItem{Value: v}
	newListItem.Next = l.first
	if l.len > 0 {
		l.first.Prev = &newListItem
	}
	l.first = &newListItem
	l.len++
	if l.len == 1 {
		l.last = &newListItem
	}

	return l.Front()
}

func (l *list) PushBack(v interface{}) *ListItem {
	newListItem := ListItem{Value: v}
	newListItem.Prev = l.last

	if l.len > 0 {
		l.last.Next = &newListItem
	}

	l.last = &newListItem
	l.len++
	if l.len == 1 {
		l.first = &newListItem
	}

	return l.Back()
}

func (l *list) Remove(i *ListItem) {
	switch {
	case i.Prev == nil && i.Next == nil:
		l.first = nil
		l.last = nil
	case i.Prev == nil:
		l.first = i.Next
		l.first.Prev = nil
	case i.Next == nil:
		l.last = i.Prev
		l.last.Next = nil
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.first == i {
		return
	}

	i.Prev.Next = i.Next
	i.Next = l.first
	i.Prev = nil
	l.first = i
}

func (l *list) MoveToBack(i *ListItem) {
	if l.last == i {
		return
	}

	i.Next.Prev = i.Prev
	i.Prev = l.last
	i.Next = nil
	l.last = i
}

func NewList() List {
	return &list{
		len:   0,
		first: nil,
		last:  nil,
	}
}
