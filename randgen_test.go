package randgen_test

import (
	"fmt"
	"randgen"
	"testing"
)

func TestGenInt(t *testing.T) {
	fmt.Println(randgen.GenInt(10))
}

func TestGenInt8(t *testing.T) {
	fmt.Println(randgen.GenInt8(10))
}

func TestGenInt16(t *testing.T) {
	fmt.Println(randgen.GenInt16(10))
}

func TestGenInt32(t *testing.T) {
	fmt.Println(randgen.GenInt32(10))
}

func TestGenInt64(t *testing.T) {
	fmt.Println(randgen.GenInt64(10))
}

func TestGenUint(t *testing.T) {
	fmt.Println(randgen.GenUint(10))
}

func TestGenUint8(t *testing.T) {
	fmt.Println(randgen.GenUint8(10))
}

func TestGenUint16(t *testing.T) {
	fmt.Println(randgen.GenUint16(10))
}

func TestGenUint32(t *testing.T) {
	fmt.Println(randgen.GenUint32(10))
}

func TestGenUint64(t *testing.T) {
	fmt.Println(randgen.GenUint64(10))
}

func TestGenFloat32(t *testing.T) {
	max := float32(3.4e+38)
	want := randgen.GenFloat32(max) <= max
	if !want {
		t.Errorf("GenFloat32(%v) = %v, want %v", max, want, true)
	}
}

func TestGenFloat64(t *testing.T) {
	max := float64(1.7e+308)
	want := randgen.GenFloat64(max) <= max
	if !want {
		t.Errorf("GenFloat32(%v) = %v, want %v", max, want, true)
	}
}

func TestGenString(t *testing.T) {
	fmt.Println(randgen.GenString(10))
}

func TestGenArrayString(t *testing.T) {
	fmt.Println(randgen.GenArrayString(10, 10))
}

func TestGenArrayInt(t *testing.T) {
	fmt.Println(randgen.GenArrayInt(10, 10))
}

func TestGenArrayInt8(t *testing.T) {
	fmt.Println(randgen.GenArrayInt8(10, 10))
}

func TestGenArrayInt16(t *testing.T) {
	fmt.Println(randgen.GenArrayInt16(10, 10))
}

func TestGenArrayInt32(t *testing.T) {
	fmt.Println(randgen.GenArrayInt32(10, 10))
}

func TestGenArrayInt64(t *testing.T) {
	fmt.Println(randgen.GenArrayInt64(10, 10))
}

func TestGenArrayUint(t *testing.T) {
	fmt.Println(randgen.GenArrayUint(10, 10))
}

func TestGenArrayUint8(t *testing.T) {
	fmt.Println(randgen.GenArrayUint8(10, 10))
}

func TestGenArrayUint16(t *testing.T) {
	fmt.Println(randgen.GenArrayUint16(10, 10))
}

func TestGenArrayUint32(t *testing.T) {
	fmt.Println(randgen.GenArrayUint32(10, 10))
}

func TestGenArrayUint64(t *testing.T) {
	fmt.Println(randgen.GenArrayUint64(10, 10))
}

func TestGenArrayFloat32(t *testing.T) {
	fmt.Println(randgen.GenArrayFloat32(10, 10))
}

func TestGenArrayFloat64(t *testing.T) {
	fmt.Println(randgen.GenArrayFloat64(10, 10))
}
