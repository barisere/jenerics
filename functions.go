package jenerics

type Predicate[T any] func(T) bool
type Thunk[T any] func() T

func Id[T any](v T) T {
	return v
}

func Const[T any](value T) Thunk[T] {
	return func() T {
		return value
	}
}

func Compose[T, R1, R2 any](f func(R1) R2, g func(T) R1) func(T) R2 {
	return func(t T) R2 {
		return f(g(t))
	}
}

func DoTimes(count uint, do func(uint)) {
	for i := uint(0); i < count; i++ {
		do(i)
	}
}

func IsEven[T Integers](n T) bool {
	return n%2 == 0
}

func IsOdd[T Integers](n T) bool {
	return Not(IsEven[T])(n)
}

func Min[T Ordered](a, b T) T {
	if Compare(a, b) == OrderingLess {
		return a
	}
	return b
}

func Max[T Ordered](a, b T) T {
	if Compare(a, b) == OrderingGt {
		return a
	}
	return b
}

func Not[T any](pred Predicate[T]) Predicate[T] {
	return func(t T) bool {
		return !pred(t)
	}
}

func Plus[V Numeric](v V) V { return v + 1 }
