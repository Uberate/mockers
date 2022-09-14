package hash

import (
	"crypto/sha256"
	"fmt"
	"mockers/pkg/utils"
	"reflect"
)

// Hashable define the object is hash able, can generator hash str, and can compare by hash. And all Hashable
// implements can serialization and deserialization by Serialization() and Deserialization([]byte).
// The Serialization can return different value by more than one time invoked. But the result of Serialization should
// return same hash value after Deserialization([]byte) by other object in same type.
type Hashable interface {
	// Hash return a hash string, not limit the length, type and format. But by more invoke it, should return same hash
	// value. This file provided some hash function that you can use. The input is byte array.
	Hash() string

	// IsHash return true when object's hash value is equals input string.
	IsHash(string string) bool

	// Serialization return the bytes of Hashable value. The value is can deserialization to object by
	// Deserialization([]byte).
	Serialization() []byte

	// Deserialization will receive a byte array, and set the object value from the input. If Deserialization([]byte)
	// can't parse the bytes(such like different type of object to Serialization and Deserialization), should do nothing
	// and return an error.
	Deserialization([]byte) error
}

// IsHash return true when specify Hashable hash value is equals hashValue. But if origin is nil, return false directly.
// Because nil have no hash value. All Hashable can implement IsHash by this function.
//
// For example code:
//
// // OneHashable is implemented the Hashable interface.
// type OneHashable struct{}
//
// func (o OneHashable)Hash()string{...}
//
// func (o OneHashable)IsHash(hashValue string)bool{
//     return IsHash(o, hashValue)
// }
func IsHash(origin Hashable, hashValue string) bool {
	if origin == nil {
		return false
	}
	return origin.Hash() == hashValue
}

//--------------------------------------------------

// AutoHash will cal the has of a struct, and if the any fields of struct, they will be ignored in AutoHash(). And this
// function use reflect, in golang, is slowly. AutoHash will cal all field hash, and sum it. AutoHash will ignore
// struct Hash interface. But for field which in struct, if it was implement the Hashable interface, will invoke it hash
// function.
// And for other type, has a rule to cal the value. But some type can't generate the Hash value, it will panic. About
// the rule of type here:
// It can auto hash the type in :
// - int(8,16,32,64 and int)
// - uint(8,16,32,64 and uint)
// - bool
// - string
// - float(32,64)
// - any-struct(has no-hashable field)
// - array(the element is hashable)
// - pointer(the element is hashable)
// - any-interface( the element is hashable)
// else all of other type is un-hashable.
// and will ignore the tag "hash-ignore", it can use at un-hashable field.
//
// Note that, this function is use reflect, in golang, is slowly!
func AutoHash(object interface{}) string {
	// nil object return directly.
	if object == nil {
		return ""
	}

	return SHA224(reflectHash(object, 0))
}

func reflectHash(object interface{}, deep int) []byte {

	value := reflect.ValueOf(object)
	typ := reflect.TypeOf(object)
	// if the deep is zero, it will hash specify object, and ignore the Hashable of specify object.
	if deep != 0 {
		// try to convert to Hashable

		if v, ok := object.(Hashable); ok {
			return []byte(v.Hash())
		}
	}

	if typ.Kind() == reflect.Interface {
		value = value.Elem()
	}
	if typ.Kind() == reflect.Pointer {
		value = value.Elem()
		typ = typ.Elem()
	}

	switch value.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Bool, reflect.String, reflect.Float32, reflect.Float64:
		switch value.Kind() {
		case reflect.String:
			return []byte(value.String())
		case reflect.Uint:
			return utils.UIntToBytes(uint(value.Uint()))
		case reflect.Uint8:
			return utils.AnyNumberToBytes(value.Uint(), 8)
		case reflect.Uint16:
			return utils.AnyNumberToBytes(value.Uint(), 16)
		case reflect.Uint32:
			return utils.AnyNumberToBytes(value.Uint(), 32)
		case reflect.Uint64:
			return utils.AnyNumberToBytes(value.Uint(), 64)
		case reflect.Int:
			return utils.IntToBytes(int(value.Int()))
		case reflect.Int8:
			return utils.AnyNumberToBytes(uint64(value.Int()), 8)
		case reflect.Int16:
			return utils.AnyNumberToBytes(uint64(value.Int()), 16)
		case reflect.Int32:
			return utils.AnyNumberToBytes(uint64(value.Int()), 32)
		case reflect.Int64:
			return utils.AnyNumberToBytes(uint64(value.Int()), 64)
		case reflect.Bool:
			return utils.BoolToBytes(value.Bool())
		case reflect.Float32:
			return utils.Float32ToBytes(float32(value.Float()))
		case reflect.Float64:
			return utils.Float64ToBytes(value.Float())
		}
	case reflect.Array, reflect.Slice:
		var resBytes []byte
		length := value.Len()
		var index int
		for index = 0; index < length; index++ {
			subValue := value.Index(length).Interface()
			resBytes = append(resBytes, reflectHash(subValue, deep+1)...)
		}
		return resBytes
	case reflect.Struct:
		// iterator the fields

		var resBytes []byte
		var index int
		for index = 0; index < typ.NumField(); index++ {
			subType := typ.Field(index)
			if _, ok := subType.Tag.Lookup("ignorehash"); ok {
				return []byte{}
			}

			subValue := value.FieldByName(subType.Name)
			if subValue.CanInterface() {
				resBytes = append(resBytes, reflectHash(subValue.Interface(), deep+1)...)
			}

		}
		return resBytes
	default:
		panic(any(fmt.Sprintf("%s is nohashable", typ.Kind().String())))

	}

	return []byte{}
}

// ---------------------------------------------------------------------------------------------------------------------
// some hash functions here

// SHA256 return a hash value by sha256.Sum256
func SHA256(inputs []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(inputs))
}

// SHA224 return a hash value by sha256.Sum224
func SHA224(inputs []byte) string {
	return fmt.Sprintf("%x", sha256.Sum224(inputs))
}
