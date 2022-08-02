package jenerics_test

import (
	"strconv"
	"testing"

	. "github.com/barisere/jenerics"
	"github.com/barisere/jenerics/specs"
)

func TestMapIterator(t *testing.T) {
	sourceMap := GoMap[int, int]{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
	it := sourceMap.ValueIterator()
	specs.Test_iterator_satifies_functor_laws(t, it, Plus[int], strconv.Itoa)
	specs.Next_visits_all_its_elements_of_finite_iterator(t, it.Clone(), 5)
	specs.Zipped_iterators_produce_only_as_many_as_the_shortest(
		t,
		sourceMap.ValueIterator(),
		GoMap[string, string]{"I": "one", "II": "two", "III": "three"}.KeyIterator(),
	)
}
