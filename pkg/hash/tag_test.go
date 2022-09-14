package hash

import (
	"fmt"
	"testing"
)

func TestTags(t1 *testing.T) {
	//--------------------------------------------------
	// test put
	tags := Tags{}
	tags.Put("test", "test")
	value, ok := tags.Get("test")
	if !ok || value != "test" {
		t1.Failed()
	}
	value, ok = tags.Get("nothing")
	if ok {
		t1.Failed()
	}
	if tags.Length() != 1 {
		t1.Failed()
	}

	//--------------------------------------------------
	// test cover
	tags.Put("test", "2")
	value, ok = tags.Get("test")
	if !ok || value != "2" {
		t1.Failed()
	}
	value, ok = tags.Get("nothing")
	if ok {
		t1.Failed()
	}
	if tags.Length() != 1 {
		t1.Failed()
	}

	//--------------------------------------------------
	// test delete
	tags.Delete("test")
	value, ok = tags.Get("test")
	if ok {
		t1.Failed()
	}
	if tags.Length() != 0 {
		t1.Failed()
	}

	//--------------------------------------------------
	// test more value
	for _, x := range []string{"1", "2", "3", "4"} {
		tags.Put(x, x)
	}
	if tags.Length() != 4 {
		t1.Failed()
	}
	for _, x := range []string{"1", "2", "3", "4"} {
		value, ok = tags.Get(x)
		if !ok || value != x {
			t1.Failed()
		}
	}
	tags.Delete("2")
	if tags.Length() != 3 {
		t1.Failed()
	}
	for _, x := range []string{"1", "3", "4"} {
		value, ok = tags.Get(x)
		if !ok || value != x {
			t1.Failed()
		}
	}
	value, ok = tags.Get("2")
	if ok {
		t1.Failed()
	}

	//--------------------------------------------------
	// test serialization and hash
	bytes := tags.Serialization()
	oldHash := tags.Hash()
	fmt.Printf("%s\n", string(bytes))
	tags.Put("test", "test")
	if tags.IsHash(oldHash) {
		t1.Failed()
	}
	err := tags.Deserialization(bytes)
	if err != nil {
		fmt.Println(err)
		t1.Failed()
	}
	value, ok = tags.Get("test")
	if ok {
		t1.Failed()
	}
	for _, x := range []string{"1", "3", "4"} {
		value, ok = tags.Get(x)
		if !ok || value != x {
			t1.Failed()
		}
	}
	if !tags.IsHash(oldHash) {
		t1.Failed()
	}
}