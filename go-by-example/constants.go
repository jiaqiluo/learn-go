package main

import (
	"fmt"
	"math"
)

const s string = "constant"

var a = "this is an variable"

func main() {
	fmt.Println(s)
	fmt.Println(a)

	const n = 50000000

	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

}
