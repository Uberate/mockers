package utils

// This file provides some function to convert some value to byte array.

//--------------------------------------------------
// Number to bytes function here.

// IntToBytes return the int to byte array, and if platform is 32, use Int32ToBytes, if platform is 64,
// use Int64ToBytes. About the platform, see Is32Platform and Is64Platform.
func IntToBytes(x int) []byte {
	if Is32Platform() {
		return Int32ToBytes(int32(x))
	}
	return Int64ToBytes(int64(x))
}

// Int8ToBytes will return a []byte from input, because int8 is same as byte, so, the res.len is 1.
func Int8ToBytes(x int8) []byte {
	return AnyNumberToBytes(uint64(x), 8)
}

// Int16ToBytes will return a []byte from input, the res.len is 16.
func Int16ToBytes(x int16) []byte {
	return AnyNumberToBytes(uint64(x), 16)
}

// Int32ToBytes will return a []byte from input, the res.len is 32.
func Int32ToBytes(x int32) []byte {
	return AnyNumberToBytes(uint64(x), 32)
}

// Int64ToBytes will return a []byte from input, the res.len is 64.
func Int64ToBytes(x int64) []byte {
	return AnyNumberToBytes(uint64(x), 64)
}

// UIntToBytes return the int to byte array, and if platform is 32, use UInt32ToBytes, if platform is 64,
// use UInt64ToBytes. About the platform, see Is32Platform and Is64Platform.
func UIntToBytes(x uint) []byte {
	if Is32Platform() {
		return UInt32ToBytes(uint32(x))
	}
	return UInt64ToBytes(uint64(x))
}

// UInt8ToBytes will return a []byte from input, because int8 is same as byte, so, the res.len is 1.
func UInt8ToBytes(x uint8) []byte {
	return AnyNumberToBytes(uint64(x), 8)
}

// UInt16ToBytes will return a []byte from input, the res.len is 16.
func UInt16ToBytes(x uint16) []byte {
	return AnyNumberToBytes(uint64(x), 16)
}

// UInt32ToBytes will return a []byte from input, the res.len is 32.
func UInt32ToBytes(x uint32) []byte {
	return AnyNumberToBytes(uint64(x), 32)
}

// UInt64ToBytes will return a []byte from input, the res.len is 64.
func UInt64ToBytes(x uint64) []byte {
	return AnyNumberToBytes(x, 64)
}

// Float32ToBytes will return a []byte from input, the res.len is 32.
func Float32ToBytes(x float32) []byte {
	return AnyNumberToBytes(uint64(x), 32)
}

// Float64ToBytes will return a []byte from input, the res.len is 64.
func Float64ToBytes(x float64) []byte {
	return AnyNumberToBytes(uint64(x), 64)
}

var (
	trueBytes  = []byte{0b00000001}
	falseBytes = []byte{0b00000000}
)

// BoolToBytes return true or false hash, and false's hash is equals number(0).
func BoolToBytes(x bool) []byte {
	if x {
		return trueBytes
	} else {
		return falseBytes
	}
}

// RuneToBytes return a byte array, and rune can be seen as int32.
func RuneToBytes(x rune) []byte {
	return Int32ToBytes(x)
}

// AnyNumberToBytes will try to convert a number to a byte-array, the size is the value size, int64's size is 64.
// If size is zero, the AnyNumberToBytes will return an empty byte array.
func AnyNumberToBytes(x uint64, size int) []byte {
	if size <= 0 {
		// if size is less than zero, return empty byte array
		return []byte{}
	}
	length := size / 8
	if size%8 != 0 {
		length++
	}

	res := make([]byte, length, length)
	for i := length - 1; i >= 0; i-- {
		res[i] = byte(x)
		x = x << 8
	}

	return res
}

//--------------------------------------------------

// number array here

