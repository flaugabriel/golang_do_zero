package main

import(
	 "fmt"
	 "reflect"
)

func main(){
	var array1 [5]string

	fmt.Println(array1)

	array2 := [5]string{"p1", "p2", "p3", "p4", "p5"}
	fmt.Println(array2)
	array3 := [...]int{1,2,3,4,5}
	fmt.Println(array3)

	slice := []int{10,11,12,13,14,15,16}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "posição alterada"
	fmt.Println(array2)

	// arrays internos
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacity

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4)) // length
	fmt.Println(cap(slice4)) // capacity

}