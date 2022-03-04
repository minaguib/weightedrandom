package container_test

import (
	"fmt"

	"github.com/minaguib/weightedrandom/container"
)

func Example() {

	container := container.New[string]()

	container.Add("Red", 500)
	container.Add("Blue", 250)
	container.Add("Green", 125)
	container.Add("Yellow", 120)
	container.Add("Transparent", 4)
	container.Add("Vantablack", 1)

	// Simulate 1,000 marble picks
	for i := 0; i < 1000; i++ {
		marble, _ := container.Pick()
		fmt.Println("Picked: ", marble)
	}

}
