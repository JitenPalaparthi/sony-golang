package main

import "fmt"

func main() {

	// var num int
	// const c constant
	// var arr [5]int
	// var slice []int

	var mp map[string]int // delare a map

	mp = make(map[string]int) // instantiate a map

	mp["guntur-1"] = 522001
	mp["bangalore-1"] = 560031
	mp["bangalore-2"] = 560033
	mp["bangalore-3"] = 560034
	mp["bangalore-4"] = 560035
	mp["trivandrum-1"] = 690011

	// for i,v:= range arr
	for key, val := range mp {
		fmt.Println("Key", key, "Value", val)
	}

	fmt.Println(mp)

	val := mp["guntur-2"]

	val2 := mp["helloworld"]

	fmt.Println("Value from guntur-2", val2)

	fmt.Println("Value from guntur-2", val)

	val, ok := mp["guntur-2"]

	if ok {
		fmt.Println("Key exists and the value is", val)
	} else {
		fmt.Println("Key does not exist you can create it")
		mp["guntur-2"] = 522002
	}

	fmt.Println(mp)

	mp1 := make(map[string]map[string]any)

	mp1["pincodes"] = map[string]any{"guntur": 522001, "trivandrum": 69011} // this is equal to if slice []int{1,2,3,4}
	mp1["countrycode"] = map[string]any{"India": "91", "Pakisthan": 92, "Singapore": 65}

	for k, v := range mp1 {
		fmt.Println("Key------------", k)
		for key, val := range v {
			fmt.Println("Key:", key, "Value", val)
		}
	}

	v := mp1["pincodes"]["guntur"]
	//v2 = v["gunrur"]

	fmt.Println(v)

	delete(mp1, "pincodes")

	v3, ok := mp1["pincodes"]
	if !ok {
		fmt.Println("Yes it got deleted")
	} else {
		fmt.Println("Not deleted", v3)
	}

}
