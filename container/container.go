// Package container is one of two convenience container APIs wrapping the underlying weightedrandom library
//
// Use Container if:
//
// * You don't want to maintain your own slice/array to use the underlying weightedrandom library's index-based methods
//
// * You're capable of using go 1.18 and generics
//
// * You prefer this generics convenience API to the sibling "jar" interface-based convenience API
//
package container

import "github.com/minaguib/weightedrandom"

// Container represents a container of items (with weights attached to each)
type Container [T any]struct {
	items   []T
	weights []float64
	wr      *weightedrandom.WeightedRandom
}

// New returns a new empty Container
func New[T any]() *Container[T] {
	return &Container[T]{}
}

// Add adds an item to the Container along with its weight(probability)
func (container *Container[T]) Add(item T, weight float64) {
	container.items = append(container.items, item)
	container.weights = append(container.weights, weight)
	// Invalidate any previous alias table:
	container.wr = nil
}

// Ensures we have an inner weightedrandom initialized
func (container *Container[T]) ensureWR() error {
	if container.wr != nil {
		return nil
	}
	wr, err := weightedrandom.New(container.weights)
	if err != nil {
		return err
	}
	container.wr = wr
	return nil
}

// Pick picks a random item previously Added to the Container and returns it
// If no items were previously added to container, an error is returned
func (container *Container[T]) Pick() (T, error) {
	if err := container.ensureWR(); err != nil {
		return *new(T), err
	}
	idx := container.wr.Pick()
	return container.items[idx], nil
}
