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

func Zip[T, U any](ts Iterator[T], us Iterator[U]) Iterator[*Pair[T, U]] {
	return &zipper[T, U]{tIter: ts, uIter: us}
}

type zipper[T, U any] struct {
	tIter Iterator[T]
	uIter Iterator[U]
	tDone bool
	uDone bool
}

func (z *zipper[T, U]) Next() (value *Pair[T, U], done bool) {
	var t T
	var u U
	done = z.tDone || z.uDone
	if done {
		return
	}
	t, z.tDone = z.tIter.Next()
	u, z.uDone = z.uIter.Next()
	done = z.tDone || z.uDone
	value = &Pair[T, U]{t, u}
	return value, done
}

func (z zipper[T, U]) Clone() Iterator[*Pair[T, U]] {
	return &zipper[T, U]{
		tIter: z.tIter.Clone(),
		uIter: z.uIter.Clone(),
		tDone: z.tDone,
		uDone: z.uDone,
	}
}

type echoIterator[T any] struct {
	value   T
	done    bool
	refresh func()
}

func (e *echoIterator[T]) Next() (T, bool) {
	e.refresh()
	return e.value, e.done
}

type unzipper[T, U any] struct {
	source Iterator[Pair[T, U]]
}

func Unzip[T, U any](values Iterator[Pair[T, U]]) (ts Iterator[T], us Iterator[U]) {
	tEcho, uEcho := &echoIterator[T]{}, &echoIterator[U]{}
	refresh := func() {
		pair, done := values.Next()
		tEcho.value, tEcho.done = pair.First, done
		uEcho.value, uEcho.done = pair.Second, done
	}
	tEcho.refresh, uEcho.refresh = refresh, refresh
	return ts, us
}

func Some[T any](values Iterator[T], pred Predicate[T]) bool {
	filtered := Filter(values, pred)
	return len(Collect(filtered)) > 0
}

func MaxItem[T Ordered](items Iterator[T]) T {
	return lastInOrder(items, OrderingGt)
}

func lastInOrder[T Ordered](items Iterator[T], ordering Ordering) T {
	last, done := items.Next()
	for {
		if done {
			break
		}
		item, d := items.Next()
		if Compare(item, last) == ordering {
			last = item
		}
		done = d
	}
	return last
}

func MinItem[T Ordered](values Iterator[T]) T {
	return lastInOrder(values, OrderingLess)
}
