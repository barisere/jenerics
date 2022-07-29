package containers_test

import (
	"strconv"
	"testing"

	j "github.com/barisere/jenerics"
	. "github.com/barisere/jenerics/containers"
	"github.com/barisere/jenerics/specs"
)

func TestSliceIteratorSatifiesFunctorLaws(t *testing.T) {
	it := Slice[int]{1, 2, 3, 4, 5}.Iter()
	specs.Test_iterator_satifies_functor_laws(t, it, j.Plus[int], strconv.Itoa)
}
