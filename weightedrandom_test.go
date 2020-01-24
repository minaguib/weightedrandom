package weightedrandom

import (
	"math/rand"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_Nil(t *testing.T) {
	wr, err := New(nil)
	if wr != nil || err != ErrNoWeights {
		t.Errorf("Should have errored on nil weights")
	}
}

func Test_Empty(t *testing.T) {
	wr, err := New([]float64{})
	if wr != nil || err != ErrNoWeights {
		t.Errorf("Should have errored on empty weights")
	}
}

func Test_Zero(t *testing.T) {
	wr, err := New([]float64{0})
	if err != nil {
		t.Errorf("Should not have returned an error %v", err)
	}
	// Should always return index 0
	for i := 0; i < 100000; i++ {
		idx := wr.Pick()
		if idx != 0 {
			t.Errorf("Expected pick to return index 0, got: %v", idx)
		}
	}
}

func Test_One(t *testing.T) {
	wr, err := New([]float64{99})
	if err != nil {
		t.Errorf("Should not have returned an error %v", err)
	}
	// Should always return index 0
	for i := 0; i < 100000; i++ {
		idx := wr.Pick()
		if idx != 0 {
			t.Errorf("Expected pick to return index 0, got: %v", idx)
		}
	}
}

// Independent invocations (default) should use new seeds and should not produce same picks
// Same-seed invocations should produce same picks
func Test_Random(t *testing.T) {

	weights := []float64{50, 50}

	// Independent instances (default)
	wrIndependent1, _ := New(weights)
	time.Sleep(100 * time.Millisecond)
	wrIndependent2, _ := New(weights)
	pickedIndependent1 := [100]int{}
	pickedIndependent2 := [100]int{}

	// Same-seed instances
	wrSame1, _ := New(weights)
	wrSame2, _ := New(weights)
	wrSame1.Rand = rand.New(rand.NewSource(1))
	wrSame2.Rand = rand.New(rand.NewSource(1))
	pickedSame1 := [100]int{}
	pickedSame2 := [100]int{}

	for i := 0; i < 100; i++ {
		pickedSame1[i] = wrSame1.Pick()
		pickedSame2[i] = wrSame2.Pick()
		pickedIndependent1[i] = wrIndependent1.Pick()
		pickedIndependent2[i] = wrIndependent2.Pick()
	}

	if !cmp.Equal(pickedSame1, pickedSame2) {
		t.Errorf("Two same-seed invocations produced different picks!")
	}

	if cmp.Equal(pickedIndependent1, pickedIndependent2) {
		t.Errorf("Two independent invocations produced the same 100 picks!")
	}

}

// More output distribution tests in weightedrandom_example_test.go
