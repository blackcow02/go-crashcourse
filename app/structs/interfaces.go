package structs

import (
	"fmt"
	"sort"
	"strconv"
)

// Human represents a human
type Human struct {
	name  string
	age   int
	phone string
}

func (h Human) String() string {
	return "(name: " + h.name + " - age: " + strconv.Itoa(h.age) + " years)"
}

//HumanGroup is a type of slices that contain Humans
type HumanGroup []Human

func (g HumanGroup) Len() int {
	return len(g)
}

func (g HumanGroup) Less(i, j int) bool {
	if g[i].age < g[j].age {
		return true
	}
	return false
}

func (g HumanGroup) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

// SortHumans will sort the below humans
func SortHumans() {
	group := HumanGroup{
		Human{name: "Bart", age: 24},
		Human{name: "Bob", age: 23},
		Human{name: "Gertrude", age: 104},
		Human{name: "Paul", age: 44},
		Human{name: "Sam", age: 34},
		Human{name: "Jack", age: 54},
		Human{name: "Martha", age: 74},
		Human{name: "Leo", age: 4},
	}

	//Let's print this group as it is
	fmt.Println("The unsorted group is:")
	for _, v := range group {
		fmt.Println(v)
	}

	//Now let's sort it using the sort.Sort function
	sort.Sort(group)

	//Print the sorted group
	fmt.Println("\nThe sorted group is:")
	for _, v := range group {
		fmt.Println(v)
	}
}
