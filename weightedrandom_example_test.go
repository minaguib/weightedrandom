package weightedrandom_test

import (
	"fmt"
	"math/rand"

	"github.com/minaguib/weightedrandom"
)

func Example_jarOfMarbles() {

	marbles := []struct {
		color  string
		weight uint
		picked uint
	}{
		{"Red", 500, 0},
		{"Blue", 250, 0},
		{"Green", 125, 0},
		{"Yellow", 120, 0},
		{"Transparent", 4, 0},
		{"Vantablack", 1, 0},
	}

	// Calculate slice of weights
	weights := make([]float64, len(marbles))
	for i, info := range marbles {
		weights[i] = float64(info.weight)
	}

	// Initialize a new *WeightedRandom
	wr, err := weightedrandom.New(weights)
	if err != nil {
		return
	}

	// For example/test purposes only, override internal randomizer with a deterministic one:
	wr.Rand = rand.New(rand.NewSource(99))

	// Pick 100,000 times and summarize
	for i := 0; i < 100000; i++ {
		idx := wr.Pick()
		marbles[idx].picked++
		//fmt.Printf("Picked: %v\n", marbles[idx].color)
	}

	// Output report
	for _, info := range marbles {
		fmt.Printf("Color %-11s weight=%3d picked=%5d times\n", info.color, info.weight, info.picked)
	}

	// Output:
	// Color Red         weight=500 picked=49999 times
	// Color Blue        weight=250 picked=24991 times
	// Color Green       weight=125 picked=12559 times
	// Color Yellow      weight=120 picked=11932 times
	// Color Transparent weight=  4 picked=  415 times
	// Color Vantablack  weight=  1 picked=  104 times

}

func Example_fiftyFifty() {

	weights := []float64{50, 50}
	wr, err := weightedrandom.New(weights)
	if err != nil {
		return
	}

	// For example/test purposes only, override internal randomizer with a deterministic one:
	wr.Rand = rand.New(rand.NewSource(99))

	// Pick 100,000 times and summarize
	picked := []uint{0, 0}
	for i := 0; i < 100000; i++ {
		idx := wr.Pick()
		picked[idx]++
	}

	fmt.Printf("Picked: %v\n", picked)

	// Output:
	// Picked: [50070 49930]

}
