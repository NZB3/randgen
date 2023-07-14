package randgen

import (
	"math/rand"
	"time"
)

func GenInt(maxInt int) int {
	return newRand().Intn(maxInt)
}

func GenInt8(maxInt int8) int8 {
	return int8(newRand().Intn(int(maxInt)))
}

func GenInt16(maxInt int16) int16 {
	return int16(newRand().Intn(int(maxInt)))
}

func GenInt32(maxInt int32) int32 {
	return int32(newRand().Intn(int(maxInt)))
}

func GenInt64(maxInt int64) int64 {
	return int64(newRand().Intn(int(maxInt)))
}

func GenUint(maxUint uint) uint {
	return uint(newRand().Intn(int(maxUint)))
}

func GenUint8(maxUint uint8) uint8 {
	return uint8(newRand().Intn(int(maxUint)))
}

func GenUint16(maxUint uint16) uint16 {
	return uint16(newRand().Intn(int(maxUint)))
}

func GenUint32(maxUint uint32) uint32 {
	return uint32(newRand().Intn(int(maxUint)))
}

func GenUint64(maxUint uint64) uint64 {
	return uint64(newRand().Intn(int(maxUint)))
}

func GenFloat32(maxFloat float32) float32 {
	return float32(newRand().Float32() * maxFloat)
}

func GenFloat64(maxFloat float64) float64 {
	return newRand().Float64() * maxFloat
}

func GenArrayInt(arraySize int, maxInt int) []int {
	arr := make([]int, arraySize)

	for i := range arr {
		number := GenInt(maxInt)
		arr[i] = number
	}
	return arr
}

func GenArrayInt8(arraySize int, maxInt int8) []int8 {
	arr := make([]int8, arraySize)

	for i := range arr {
		number := GenInt8(maxInt)
		arr[i] = number
	}
	return arr
}

func GenArrayInt16(arraySize int, maxInt int16) []int16 {
	arr := make([]int16, arraySize)

	for i := range arr {
		number := GenInt16(maxInt)
		arr[i] = number
	}
	return arr
}

func GenArrayInt32(arraySize int, maxInt int32) []int32 {
	arr := make([]int32, arraySize)

	for i := range arr {
		number := GenInt32(maxInt)
		arr[i] = number
	}
	return arr
}

func GenArrayInt64(arraySize int, maxInt int64) []int64 {
	arr := make([]int64, arraySize)

	for i := range arr {
		number := GenInt64(maxInt)
		arr[i] = number
	}
	return arr
}

func GenArrayUint(arraySize int, maxUint uint) []uint {
	arr := make([]uint, arraySize)

	for i := range arr {
		number := GenUint(maxUint)
		arr[i] = number
	}
	return arr
}

func GenArrayUint8(arraySize int, maxUint uint8) []uint8 {
	arr := make([]uint8, arraySize)

	for i := range arr {
		number := GenUint8(maxUint)
		arr[i] = number
	}
	return arr
}

func GenArrayUint16(arraySize int, maxUint uint16) []uint16 {
	arr := make([]uint16, arraySize)

	for i := range arr {
		number := GenUint16(maxUint)
		arr[i] = number
	}
	return arr
}

func GenArrayUint32(arraySize int, maxUint uint32) []uint32 {
	arr := make([]uint32, arraySize)

	for i := range arr {
		number := GenUint32(maxUint)
		arr[i] = number
	}
	return arr
}

func GenArrayUint64(arraySize int, maxUint uint64) []uint64 {
	arr := make([]uint64, arraySize)

	for i := range arr {
		number := GenUint64(maxUint)
		arr[i] = number
	}
	return arr
}

func GenArrayFloat32(arraySize int, maxFloat float32) []float32 {
	arr := make([]float32, arraySize)

	for i := range arr {
		number := GenFloat32(maxFloat)
		arr[i] = number
	}
	return arr
}

func GenArrayFloat64(arraySize int, maxFloat float64) []float64 {
	arr := make([]float64, arraySize)

	for i := range arr {
		number := GenFloat64(maxFloat)
		arr[i] = number
	}
	return arr
}

func newRand() *rand.Rand {
	seed := rand.NewSource(time.Now().UnixNano())
	return rand.New(seed)
}
