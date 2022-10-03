package collection

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTestStruct_Do(t *testing.T) {
	x := TestStruct[string]{}
	x2 := TestStruct[int]{}

	fmt.Println(reflect.TypeOf(x.Do("sdfa")))
	fmt.Println(reflect.TypeOf(x2.Do("alskdfjlakf")))
	//x.Do(19)

	//InterfaceFunc(&x)
}

func InterfaceFunc(t TestInterface) {
}
