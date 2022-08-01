package containers_test

import (
	"strconv"
	"testing"

	j "github.com/barisere/jenerics"
	. "github.com/barisere/jenerics/containers"
	"github.com/barisere/jenerics/specs"
)

func TestSliceWithIteratorSpecs(t *testing.T) {
	ints := Slice[int]{1, 2, 3, 4, 5}
	specs.Test_iterator_satifies_functor_laws(t, ints.Iter(), j.Plus[int], strconv.Itoa)
	specs.Test_filter_drops_non_matching_items(t, ints.Iter(), j.IsOdd[int])
}
