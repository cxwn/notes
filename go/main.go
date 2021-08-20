package main

import "fmt"

type Name string

func main() {
	var nm Name = "ivan du"
	fmt.Println(nm)
	nm.setName()
	fmt.Println(nm)
}

func (n *Name) setName() {
	*n = "duruihong"
}