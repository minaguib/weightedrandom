// Copyright 2020 Mina Naguib
// See accompanying LICENSE file for legalese
// See accompanying README.md for algorithm references and background

// Package weightedrandom implements extremely efficient weighted random picking
package weightedrandom

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	// ErrNoWeights indicates New was called with nil or empty weights
	ErrNoWeights = fmt.Errorf("Supplied weights can not be empty")
)

// WeightedRandom stores the initialized probability and alias table produced by New
type WeightedRandom struct {
	prob  []float64
	alias []int
	Rand  *rand.Rand
}

func average(weights []float64) float64 {
	sum := float64(0)
	for _, weight := range weights {
		sum += weight
	}
	return sum / float64(len(weights))
}

// New accepts as input a slice of float64 weights.  At least 1 weight must be given otherwise error ErrNoWeights is returned
// It computes an internal alias table and returns a *WeightedRandom which is then suitable for calling .Pick() on
//
func New(weights []float64) (*WeightedRandom, error) {

	n := len(weights)

	if n == 0 {
		return nil, ErrNoWeights
	}

	// Make a copy of weights as we'll mutate it later
	weights2 := make([]float64, n)
	copy(weights2, weights)
	weights = weights2

	wr := &WeightedRandom{
		prob:  make([]float64, n),
		alias: make([]int, n),
		Rand:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	average := average(weights)

	// Fan out weights indexes to small or large
	var small []int
	var large []int
	for i, weight := range weights {
		if weight >= average {
			large = append(large, i)
		} else {
			small = append(small, i)
		}
	}

	// Fan out small and large into prob and alias
	for len(small) > 0 && len(large) > 0 {
		small_i := small[0]
		small = small[1:]
		large_i := large[0]
		large = large[1:]

		wr.prob[small_i] = weights[small_i] / average
		wr.alias[small_i] = large_i
		weights[large_i] -= average - weights[small_i]

		if weights[large_i] < average {
			small = append(small, large_i)
		} else {
			large = append(large, large_i)
		}
	}

	// Any indexes remaining in small or large assume normalized average
	for _, small_i := range small {
		wr.prob[small_i] = 1.0
	}
	for _, large_i := range large {
		wr.prob[large_i] = 1.0
	}

	return wr, nil
}

// Pick picks a random element and returns its index (as per the order supplied when New was called)
func (wr *WeightedRandom) Pick() int {

	col := wr.Rand.Intn(len(wr.prob))

	if wr.Rand.Float64() < wr.prob[col] {
		return col
	} else {
		return wr.alias[col]
	}

}