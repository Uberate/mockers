package collection

import "fmt"

func NewHeadListHead[T any]() *HeadListHead[T] {
	return &HeadListHead[T]{len: 0}
}

type HeadListHead[T any] struct {
	FirstEle *HeadListNode[T]
	LastEle  *HeadListNode[T]

	len uint64
}

func (hlh *HeadListHead[T]) Iterator(f func(index uint64, data T)) {
	np := hlh.FirstEle
	index := uint64(0)
	for np != nil {
		f(index, np.Data)
		np = np.Next
		index++
	}
}

func (hlh *HeadListHead[T]) Len() uint64 {
	return hlh.len
}

func (hlh *HeadListHead[T]) index(index uint64) *HeadListNode[T] {
	if index >= hlh.Len() || index < 0 {
		return nil
	}

	np := hlh.FirstEle
	i := uint64(0)
	for np != nil && i <= index {
		if i == index {
			return np
		}
		np = np.Next
		i++
	}
	return nil
}

func (hlh *HeadListHead[T]) Index(index uint64) T {
	node := hlh.index(index)
	if node == nil {
		panic(fmt.Errorf("can't find index: [%d]", index))
	}
	return node.Data
}

func (hlh *HeadListHead[T]) Remove(index uint64) {
	node := hlh.index(index)
	if node == nil {
		return
	}
	node.Remove()
	hlh.len--
}

func (hlh *HeadListHead[T]) Append(data T) {
	hlh.len++
	// is an empty list
	if hlh.LastEle == nil {
		s := NewHeadListNode[T](data)
		hlh.LastEle = s
		hlh.FirstEle = s
		s.Head = hlh
		return
	}
	hlh.LastEle.AppendNext(data)
}

func NewHeadListNode[T any](data T) *HeadListNode[T] {
	res := &HeadListNode[T]{}
	res.Data = data
	return res
}

type HeadListNode[T any] struct {
	Data T

	Pre  *HeadListNode[T]
	Next *HeadListNode[T]
	Head *HeadListHead[T]
}

// Remove will remove this node from list. If this node is the start in the list, reset HeadListHead.First to next.
// If this node is the ended in the list, reset HeadListHeader.LastEle to pre.
func (hln *HeadListNode[T]) Remove() *HeadListNode[T] {
	if hln.Pre == nil {
		// is the head
		hln.Head.FirstEle = hln.Next
	} else {
		hln.Pre.Next = hln.Next
	}
	if hln.Next == nil {
		// is the end
		hln.Head.LastEle = hln.Pre
	} else {
		hln.Next.Pre = hln.Pre
	}
	// to remove list info
	hln.Next = nil
	hln.Pre = nil
	hln.Head = nil
	return hln
}

// AppendPre will insert a data front of this node. And if this node is head, set HeadListHead.FirstEle to new node.
func (hln *HeadListNode[T]) AppendPre(data T) *HeadListNode[T] {
	s := NewHeadListNode[T](data)

	// set new node
	s.Next = hln
	s.Pre = hln.Pre

	// set pre and next value
	hln.Pre = s

	if s.Pre != nil {
		s.Pre.Next = s
	} else {
		// is head.
		hln.Head.FirstEle = s
	}
	s.Head = hln.Head
	return s
}

// AppendNext will insert a node after this node. If this node is ended, set HeadListHead.LastEle to new node.
func (hln *HeadListNode[T]) AppendNext(data T) *HeadListNode[T] {
	s := NewHeadListNode[T](data)

	s.Next = hln.Next
	s.Pre = hln

	hln.Next = s

	if s.Next != nil {
		s.Next.Pre = s
	} else {
		// is ended
		hln.Head.LastEle = s
	}

	s.Head = hln.Head
	return s
}
