package main

import "fmt"

func main(){
	var names []string
	names = append(names, "Steven")
	names = append(names, "Christopher")
	names = append(names, "Jackson")

	fmt.Println(names[0])
	fmt.Println(names)
}
