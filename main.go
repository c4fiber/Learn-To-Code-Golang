package main

import "fmt"

type hotdog int

type person struct {
	fname string
	lname string // 첫글자가 대문자면 패키지 밖에서 접근 가능
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning, James."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred."`)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	// basic example
	xi := []int{2, 4, 7, 9, 42}
	fmt.Println(xi)

	m := map[string]int{
		"Todd": 45,
		"Job":  42,
	}
	fmt.Println(m)

	// struct example
	p1 := person{
		"Miss",
		"Mondeypenny",
	}
	fmt.Println(p1)

	sal := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	fmt.Println(sal)

	// interface, implicitly implemented
	saySomething(p1)
	saySomething(sal)
}
