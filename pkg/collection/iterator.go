package collection

type Iterator[T any] interface {
	Next() T
	Pre() T
	HasNext() bool
	HasPre() bool
	Remove()
}
