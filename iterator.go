package jenerics

type Iterator[T any] interface {
	Cloneable[Iterator[T]]
	Next() (value T, done bool)
}

type mapper[I, O any] struct {
	it Iterator[I]
	f  func(I) O
}

func (self mapper[I, O]) Next() (O, bool) {
	value, done := self.it.Next()
	return self.f(value), done
}

func (self mapper[I, O]) Clone() Iterator[O] {
	return mapper[I, O]{it: self.it.Clone(), f: self.f}
}

func Map[I, O any](it Iterator[I], f func(I) O) Iterator[O] {
	return mapper[I, O]{it, f}
}

func Collect[T any](it Iterator[T]) []T {
	values := make([]T, 0)
	for v, done := it.Next(); !done; v, done = it.Next() {
		values = append(values, v)
	}
	return values
}

type filter[T any] struct {
	it   Iterator[T]
	keep Predicate[T]
}

func (self filter[T]) Next() (value T, done bool) {
	for value, done = self.it.Next(); !done && !self.keep(value); value, done = self.it.Next() {
	}
	return value, done
}

func (self filter[T]) Clone() Iterator[T] {
	return filter[T]{it: self.it.Clone(), keep: self.keep}
}

func Filter[T any](it Iterator[T], predicate func(T) bool) Iterator[T] {
	return filter[T]{it, predicate}
}
