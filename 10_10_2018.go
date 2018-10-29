package main

import (
	"fmt"
)

type iTable interface {
	GetName() string
}
type table struct{}

func (t table) GetName() string {
	return "11"
}
func t1(it iTable) {
	fmt.Printf(it.GetName())
}

type Vertex struct {
	X int
	Y int
}

func main1() {
	/**names := []string{"1", "2", "45"}
	//a := "n"
	// := &a
	p := &names
	for i, name := range names {
		fmt.Printf("%v %v %s\n", &p, &names[i], name)
	}
	t1(table{})*/
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
