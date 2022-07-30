package containers

import (
	"reflect"

	"github.com/barisere/jenerics"
)

type GoMap[K comparable, V any] map[K]V

func (self GoMap[K, V]) KeyIterator() jenerics.CloneableIterator[K] {
	return keyIterator[K, V]{newIterState(self)}
}

func (self GoMap[K, V]) ValueIterator() jenerics.CloneableIterator[V] {
	return valueIterator[K, V]{newIterState(self)}
}

type iterState[K comparable, V any] struct {
	source GoMap[K, V]
	iter   *reflect.MapIter
	done   bool
}

func newIterState[K comparable, V any](source GoMap[K, V]) iterState[K, V] {
	return iterState[K, V]{
		source: source,
		iter:   reflect.ValueOf(source).MapRange(),
	}
}

func (state iterState[K, V]) next() (K, V, bool) {
	var k K
	var v V
	state.done = !state.iter.Next()
	if state.done {
		return k, v, state.done
	}
	reflect.ValueOf(&k).Elem().Set(state.iter.Key())
	reflect.ValueOf(&v).Elem().Set(state.iter.Value())
	return k, v, state.done
}

type keyIterator[K comparable, V any] struct {
	iterState[K, V]
}

func (self keyIterator[K, V]) Clone() jenerics.Iterator[K] {
	return keyIterator[K, V]{newIterState(self.source)}
}

// Next implements Iterator
func (self keyIterator[K, V]) Next() (value K, done bool) {
	k, _, done := self.iterState.next()
	return k, done
}

type valueIterator[K comparable, V any] struct {
	iterState[K, V]
}

func (self valueIterator[K, V]) Clone() jenerics.Iterator[V] {
	return valueIterator[K, V]{newIterState(self.source)}
}

// Next implements Iterator
func (self valueIterator[K, V]) Next() (value V, done bool) {
	_, v, done := self.iterState.next()
	return v, done
}

func (self GoMap[K, V]) Collect() Slice[V] {
	values := make(Slice[V], 0, len(self))
	for _, v := range self {
		values = append(values, v)
	}
	return values
}

func MinMap[K comparable, V jenerics.Ordered](items GoMap[K, V]) V {
	return Min(items.Collect())
}

func MaxMap[K comparable, V jenerics.Ordered](items GoMap[K, V]) V {
	return Max(items.Collect())
}
