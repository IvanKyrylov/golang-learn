package taskfour

type List interface {
	Len() int                          // длина списка
	Front() *ListItem                  // первый элемент списка
	Back() *ListItem                   // последний элемент списка
	PushFront(v interface{}) *ListItem // добавить элемент в начале
	PushBack(v interface{}) *ListItem  // добавить элемент в конце
	Remove(i *ListItem)                // удалить элемент
	MoveToFront(i *ListItem)           // переместить элемент в начало
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len       int
	frontItem *ListItem
	backItem  *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	if l.len == 0 {
		return nil
	}

	return l.frontItem
}

func (l *list) Back() *ListItem {
	if l.len == 0 {
		return nil
	}

	return l.backItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	newFrontItem := new(ListItem)
	newFrontItem.Value = v

	if l.len == 0 {
		l.frontItem = newFrontItem
		l.backItem = newFrontItem
		l.len++

		return newFrontItem
	}

	newFrontItem.Next = l.frontItem

	oldFrontItem := l.frontItem
	oldFrontItem.Prev = newFrontItem

	l.frontItem = newFrontItem
	l.len++

	return newFrontItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newBackItem := new(ListItem)
	newBackItem.Value = v

	if l.len == 0 {
		l.frontItem = newBackItem
		l.backItem = newBackItem
		l.len++

		return newBackItem
	}

	newBackItem.Prev = l.backItem
	oldBackItem := l.backItem
	oldBackItem.Next = newBackItem

	l.backItem = newBackItem
	l.len++

	return newBackItem
}

func (l *list) Remove(i *ListItem) {
	if i == nil {
		return
	}

	if l.len == 0 {
		return
	}

	if (i.Next == nil && i.Prev == nil) && (l.frontItem != i && l.backItem != i) {
		return
	}

	switch {
	case (i.Next == nil && i.Prev == nil) && (l.frontItem == i && l.backItem == i):
		l.frontItem = nil
		l.backItem = nil
	case i.Next == nil:
		i.Prev.Next = nil
		l.backItem = i.Prev
	case i.Prev == nil:
		i.Next.Prev = nil
		l.frontItem = i.Next
	default:
		i.Next.Prev = i.Prev
		i.Prev.Next = i.Next
	}

	i.Next = nil
	i.Prev = nil
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == nil {
		return
	}

	if l.len == 0 {
		return
	}

	if i.Prev == nil {
		return
	}

	if i.Next == nil && l.backItem == i {
		i.Next = l.frontItem
		i.Prev.Next = nil
		l.backItem = i.Prev

		i.Prev = nil
		l.frontItem.Prev = i
		l.frontItem = i

		return
	}

	oldFrontItem := l.frontItem
	oldFrontItem.Prev = i
	l.frontItem = i

	prevItem := i.Prev
	nextItem := i.Next

	prevItem.Next = nextItem
	nextItem.Prev = prevItem

	i.Prev = nil
	i.Next = oldFrontItem
}

func NewList() List {
	return new(list)
}
