package collection

type TreeNode[T any] struct {
	Data T
	Name string
	Subs map[string]TreeNode[T]
}

type TestInterface interface {
	Do(any2 any) any
}

type TestStruct[T any] struct {
}

func (ts *TestStruct[T]) Do(t T) any {
	return t
}
