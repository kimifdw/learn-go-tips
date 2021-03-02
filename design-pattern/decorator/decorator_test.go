package decorator

import (
	"testing"
)

func TestExampleDecorator(t *testing.T) {
	t.Run("decorator desigin test", func(t *testing.T) {
		o := &OriCalculate{1}
		m := &MutCalculate{o, 2}
		a := &AddCalculate{o, 3}

		t.Logf("Oricalculate is %d \n MutCalculate is %d\n AddCalculate is %d \n", o.Cal(), m.Cal(), a.Cal())
	})

}
