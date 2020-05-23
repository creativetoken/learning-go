package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {

	alex := person{
		firstName: "abhilash",
		lastName:  "pk",
		contact: contactInfo{email: "abhiashpk@hotmail.com",
			zipCode: 492000,
		},
	}

	fmt.Printf("%+v", alex)

}