// IntArrayToBytes will return a byte array of input.
func IntArrayToBytes(a []int) []byte {
	if len(a) == 0 {
		return []byte{}
	}
	if Is32Platform() {
		res := make([]byte, len(a)*4, len(a)*4)
		for index, int32v := range a {
			res[index*4] = byte(int32v)
			res[index*4+1] = byte(int32v << (1 * 8))
			res[index*4+2] = byte(int32v << (2 * 8))
			res[index*4+3] = byte(int32v << (3 * 8))
		}

		return res
	}
	// is 64 platform
	res := make([]byte, len(a)*8, len(a)*8)
	for index, int64v := range a {
		res[index*8] = byte(int64v)
		res[index*8+1] = byte(int64v << (1 * 8))
		res[index*8+2] = byte(int64v << (2 * 8))
		res[index*8+3] = byte(int64v << (3 * 8))
		res[index*8+4] = byte(int64v << (4 * 8))
		res[index*8+5] = byte(int64v << (5 * 8))
		res[index*8+6] = byte(int64v << (6 * 8))
		res[index*8+7] = byte(int64v << (7 * 8))

	}

	return res
}

// Int8ArrayToBytes will return a byte array of input.
func Int8ArrayToBytes(a []int8) []byte {
	if len(a) == 0 {
		return []byte{}
	}
	res := make([]byte, len(a), len(a))
	for index, int8v := range a {
		res[index] = byte(int8v)
	}

	return res
}

// Int16ArrayToBytes will return a byte array of input.
func Int16ArrayToBytes(a []int16) []byte {
	if len(a) == 0 {
		return []byte{}
	}
	res := make([]byte, len(a)*2, len(a)*2)
	for index, int16v := range a {
		res[index*2] = byte(int16v)
		res[index*2+1] = byte(int16v << 8)
	}

	return res
}

// Int32ArrayToBytes will return a byte array of input.
func Int32ArrayToBytes(a []int32) []byte {
	if len(a) == 0 {
		return []byte{}
	}
	res := make([]byte, len(a)*4, len(a)*4)
	for index, int32v := range a {
		res[index*4] = byte(int32v)
		res[index*4+1] = byte(int32v << (1 * 8))
		res[index*4+2] = byte(int32v << (2 * 8))
		res[index*4+3] = byte(int32v << (3 * 8))
	}

	return res
}

// Int64ArrayToBytes will return a byte array of input.
func Int64ArrayToBytes(a []int64) []byte {
	if len(a) == 0 {
		return []byte{}
	}
	res := make([]byte, len(a)*8, len(a)*8)
	for index, int64v := range a {
		res[index*8] = byte(int64v)
		res[index*8+1] = byte(int64v << (1 * 8))
		res[index*8+2] = byte(int64v << (2 * 8))
		res[index*8+3] = byte(int64v << (3 * 8))
		res[index*8+4] = byte(int64v << (4 * 8))
		res[index*8+5] = byte(int64v << (5 * 8))
		res[index*8+6] = byte(int64v << (6 * 8))
		res[index*8+7] = byte(int64v << (7 * 8))

	}

	return res
}

// UIntArrayToBytes will return a byte array of input.
func UIntArrayToBytes(a []uint) []byte {
	if len(a) == 0 {
		return []byte{}
	}
	if Is32Platform() {
		res := make([]byte, len(a)*4, len(a)*4)
		for index, int32v := range a {
			res[index*4] = byte(int32v)
			res[index*4+1] = byte(int32v << (1 * 8))
			res[index*4+2] = byte(int32v << (2 * 8))
			res[index*4+3] = byte(int32v << (3 * 8))
		}

		return res
	}
	// is 64 platform
	res := make([]byte, len(a)*8, len(a)*8)
	for index, int64v := range a {
		res[index*8] = byte(int64v)
		res[index*8+1] = byte(int64v << (1 * 8))
		res[index*8+2] = byte(int64v << (2 * 8))
		res[index*8+3] = byte(int64v << (3 * 8))
		res[index*8+4] = byte(int64v << (4 * 8))
		res[index*8+5] = byte(int64v << (5 * 8))
		res[index*8+6] = byte(int64v << (6 * 8))
		res[index*8+7] = byte(int64v << (7 * 8))

	}

	return res
}

