package main

import "fmt"

type name struct {
	firstName string
	lastName  string
}

func (n *name) getName() string {
	return fmt.Sprintf("struct name %s %s", n.firstName, n.lastName)
}

type person struct {
	name
	sex string
}

func (p *person) getName() string {
	return fmt.Sprintf("struct person %s %s", p.firstName, p.lastName)
}

func main() {
	p := person{
		name: name{
			firstName: "Zhang",
			lastName:  "Wind",
		},
		sex: "boy",
	}
	fmt.Println(p.firstName)
	fmt.Println(p.getName())
	fmt.Println(p.name.getName())

	p.firstName = "Wang"

	fmt.Println(p.firstName)
	fmt.Println(p.getName())
	fmt.Println(p.name.getName())

}
