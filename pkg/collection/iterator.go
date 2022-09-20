package collection

type ReadonlyIterator[T any] interface {
	Next() T
	Pre() T
	HasNext() bool
	HasPre() bool
}

type Iterator[T any] interface {
	ReadonlyIterator[T]

	Remove()
}
