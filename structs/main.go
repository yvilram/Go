package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	//contact contactInfo
	contactInfo
}

func main() {
	/* //yago := person{"Yago", "Vila"}
	yago := person{firstName: "Yago", lastName: "Vila"}
	fmt.Println(yago)
	*/

	/* var yago person
	yago.firstName = "Yago"
	yago.lastName = "Vila"
	*/

	/* yago := person {
		firstName : "Yago",
		lastName : "Vila",
		contact: contactInfo{
			email : "yago@example.com",
			zip : 12345,
		},
	} */

	yago := person{
		firstName: "Yago",
		lastName:  "Vila",
		contactInfo: contactInfo{
			email: "yago@example.com",
			zip:   12345,
		},
	}

	//yagoPointer := &yago
	yago.updateName("Rosamelano")
	yago.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(NewFirstName string) {
	(*pointerToPerson).firstName = NewFirstName
}
