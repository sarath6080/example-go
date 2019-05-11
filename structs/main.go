package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {

	jim := person{
		firstName: "sarath",
		lastName:  "kumar",
		contact: contactInfo{
			email: "sarath.com",
			zip:   600000,
		},
	}
	fmt.Println(jim)
}
