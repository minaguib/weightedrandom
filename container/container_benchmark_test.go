package container

import (
	"fmt"
	"testing"
)

const maxWeights = 10000000
const multiple = 10

func makeContainer(numWeights int) *Container[int] {
	container := New[int]()
	for i := 0; i < numWeights; i++ {
		container.Add(i, float64(i%10000))
	}
	// Container WR materialization is lazy, so force one:
	container.Pick()
	return container
}

func BenchmarkAdd(b *testing.B) {

	for numWeights := 1; numWeights <= maxWeights; numWeights *= multiple {
		b.Run(fmt.Sprintf("%v weights", numWeights), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				makeContainer(numWeights)
			}
		})

	}
}

func BenchmarkPick(b *testing.B) {

	for numWeights := 1; numWeights <= maxWeights; numWeights *= multiple {
		container := makeContainer(numWeights)
		b.Run(fmt.Sprintf("from %v weights", numWeights), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				item, _ := container.Pick()
				item++
			}
		})

	}
}
