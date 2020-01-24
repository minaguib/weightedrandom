package weightedrandom

import (
	"fmt"
	"testing"
)

const maxWeights = 10000000
const multiple = 10

func makeWeights(numWeights int) []float64 {
	weights := make([]float64, numWeights)
	for i := 0; i < numWeights; i++ {
		weights[i] = float64(numWeights % 10000)
	}
	return weights
}

func BenchmarkNew(b *testing.B) {

	for numWeights := 1; numWeights <= maxWeights; numWeights *= multiple {
		weights := makeWeights(numWeights)
		b.Run(fmt.Sprintf("%v weights", numWeights), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				New(weights)
			}
		})

	}
}

func BenchmarkPick(b *testing.B) {

	for numWeights := 1; numWeights <= maxWeights; numWeights *= multiple {
		weights := makeWeights(numWeights)
		b.Run(fmt.Sprintf("from %v weights", numWeights), func(b *testing.B) {
			wr, _ := New(weights)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				wr.Pick()
			}
		})

	}
}
