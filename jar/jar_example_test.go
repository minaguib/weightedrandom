package jar_test

import (
	"fmt"

	"github.com/minaguib/weightedrandom/jar"
)

func Example() {

	jar := jar.New()

	jar.Add("Red", 500)
	jar.Add("Blue", 250)
	jar.Add("Green", 125)
	jar.Add("Yellow", 120)
	jar.Add("Transparent", 4)
	jar.Add("Vantablack", 1)

	// Simulate 1,000 marble picks
	for i := 0; i < 1000; i++ {
		marble, _ := jar.Pick()
		fmt.Println("Picked: ", marble)
	}

}
