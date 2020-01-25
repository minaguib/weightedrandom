// Package jar provides a convenience container API wrapping the underlying weightedrandom library
//
// Use Jar if:
//
// * You don't want to maintain your own slice/array to use the underlying weightedrandom library's index-based methods
//
// * You don't mind the performance overhead of casting your item to/from generic interface{}
//
package jar

import "github.com/minaguib/weightedrandom"

// Jar represents a jar of items (with weights attached to each)
type Jar struct {
	items   []interface{}
	weights []float64
	wr      *weightedrandom.WeightedRandom
}

// New returns a new empty Jar
func New() *Jar {
	return &Jar{}
}

// Add adds an item to the Jar along with its weight(probability)
func (jar *Jar) Add(item interface{}, weight float64) {
	jar.items = append(jar.items, item)
	jar.weights = append(jar.weights, weight)
	// Invalidate any previous alias table:
	jar.wr = nil
}

// Pick picks a random item previously Added to the Jar and returns it
// If no items were previously added to jar, an error is returned
func (jar *Jar) Pick() (interface{}, error) {
	if jar.wr == nil {
		wr, err := weightedrandom.New(jar.weights)
		if err != nil {
			return nil, err
		}
		jar.wr = wr
	}
	idx := jar.wr.Pick()
	return jar.items[idx], nil
}
