// This package is used for the Go Crashcourse
package main

import (
	"flag"
	"fmt"

	"github.com/nii236/go-crashcourse/app/concurrency"
	"github.com/nii236/go-crashcourse/app/hello"
	"github.com/nii236/go-crashcourse/app/serve"
	"github.com/nii236/go-crashcourse/app/structs"
)

func main() {
	var cmd string

	flag.StringVar(&cmd, "cmd", cmd, `hello, serve1, serve2, routines, channels, structs, interfaces, wg or wwg`)
	flag.Parse()

	switch cmd {
	case "hello":
		fmt.Println("Running hello world!")
		hello.SayHello()
	case "serve1":
		fmt.Println("Running server!")
		serve.PartOne()
	case "serve2":
		fmt.Println("Running server!")
		serve.PartTwo()
	case "routines":
		concurrency.Routines()
	case "channels":
		concurrency.Channels()
	case "wg":
		fmt.Println("Running concurrency example with WaitGroups!")
		concurrency.WithWG(concurrency.Data)
	case "wwg":
		fmt.Println("Running concurrency example without WaitGroups!")
		concurrency.WithoutWG(concurrency.Data)
	case "sort":
		fmt.Println("Running sort interface example!")
		structs.SortHumans()
	case "structs":
		fmt.Println("Running structs example!")
		structs.StructExample()
	case "interfaces":
		fmt.Println("Running interfaces example!")
		s := &OldGuy{}
		structs.InterfaceExample(s)
	default:

		flag.Usage()
	}
}

//OldGuy is a struct showing how you can pass external interfaces in to a func
type OldGuy struct {
}

//SayHello is an external method to OldGuy
func (s OldGuy) SayHello() {
	fmt.Println("Hello from outside the package! I'm good for stubbing!")
}
