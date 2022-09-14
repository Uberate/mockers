package utils

// platformBits return the system bits count, only support 64 and 32
// uint in 64 platform, the size of it is 64, ^uint(0) = 0xFFFFFFFF, >> 63 == 1
// 32 << (0xFFFF>>63) = 0 == 32 * 1 == 32
// uint in 32 platform, the size of it is 32, ^uint(0) = 0xFFFF, >> 63 == 0
// 32 << (0xFFFFFFFF>>63) = 1 == 32 * 2 == 64
var platformBits = 32 << (^uint(0) >> 63)

func Is32Platform() bool {
	return platformBits == 32
}

func Is64Platform() bool {
	return platformBits == 64
}
