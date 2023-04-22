package main

import (
	"fmt"
	"sync"
)

// waitgroup
// counter
func main() {

	wg := new(sync.WaitGroup)

	strs := []string{"Ashok", "Khalil", "Kishore", "Jiten", "Nithin", "Himasaili", "Ranjani"}

	for _, s := range strs {
		wg.Add(1)
		go func(wg *sync.WaitGroup, str string) {
			rev := ReverseStr(str)
			fmt.Println(rev)
			wg.Done()
		}(wg, s)
	}

	wg.Wait()
	//runtime.Goexit()
}

func ReverseStr(str string) string {
	rstr := ""
	for _, s := range str {
		rstr = string(s) + rstr
	}
	return rstr
}

func Reverse(str string) {
	rstr := ""
	for _, s := range str {
		rstr = string(s) + rstr
	}
	println(rstr)
}