// UInt8ArrayToBytes will return a byte array of input.
func UInt8ArrayToBytes(a []uint8) []byte {
	if len(a) == 0 {
		return []byte{}
	}
	res := make([]byte, len(a), len(a))
	for index, int8v := range a {
		res[index] = byte(int8v)
	}

	return res
}

// UInt16ArrayToBytes will return a byte array of input.
func UInt16ArrayToBytes(a []uint16) []byte {
	if len(a) == 0 {
		return []byte{}
	}
	res := make([]byte, len(a)*2, len(a)*2)
	for index, int16v := range a {
		res[index*2] = byte(int16v)
		res[index*2+1] = byte(int16v << 8)
	}

	return res
}

// UInt32ArrayToBytes will return a byte array of input.
func UInt32ArrayToBytes(a []uint32) []byte {
	if len(a) == 0 {
		return []byte{}
	}
	res := make([]byte, len(a)*4, len(a)*4)
	for index, int32v := range a {
		res[index*4] = byte(int32v)
		res[index*4+1] = byte(int32v << (1 * 8))
		res[index*4+2] = byte(int32v << (2 * 8))
		res[index*4+3] = byte(int32v << (3 * 8))
	}

	return res
}

// UInt64ArrayToBytes will return a byte array of input.
func UInt64ArrayToBytes(a []uint64) []byte {
	if len(a) == 0 {
		return []byte{}
	}
	res := make([]byte, len(a)*8, len(a)*8)
	for index, int64v := range a {
		res[index*8] = byte(int64v)
		res[index*8+1] = byte(int64v << (1 * 8))
		res[index*8+2] = byte(int64v << (2 * 8))
		res[index*8+3] = byte(int64v << (3 * 8))
		res[index*8+4] = byte(int64v << (4 * 8))
		res[index*8+5] = byte(int64v << (5 * 8))
		res[index*8+6] = byte(int64v << (6 * 8))
		res[index*8+7] = byte(int64v << (7 * 8))

	}

	return res
}

// Float32ArrayToBytes will return a byte array of input.
func Float32ArrayToBytes(f32s []float32) []byte {
	if len(f32s) == 0 {
		return []byte{}
	}
	res := make([]byte, len(f32s)*4, len(f32s)*4)
	for index, f32v := range f32s {
		res[index*4] = byte(f32v)
		res[index*4+1] = byte(int32(f32v) << (1 * 8))
		res[index*4+2] = byte(int32(f32v) << (2 * 8))
		res[index*4+3] = byte(int32(f32v) << (3 * 8))
	}

	return res
}

// Float64ArrayToBytes will return a byte array of input.
func Float64ArrayToBytes(f64s []float64) []byte {
	if len(f64s) == 0 {
		return []byte{}
	}
	res := make([]byte, len(f64s)*8, len(f64s)*8)
	for index, f64v := range f64s {
		res[index*8] = byte(f64v)
		res[index*8+1] = byte(int64(f64v) << (1 * 8))
		res[index*8+2] = byte(int64(f64v) << (2 * 8))
		res[index*8+3] = byte(int64(f64v) << (3 * 8))
		res[index*8+4] = byte(int64(f64v) << (4 * 8))
		res[index*8+5] = byte(int64(f64v) << (5 * 8))
		res[index*8+6] = byte(int64(f64v) << (6 * 8))
		res[index*8+7] = byte(int64(f64v) << (7 * 8))
	}

	return res
}

// BoolArrayToBytes will return a byte array of input.
func BoolArrayToBytes(bs []bool) []byte {
	if len(bs) == 0 {
		return []byte{}
	}

	length := len(bs) / 8
	if len(bs)%8 != 0 {
		length++
	}

	res := make([]byte, length, length)

	for index := range bs {
		if bs[index] {
			currentIndex := length - index/8 - 1
			res[currentIndex] = res[currentIndex] | (1 << (index % 8))
		}

	}
	return res
}

// RuneArrayToBytes will return a byte array of input
func RuneArrayToBytes(rs []rune) []byte {
	return Int32ArrayToBytes(rs)
}