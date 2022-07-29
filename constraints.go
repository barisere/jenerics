package jenerics

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Complex interface {
	~complex64 | ~complex128
}

type Numeric interface {
	Signed | Unsigned | Complex
}

type Ordering uint8

const (
	OrderingLess Ordering = iota
	OrderingEq   Ordering = iota
	OrderingGt   Ordering = iota
)

func Compare[T Ordered](a, b T) Ordering {
	switch true {
	case a < b:
		return OrderingLess
	case a == b:
		return OrderingEq
	default:
		return OrderingGt
	}
}

type Comparable[T any] interface {
	Compare(T, T) Ordering
}

type Cloneable[T any] interface {
	Clone() T
}

type Ordered interface {
	Signed | Unsigned | ~string
}
