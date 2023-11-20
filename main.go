package main

import (
	"dsa/dsa"
	"fmt"
)

func main() {
	fmt.Println("Welcome")

	dArr := dsa.NewDynamicArr(9)
	dArr.Add(3)
	dArr.Add(5)
	dArr.Add(8)
	dArr.Add(0)

	fmt.Println(dArr.String())
	dArr.RemoveAt(1)
	dArr.RemoveAt(0)
	fmt.Println(dArr.String())
	fmt.Println(dArr.Get(0))

	for i, v := range dArr.Value {
		fmt.Println(i, ") ", v)
	}
}
