package jenerics

import (
	"reflect"
)

type GoMap[K comparable, V any] map[K]V

func (self GoMap[K, V]) KeyIterator() Iterator[K] {
	return keyIterator[K, V]{newIterState(self)}
}

func (self GoMap[K, V]) ValueIterator() Iterator[V] {
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

func (self keyIterator[K, V]) Clone() Iterator[K] {
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

func (self valueIterator[K, V]) Clone() Iterator[V] {
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

type orderedKey interface {
	comparable
	Ordered
}

func MinMapKey[K orderedKey, V Ordered](items GoMap[K, V]) K {
	return MinItem[K](items.KeyIterator())
}

func MinMapValue[K comparable, V Ordered](items GoMap[K, V]) V {
	return MinItem[V](items.ValueIterator())
}

func MaxMapKey[K orderedKey, V Ordered](items GoMap[K, V]) K {
	return MaxItem[K](items.KeyIterator())
}

func MaxMapValue[K comparable, V Ordered](items GoMap[K, V]) V {
	return MaxItem[V](items.ValueIterator())
}
