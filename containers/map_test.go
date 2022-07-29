package containers_test

import (
	"strconv"
	"testing"

	j "github.com/barisere/jenerics"
	. "github.com/barisere/jenerics/containers"
	"github.com/barisere/jenerics/specs"
)

func TestMapIteratorsSatifiesFunctorLaws(t *testing.T) {
	it := GoMap[int, int]{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}.ValueIterator()
	specs.Test_iterator_satifies_functor_laws(t, it, j.Plus[int], strconv.Itoa)
}
