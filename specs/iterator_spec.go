package specs

import (
	"sort"
	"testing"

	. "github.com/barisere/jenerics"
	"github.com/stretchr/testify/assert"
)

func Next_visits_all_its_elements_of_finite_iterator[T any](
	t *testing.T,
	it Iterator[T],
	expected_count uint,
) {
	all_elements := Collect(it)
	assert.Len(t, all_elements, int(expected_count))
}

func Test_iterator_satifies_functor_laws[T, R1 any, R2 Ordered](
	t *testing.T,
	it Iterator[T],
	g func(T) R1,
	f func(R1) R2,
) {
	t.Run("map f . map g == map (f . g)", func(t *testing.T) {
		test_composition_of_mappings(t, it, g, f)
	})
}

func test_composition_of_mappings[T, R1 any, R2 Ordered](
	t *testing.T,
	iterator Iterator[T],
	g func(T) R1,
	f func(R1) R2,
) {
	clone_a := iterator.Clone()
	clone_b := iterator.Clone()

	map_with_g := Map(clone_a, g)
	and_then_with_f := Map[R1](Slice[R1](Collect(map_with_g)).Iter(), f)

	map_with_compose_f_g := Map(clone_b, Compose(f, g))

	assertEqualIterators(t, map_with_compose_f_g, and_then_with_f)
}

func assertEqualIterators[T Ordered](t *testing.T, a Iterator[T], b Iterator[T]) {
	as := Slice[T](Collect(a))
	bs := Slice[T](Collect(b))
	sort.Slice(as, func(i, j int) bool { return as[i] < as[j] })
	sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
	assert.Equal(t, as, bs)
}

// In other words, `Filter(iterator, pred) ≍ Partition(iterator, pred).First`
func Test_filter_drops_non_matching_items[T any](
	t *testing.T,
	it Iterator[T],
	pred Predicate[T],
) {
	non_matches := Collect(Filter(it.Clone(), Not(pred)))
	matches := Collect(Filter(it, pred))

	Slice[T](matches).ForEach(func(item T) {
		assert.True(t, pred(item))
		assert.NotContains(t, non_matches, item, "matches and non matches must be disjoint")
	})
}

func Zipped_iterators_produce_only_as_many_as_the_shortest[T, U any](
	t *testing.T,
	ts Iterator[T],
	us Iterator[U],
) {
	lenTs := len(Collect(ts.Clone()))
	lenUs := len(Collect(us.Clone()))
	lenShorter := Min(lenTs, lenUs)
	zipped := Zip(ts, us)

	assert.Len(t, Collect(zipped), lenShorter)
}
