package jenerics_test

import (
	"strconv"
	"testing"

	. "github.com/barisere/jenerics"
	"github.com/barisere/jenerics/specs"
)

func TestSliceWithIteratorSpecs(t *testing.T) {
	ints := Slice[int]{1, 2, 3, 4, 5}
	specs.Test_iterator_satifies_functor_laws(t, ints.Iter(), Plus[int], strconv.Itoa)
	specs.Test_filter_drops_non_matching_items(t, ints.Iter(), IsOdd[int])
	specs.Next_visits_all_its_elements_of_finite_iterator[int](t, ints.Iter(), 5)
	specs.Zipped_iterators_produce_only_as_many_as_the_shortest[int, string](
		t,
		ints.Iter(),
		Slice[string]{"one", "two", "three"}.Iter(),
	)
	specs.Zipped_iterators_produce_only_as_many_as_the_shortest[int](
		t,
		ints.Iter(),
		Zip[string, string](
			Slice[string]{"I", "II", "III", "IV"}.Iter(),
			Slice[string]{"one", "two", "three"}.Iter(),
		),
	)
}
