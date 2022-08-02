package containers_test

import (
	"strconv"
	"testing"

	j "github.com/barisere/jenerics"
	. "github.com/barisere/jenerics/containers"
	"github.com/barisere/jenerics/specs"
)

func TestMapIterator(t *testing.T) {
	sourceMap := GoMap[int, int]{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
	it := sourceMap.ValueIterator()
	specs.Test_iterator_satifies_functor_laws(t, it, j.Plus[int], strconv.Itoa)
	specs.Next_visits_all_its_elements_of_finite_iterator(t, it.Clone(), 5)
}
