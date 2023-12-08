package main

import "fmt"

func main() {
	var p *[]int = new([]int)

	fmt.Println(cap(*p))

	//var v = make([]int, 10, 11)

	*p = append(*p, 33)
	fmt.Println(cap(*p))
	var x = *p

	(*p)[0] = 43

	*p = append(*p, 2)

	(*p)[0] = 666

	//*p = append(*p, 33)

	fmt.Println(*p)

	fmt.Println(x)
}
