package main

import "fmt"

type Vertex struct {
	X int
	Z int
	Y string

}

func main(){
	i,j := 42, 36

	p := &i			// point to I
	fmt.Println("*p = ",*p)		//read I through the pointer
	fmt.Println("p = ", p)
	fmt.Println("&i = ", &i)
	*p = 21 			//set I through the pointer
	fmt.Println("i = ",i)		// see the new value if I

	g := &j				// point to j
	fmt.Println("g = ", g)
	fmt.Println("*g = ", *g)
	*g = *g / 3			// divide j through the pointer
	fmt.Println("j = ", j)		// see new value of j
	fmt.Println("g = ", g)
	fmt.Println("*g = ", *g)


	v := Vertex{4,2,"hello"}
	fmt.Println(v)
	ptr := &v
	ptr.X = 1e9
	fmt.Println(v)
	fmt.Println(*ptr)
	fmt.Println(ptr)


	var pow = []int{1,2,4,8,16,32,64,128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}