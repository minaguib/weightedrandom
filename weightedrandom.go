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

func duplicateAndAverage(weights []float64) (newWeights []float64, average float64) {
	newWeights = make([]float64, len(weights))
	sum := float64(0)
	for i, weight := range weights {
		sum += weight
		newWeights[i] = weight
	}
	average = sum / float64(len(weights))
	return newWeights, average
}

// New accepts as input a slice of float64 weights.  At least 1 weight must be given otherwise error ErrNoWeights is returned
// It computes an internal alias table and returns a *WeightedRandom which is then suitable for calling .Pick() on
//
func New(weights []float64) (*WeightedRandom, error) {

	n := len(weights)

	if n == 0 {
		return nil, ErrNoWeights
	}

	wr := &WeightedRandom{
		prob:  make([]float64, n),
		alias: make([]int, n),
		Rand:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	weights, average := duplicateAndAverage(weights)

	// Fan out weights indexes to small or large
	small := make([]int, 0, n)
	large := make([]int, 0, n)
	for i, weight := range weights {
		if weight >= average {
			large = append(large, i)
		} else {
			small = append(small, i)
		}
	}

	// Fan out small and large into prob and alias
	for len(small) > 0 && len(large) > 0 {
		smallIdx := small[0]
		small = small[1:]
		largeIdx := large[0]
		large = large[1:]

		wr.prob[smallIdx] = weights[smallIdx] / average
		wr.alias[smallIdx] = largeIdx
		weights[largeIdx] -= average - weights[smallIdx]

		if weights[largeIdx] < average {
			small = append(small, largeIdx)
		} else {
			large = append(large, largeIdx)
		}
	}

	// Any indexes remaining in small or large assume normalized average
	for _, smallIdx := range small {
		wr.prob[smallIdx] = 1.0
	}
	for _, largeIdx := range large {
		wr.prob[largeIdx] = 1.0
	}

	return wr, nil
}

// Pick picks a random element and returns its index (as per the order supplied when New was called)
func (wr *WeightedRandom) Pick() int {

	col := wr.Rand.Intn(len(wr.prob))

	if wr.Rand.Float64() < wr.prob[col] {
		return col
	}
	return wr.alias[col]

}
