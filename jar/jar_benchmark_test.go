package jar

import (
	"fmt"
	"testing"
)

const maxWeights = 10000000
const multiple = 10

func makeJar(numWeights int) *Jar {
	jar := New()
	for i := 0; i < numWeights; i++ {
		jar.Add(i, float64(i%10000))
	}
	// Jar WR materialization is lazy, so force one:
	jar.Pick()
	return jar
}

func BenchmarkAdd(b *testing.B) {

	for numWeights := 1; numWeights <= maxWeights; numWeights *= multiple {
		b.Run(fmt.Sprintf("%v weights", numWeights), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				makeJar(numWeights)
			}
		})

	}
}

func BenchmarkPick(b *testing.B) {

	for numWeights := 1; numWeights <= maxWeights; numWeights *= multiple {
		jar := makeJar(numWeights)
		b.Run(fmt.Sprintf("from %v weights", numWeights), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				item, _ := jar.Pick()
				x := item.(int)
				x++
			}
		})

	}
}
