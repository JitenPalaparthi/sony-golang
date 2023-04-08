package main

import "fmt"

func main() {

	//var num int = 100

	//day := 1 // variable

	// const a any = 100

	const day = 1

	const pi float32 = 2.4 + 0.6 + 0.14

	const daysinaweek = 7 * day // this is of type int

	const ok = (1 == 1)

	const ok1 = (ok && false)

	const (
		pi1  = 4.14
		days = 7
	)

	fmt.Println(pi, daysinaweek, ok, ok1)

}
