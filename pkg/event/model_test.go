package event

import (
	"fmt"
	"mockers/pkg/hash"
	"testing"
)

func TestModel_Hash(t *testing.T) {
	m := Model{}
	fmt.Println(hash.AutoHash(m))
	fmt.Println(hash.AutoHash(&m))

	m.Action = "test"
	fmt.Println(hash.AutoHash(m))
	fmt.Println(hash.AutoHash(&m))

	m.Action = ""
	fmt.Println(hash.AutoHash(m))
	fmt.Println(hash.AutoHash(&m))
	m.CreateTime = 234324
	fmt.Println(hash.AutoHash(m))
	fmt.Println(hash.AutoHash(&m))

}
