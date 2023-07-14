package randgen

import (
	"math/rand"
	"strings"
	"time"
)

// GenInt generates a random int number between 0 and max
func GenInt(max int) int {
	return newRand().Intn(max)
}

// GenInt8 generates a random int8 number between 0 and max
func GenInt8(max int8) int8 {
	return int8(newRand().Intn(int(max)))
}

// GenInt16 generates a random int16 number between 0 and max
func GenInt16(max int16) int16 {
	return int16(newRand().Intn(int(max)))
}

// GenInt32 generates a random int32 number between 0 and max
func GenInt32(max int32) int32 {
	return int32(newRand().Intn(int(max)))
}

// GenInt64 generates a random int64 number between 0 and max
func GenInt64(max int64) int64 {
	return int64(newRand().Intn(int(max)))
}

// GenUint generates a random uint number between 0 and max
func GenUint(max uint) uint {
	return uint(newRand().Intn(int(max)))
}

// GenUint8 generates a random uint8 number between 0 and max
func GenUint8(max uint8) uint8 {
	return uint8(newRand().Intn(int(max)))
}

// GenUint16 generates a random uint16 number between 0 and max
func GenUint16(max uint16) uint16 {
	return uint16(newRand().Intn(int(max)))
}

// GenUint32 generates a random uint32 number between 0 and max
func GenUint32(max uint32) uint32 {
	return uint32(newRand().Intn(int(max)))
}

// GenUint64 generates a random uint64 number between 0 and max
func GenUint64(max uint64) uint64 {
	return uint64(newRand().Intn(int(max)))
}

// GenFloat32 generates a random float32 number between 0 and max
func GenFloat32(max float32) float32 {
	return float32(newRand().Float32() * max)
}

// GenFloat64 generates a random float64 number between 0 and max
func GenFloat64(max float64) float64 {
	return newRand().Float64() * max
}

// GenString generates a random string with length n
// thanks to https://stackoverflow.com/users/1705598/icza
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func GenString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, newRand().Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = newRand().Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

// GenArrayInt generates an array of random int numbers between 0 and max
func GenArrayInt(arraySize int, max int) []int {
	arr := make([]int, arraySize)

	for i := range arr {
		number := GenInt(max)
		arr[i] = number
	}
	return arr
}

// GenArrayInt8 generates an array of random int8 numbers between 0 and max
func GenArrayInt8(arraySize int, max int8) []int8 {
	arr := make([]int8, arraySize)

	for i := range arr {
		number := GenInt8(max)
		arr[i] = number
	}
	return arr
}

// GenArrayInt16 generates an array of random int16 numbers between 0 and max
func GenArrayInt16(arraySize int, max int16) []int16 {
	arr := make([]int16, arraySize)

	for i := range arr {
		number := GenInt16(max)
		arr[i] = number
	}
	return arr
}

// GenArrayInt32 generates an array of random int32 numbers between 0 and max
func GenArrayInt32(arraySize int, max int32) []int32 {
	arr := make([]int32, arraySize)

	for i := range arr {
		number := GenInt32(max)
		arr[i] = number
	}
	return arr
}

// GenArrayInt64 generates an array of random int64 numbers between 0 and max
func GenArrayInt64(arraySize int, max int64) []int64 {
	arr := make([]int64, arraySize)

	for i := range arr {
		number := GenInt64(max)
		arr[i] = number
	}
	return arr
}

// GenArrayUint generates an array of random uint numbers between 0 and max
func GenArrayUint(arraySize int, max uint) []uint {
	arr := make([]uint, arraySize)

	for i := range arr {
		number := GenUint(max)
		arr[i] = number
	}
	return arr
}

// GenArrayUint8 generates an array of random uint8 numbers between 0 and max
func GenArrayUint8(arraySize int, max uint8) []uint8 {
	arr := make([]uint8, arraySize)

	for i := range arr {
		number := GenUint8(max)
		arr[i] = number
	}
	return arr
}

// GenArrayUint16 generates an array of random uint16 numbers between 0 and max
func GenArrayUint16(arraySize int, max uint16) []uint16 {
	arr := make([]uint16, arraySize)

	for i := range arr {
		number := GenUint16(max)
		arr[i] = number
	}
	return arr
}

// GenArrayUint32 generates an array of random uint32 numbers between 0 and max
func GenArrayUint32(arraySize int, max uint32) []uint32 {
	arr := make([]uint32, arraySize)

	for i := range arr {
		number := GenUint32(max)
		arr[i] = number
	}
	return arr
}

// GenArrayUint64 generates an array of random uint64 numbers between 0 and max
func GenArrayUint64(arraySize int, max uint64) []uint64 {
	arr := make([]uint64, arraySize)

	for i := range arr {
		number := GenUint64(max)
		arr[i] = number
	}
	return arr
}

// GenArrayFloat32 generates an array of random float32 numbers between 0 and max
func GenArrayFloat32(arraySize int, max float32) []float32 {
	arr := make([]float32, arraySize)

	for i := range arr {
		number := GenFloat32(max)
		arr[i] = number
	}
	return arr
}

// GenArrayFloat64 generates an array of random float64 numbers between 0 and max
func GenArrayFloat64(arraySize int, max float64) []float64 {
	arr := make([]float64, arraySize)

	for i := range arr {
		number := GenFloat64(max)
		arr[i] = number
	}
	return arr
}

// GenArrayString generates an array of random strings with length n
func GenArrayString(arraySize int, strMaxLen int) []string {
	arr := make([]string, arraySize)

	for i := range arr {
		strLen := GenInt(strMaxLen)
		arr[i] = GenString(strLen)
	}
	return arr
}

func newRand() *rand.Rand {
	seed := rand.NewSource(time.Now().UnixNano())
	return rand.New(seed)
}
