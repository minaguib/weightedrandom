package jar

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/minaguib/weightedrandom"
)

var jarMarblesColors = []interface{}{
	"Red",
	"Blue",
	"Green",
	"Yellow",
	"Transparent",
	"Vantablack",
}
var jarMarblesCounts = []float64{
	500,
	250,
	125,
	120,
	4,
	1,
}

func TestNew(t *testing.T) {
	jar := New()
	if jar.wr != nil {
		t.Errorf("jar.wr not nil initially")
	}
}

func TestPickEmpty(t *testing.T) {
	jar := New()
	picked, err := jar.Pick()
	if picked != nil || err == nil {
		t.Errorf("jar.Pick() did not return an error despite no added items")
	}
}

func TestPick(t *testing.T) {

	jar := New()
	for i, color := range jarMarblesColors {
		jar.Add(color, jarMarblesCounts[i])
	}

OUTER:
	for n := 0; n < 1000; n++ {
		picked, _ := jar.Pick()
		if jar.wr == nil {
			t.Errorf("jar.wr is nil after Pick")
		}
		for _, color := range jarMarblesColors {
			if picked == color {
				continue OUTER
			}
		}
		t.Errorf("Picked item %#v not in jarMarblesColors", picked)
	}

}

func TestAdd(t *testing.T) {

	jar := New()

	wr, _ := weightedrandom.New([]float64{99})

	for i, color := range jarMarblesColors {
		jar.wr = wr
		jar.Add(color, jarMarblesCounts[i])
		if jar.wr != nil {
			t.Errorf("jar.wr not reset to nil after Add")
		}
	}

	if diff := cmp.Diff(jarMarblesColors, jar.items); diff != "" {
		t.Errorf("jar.items mismatch (-want +got):\n%s", diff)
	}

	if diff := cmp.Diff(jarMarblesCounts, jar.weights); diff != "" {
		t.Errorf("jar.weights mismatch (-want +got):\n%s", diff)
	}

}
