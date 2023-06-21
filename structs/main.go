package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}

func main() {
	//hardik := person{firstname: "Hardik", lastname: "Bansal"}
	//var hardik person
	//hardik.firstname = "Hardik"
	//hardik.lastname = "Bansal"
	jim := person{firstname: "Jim", lastname: "JimJim", contact: contactInfo{email: "cdcd@gmail.com", zipCode: 123456}}
	//fmt.Printf("%+v", jim)
	//jimPointer := &jim
	jim.print()
	//jimPointer.UpdateName("Hard")
	jim.UpdateName("Hard")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointerToPerson *person) UpdateName(newName string) {
	pointerToPerson.firstname = newName
}
