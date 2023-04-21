package main

import "fmt"

func main() {

	p1 := Person{Name: "Jiten", Email: "jitenp@outlook.com", Mobile: "9618558500", Address: Address{Line1: "Sriraj Apartments", City: "Bangalore", State: "Karnataka", Country: "India", PinCode: "560086"}}
	fmt.Println(p1, "City:", p1.Address.City)

	e1 := Employee{}
	e1.Name = "Jiten"
	e1.Email = "Jitenp@outlook.com"
	e1.Mobile = "9618558500"
	e1.City = "Bangalore"
	e1.State = "Karnataka"
	e1.Line1 = "Sriraj Aparments"
	e1.Country = "India"
	e1.PinCode = "560086"
	e1.Status = "Active"
	e1.Address.Status = "Active"

	fmt.Println("Employee details:", e1)
	fmt.Println(e1.GetAddress())
}

type Address struct {
	Line1, City, State, Country, PinCode, Status string
}

func (a Address) GetAddress() string {
	return fmt.Sprint(a)
}

type Person struct {
	Name, Email, Mobile string
	Address             Address // not a promoted field. Just embedded struct using composition
}

type Employee struct {
	Name, Email, Mobile, Status string
	Address                     // promoted field . Can acheive which is similar to inheritance
}
