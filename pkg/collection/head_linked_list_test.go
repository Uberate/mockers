package collection

import (
	"fmt"
	"testing"
)

func TestHeadListHead_Append(t *testing.T) {
	tl := NewHeadListHead[int]()
	fmt.Printf("len is: %d\n", tl.Len())
	tl.Append(1)
	fmt.Printf("len is: %d\n", tl.Len())
	tl.Append(3)
	fmt.Printf("len is: %d\n", tl.Len())
	tl.Append(2)
	fmt.Printf("len is: %d\n", tl.Len())
	tl.Iterator(func(index uint64, data int) {
		fmt.Printf("index: [%d], value is: [%d]\n", index, data)
	})
	fmt.Printf("Get [%d] is [%d]\n", 0, tl.Index(0))
	fmt.Printf("Get [%d] is [%d]\n", 1, tl.Index(1))
	fmt.Printf("Get [%d] is [%d]\n", 2, tl.Index(2))
	tl.Remove(1)
	fmt.Println("Remove one: index: [1]")
	fmt.Printf("len is: %d\n", tl.Len())
	tl.Iterator(func(index uint64, data int) {
		fmt.Printf("index: [%d], value is: [%d]\n", index, data)
	})
	tl.Remove(0)
	fmt.Println("Remove one: index: [0]")
	fmt.Printf("len is: %d\n", tl.Len())
	tl.Iterator(func(index uint64, data int) {
		fmt.Printf("index: [%d], value is: [%d]\n", index, data)
	})
	tl.Remove(0)
	fmt.Println("Remove one: index: [0]")
	fmt.Printf("len is: %d\n", tl.Len())
	tl.Iterator(func(index uint64, data int) {
		fmt.Printf("index: [%d], value is: [%d]\n", index, data)
	})
}
