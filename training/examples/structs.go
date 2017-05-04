package main

import "fmt"

type Person struct {
	FirstName string
	LastName string
}

func main(){
	me := Person{FirstName: "Steven", LastName: "Jackson"}
	fmt.Println(me)
}