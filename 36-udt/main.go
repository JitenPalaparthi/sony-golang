package main

import "fmt"

func main() {

	//var mm1 myMap

	mm1 := make(myMap)

	mm1["Bangalore-1"] = 58056
	mm1["Bangalore-2"] = 58066
	mm1["Trivandrum-1"] = 69001

	fmt.Println("keys", mm1.GetKeys())
	fmt.Println("Values", mm1.GetVals())

	m2 := make(map[string]any)

	m2["Jiten.p"] = 101
	m2["Rajesh.j"] = 102
	m2["Khalil.A"] = 103

	//myMap(m2).GetKeys()

	fmt.Println("keys", myMap(m2).GetKeys())
	fmt.Println("Values", myMap(m2).GetVals())

}

type myMap map[string]any

func (mm myMap) GetKeys() []string {

	keys := make([]string, len(mm))
	index := 0
	for key, _ := range mm {
		keys[index] = key
		index++
	}

	return keys
}

func (mm myMap) GetVals() []any {

	vals := make([]any, len(mm))
	index := 0
	for _, v := range mm {
		vals[index] = v
		index++
	}

	return vals
}
