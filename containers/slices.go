package containers

import (
	"github.com/barisere/jenerics"
)

type Slice[T any] []T

func (self Slice[T]) ForEach(f func(T)) {
	self.ForEachIdx(func(t T, _ int) { f(t) })
}

func (self Slice[T]) ForEachIdx(f func(T, int)) {
	for idx, value := range self {
		f(value, idx)
	}
}

func (self Slice[T]) Iter() jenerics.CloneableIterator[T] {
	return &forwardIterator[T]{slice: self}
}

type forwardIterator[T any] struct {
	slice Slice[T]
	pos   int
	done  bool
}

func (it forwardIterator[T]) Clone() jenerics.Iterator[T] {
	return &forwardIterator[T]{it.slice[:], it.pos, it.done}
}

// Next implements jenerics.Iterator
func (it *forwardIterator[T]) Next() (value T, done bool) {
	var t T
	if it.done {
		return t, it.done
	}
	t = it.slice[it.pos]
	it.pos += 1
	it.done = it.pos >= len(it.slice)
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

func Map[From, To any](s Slice[From], f func(From) To) Slice[To] {
	mapped := make(Slice[To], 0, len(s))
	for i, v := range s {
		mapped[i] = f(v)
	}
	return mapped
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

func DoTimes(count uint, do func(uint)) {
	for i := uint(0); i < count; i++ {
		do(i)
	}
}

func FoldM[From, To any, M jenerics.Monoid[To]](s Slice[From], fn func(From) To, m M) To {
	transform := func(f From, t To) To {
		return m.Concat(fn(f), t)
	}
	return Fold(s, transform, m.Identity())
}

func Max[T jenerics.Ordered](items Slice[T]) T {
	return lastInOrder(items, jenerics.OrderingLess)
}

func Min[T jenerics.Ordered](items Slice[T]) T {
	return lastInOrder(items, jenerics.OrderingGt)
}

func lastInOrder[T jenerics.Ordered](items Slice[T], ordering jenerics.Ordering) T {
	if len(items) == 1 {
		return items[0]
	}
	return Fold(items[1:], func(t1, t2 T) T {
		if jenerics.Compare(t1, t2) == ordering {
			return t2
		}
		return t1
	}, items[0])
}

type Tuple[T, U any] struct {
	T T
	U U
}

func Zip[T, U any](ts Slice[T], us Slice[U]) Slice[Tuple[T, U]] {
	result := make(Slice[Tuple[T, U]], 0, Min([]int{len(ts), len(us)}))
	for i := range result {
		result[i] = Tuple[T, U]{ts[i], us[i]}
	}
	return result
}

func Unzip[T, U any](values Slice[Tuple[T, U]]) (ts Slice[T], us Slice[U]) {
	for i := range values {
		ts = append(ts, values[i].T)
		us = append(us, values[i].U)
	}
	return ts, us
}
