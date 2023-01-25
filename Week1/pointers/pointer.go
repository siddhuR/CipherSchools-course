package main

import "fmt"

func main() {
	// i := 10
	// j := &i
	// *j = 100

	var i int
	i = 10
	//var j *int
	j := new(int)
	j = &i
	*j = 100
	//fmt.Println(j)
	//fmt.Println(i)

	name := new(string)
	*name = "golang"
	fmt.Println(*name)
}
