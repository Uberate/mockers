package collection

type Collection[T any] interface {
	Iterator(func(index uint64, data T))
	Cursor(func(index uint64, data T))
}
