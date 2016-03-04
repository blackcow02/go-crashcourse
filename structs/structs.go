package structs

import "fmt"

// PersonProvider is an interface which requires the SayHello func
type PersonProvider interface {
	SayHello()
}

// Person is a struct example of a person
type Person struct {
	Name string
	Age  int
}

// Developer is a struct with an embedded person
type Developer struct {
	Person
	Company string
}

// Student is a struct with an embedded person
type Student struct {
	Person
	University string
}

// SayHello is how a person says hello
func (p Person) SayHello() {
	fmt.Println(p.Name, ": Hello")
}

// SayHello is how a developer says hello
func (d Developer) SayHello() {
	fmt.Println(d.Name, ": BeepBoop")
}

// SayHello is how a student says hello
func (s Student) SayHello() {
	fmt.Println(s.Name, ": I need money")
}

// StructExample shows how structs can be embedded
func StructExample() {
	p := &Person{"Andy", 32}
	d := &Developer{Person{"John", 29}, "TFG"}
	s := &Student{Person{"Patrick", 18}, "UWA"}
	p.SayHello()
	d.SayHello()
	s.SayHello()
}

// InterfaceExample shows how you can use interfaces as args and stubbing
func InterfaceExample(pp PersonProvider) {
	pp.SayHello()
}
