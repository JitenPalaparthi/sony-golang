package main

import (
	"fmt"
	"time"
)

func main() {

	strs := []string{"Ashok", "Khalil", "Kishore", "Jiten", "Nithin", "Himasaili", "Ranjani"}

	// for _, s := range strs {
	// 	go Reverse(s)
	// }

	for _, s := range strs {
		//for i := 0; i < len(strs); i++ {
		go func(str string) {
			rev := ReverseStr(str)
			fmt.Println(rev)
		}(s)
	}

	// for _, s := range strs {
	// 	//for i := 0; i < len(strs); i++ {
	// 	go func(str *string) {
	// 		rev := ReverseStr(*str)
	// 		fmt.Println(rev)
	// 	}(&s)
	// }

	// go func() {
	// 	for _, s := range strs {
	// 		//for i := 0; i < len(strs); i++ {
	// go func(){
	// 		rev := ReverseStr(s)
	// 		fmt.Println(rev)
	//	}()
	// 	}
	// }()

	time.Sleep(time.Second * 1)
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

//package main

// func main() {
// 	str := []string{"khalil", "sony", "raju", "apple", "cat", "dog", "samsung", "khalil", "sony", "raju", "apple", "cat", "dog", "samsung"}
// 	//don't use goroutine in for loop as the gloal variable increment faster than execution of go ruoitne function and cause problem
// 	for _, s := range str {
// 		go func() {
// 			rev := Reverse(s)
// 			fmt.Println(rev)
// 		}()
// 	}
// 	time.Sleep(time.Second + 1)
// }
// func Reverse(str string) string {
// 	rstr := ""
// 	for _, s := range str {
// 		rstr = string(s) + rstr
// 	}
// 	return rstr
// }
