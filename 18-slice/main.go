package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {

	slice1 := make([]int, 10)

	for i := 0; i < len(slice1); i++ {
		slice1[i] = rand.Intn(100)
	}

	fmt.Println(slice1)

	slice1, err := AddElem(slice1, 4, 400)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(slice1)

	slice2, err := AddElem(slice1, 18, 400)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(slice2)

}

func AddElem(slice []int, index, val int) ([]int, error) {
	if index >= len(slice) || index < 0 || slice == nil {
		//return nil, fmt.Errorf("either slice is nil or index out of bound")
		return nil, errors.New("either slice is nil or index out of bound")
	}
	slice1 := make([]int, index)
	copy(slice1, slice[:index])
	slice1 = append(slice1, val)
	slice = append(slice1, slice[index:]...)
	return slice, nil
}
