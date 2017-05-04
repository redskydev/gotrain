package main

import "fmt"

func main(){
	// Basic Variable Declaration
	//  First, declaring the type
	var meaningOfLife int
	//  And then setting a value
	meaningOfLife = 42

	// Variable initialization
	//  Here we're initializing the value from the get-go
	var testScore = 98

	// Short Declaration
	//  And here we're skipping the var keyword and using the short delcartion syntax which is a colon and equals
	daysUntilGophercon := 67

	// Let's take a look at what is stored in the variables
	//  In our Hello, Space example we just printed out a string. Now we're adding variables into the mix
	//  We can simply do this with the Println function by adding a comma then the variable we want displayed.
	//  Snazzy.
	fmt.Println("Meaning of life:", meaningOfLife)
	fmt.Println("Test Score:", testScore)
	fmt.Println("Days until Gophercon:", daysUntilGophercon)
}
