package main

import "fmt"

func main() {
	thing := [5]float64{1, 2, 3, 4, 5}
	squared := square(thing)
	fmt.Println(squared)
}

func square(thing [5]float64) [5]float64 {
	for i := 0; i < len(thing); i++ {
		thing[i] = thing[i] * thing[i]
	}

	return thing
}
