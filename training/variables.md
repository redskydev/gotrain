### Variables
There are a few different ways in Go to create variables. 
* We can go the explicit route of creating a variable where we then assign a value later (this is useful at the top of functions to give the user an idea of what may be coming up)

* We can default assign a value at declaration time (variable initialization), or

* We can use short declaration so we don't have to type out the data type. To do this we use the _:=_ operator.

The following example can be found in training/examples/variables_01.go
```go
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
```