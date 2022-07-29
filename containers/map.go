package containers

type GoMap[K comparable, V any] map[K]V

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
