package jenerics

type Predicate[T any] func(T) bool

func Id[T any](v T) T {
	return v
}

func Compose[T, R1, R2 any](f func(R1) R2, g func(T) R1) func(T) R2 {
	return func(t T) R2 {
		return f(g(t))
	}
}

func Plus[V Numeric](v V) V { return v + 1 }
