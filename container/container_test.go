package container

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/minaguib/weightedrandom"
)

var colors = []string{
	"Red",
	"Blue",
	"Green",
	"Yellow",
	"Transparent",
	"Vantablack",
}
var weights = []float64{
	500,
	250,
	125,
	120,
	4,
	1,
}

func TestNew(t *testing.T) {
	container := New[string]()
	if container.wr != nil {
		t.Errorf("container.wr not nil initially")
	}
}

func TestPickEmpty(t *testing.T) {
	container := New[string]()
	picked, err := container.Pick()
	if picked != "" || err == nil {
		t.Errorf("container.Pick() did not return an error despite no added items")
	}
}

func TestPick(t *testing.T) {

	container := New[string]()
	for i := range colors {
		container.Add(colors[i], weights[i])
	}

OUTER:
	for n := 0; n < 1000; n++ {
		picked, _ := container.Pick()
		if container.wr == nil {
			t.Errorf("container.wr is nil after Pick")
		}
		for _, color := range colors {
			if picked == color {
				continue OUTER
			}
		}
		t.Errorf("Picked item %#v not in original supplied list", picked)
	}

}

func TestAdd(t *testing.T) {

	container := New[string]()

	wr, _ := weightedrandom.New([]float64{99})

	for i := range colors {

		// Assign canary wr, not useful
		container.wr = wr

		container.Add(colors[i], weights[i])

		//But assert nullification of prior (canary) wr
		if container.wr != nil {
			t.Errorf("container.wr not reset to nil after Add")
		}
		if container.items[len(container.items)-1] != colors[i] {
			t.Errorf("container.items last element is not last added")
		}
		if container.weights[len(container.weights)-1] != weights[i] {
			t.Errorf("container.weights last element is not last added")
		}
	}

	if diff := cmp.Diff(colors, container.items); diff != "" {
		t.Errorf("container.items mismatch (-want +got):\n%s", diff)
	}

	if diff := cmp.Diff(weights, container.weights); diff != "" {
		t.Errorf("container.weights mismatch (-want +got):\n%s", diff)
	}

}
