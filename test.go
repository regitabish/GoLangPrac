package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	println("hello world!")
	var str string = "hello world from variable"
	println(str)

	/* for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			println(i, "偶数")
		} else {
			println(i, "奇数")
		}
	} */

	t := time.Now().UnixNano()
	rand.Seed(t)
	n := rand.Intn(6)
	fmt.Println(n)
}
