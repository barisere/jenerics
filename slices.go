package jenerics

type Slice[T any] []T

func (self Slice[T]) ForEach(f func(T)) {
	self.ForEachIdx(func(t T, _ int) { f(t) })
}

func (self Slice[T]) ForEachIdx(f func(T, int)) {
	for idx, value := range self {
		f(value, idx)
	}
}

func (self Slice[T]) Iter() Iterator[T] {
	return &forwardIterator[T]{slice: self}
}

type forwardIterator[T any] struct {
	slice Slice[T]
	pos   int
	done  bool
}

func (it forwardIterator[T]) Clone() Iterator[T] {
	return &forwardIterator[T]{it.slice[:], it.pos, it.done}
}

// Next implements Iterator
func (it *forwardIterator[T]) Next() (value T, done bool) {
	var t T
	it.done = it.pos >= len(it.slice)
	if it.done {
		return t, it.done
	}
	t = it.slice[it.pos]
	it.pos += 1
	return t, false
}

func (self Slice[T]) ToMap() GoMap[int, T] {
	var entries = make(GoMap[int, T], len(self))
	for i, item := range self {
		entries[i] = item
	}
	return entries
}

func Unique[T any, K comparable](s Slice[T], toKey func(T) K) Slice[T] {
	var entries = make(GoMap[K, T])
	for _, item := range s {
		entries[toKey(item)] = item
	}
	return entries.Collect()
}

func Fold[From, To any](s Slice[From], combine func(From, To) To, start To) To {
	for _, v := range s {
		start = combine(v, start)
	}
	return start
}

func FoldRight[From, To any](s Slice[From], combine func(From, To) To, start To) To {
	for i := len(s) - 1; i > 0; i-- {
		start = combine(s[i], start)
	}
	return start
}

func FoldM[From, To any, M Monoid[To]](s Slice[From], fn func(From) To, m M) To {
	transform := func(f From, t To) To {
		return m.Concat(fn(f), t)
	}
	return Fold(s, transform, m.Identity())
}
