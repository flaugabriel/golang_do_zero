package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)

	// println é uma interface do tipo generico
	fmt.Println(1, 2, "teste", false, true, float64(6516))

	// não é recomendoado usar assim.
	mapa := map[interface{}]interface{}{
		1:           "string",
		float32(23): true,
		true:        float64(51),
	}

	fmt.Println(mapa)
}
