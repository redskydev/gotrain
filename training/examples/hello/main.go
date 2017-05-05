// This program shows how to Marshal (encode) a struct into json.

// Always gotta set a package name. For most of these tutorials we'll use main and for many small applications it's all
//  you'll need!
package main
/* More on Packages at <RSD:PACKAGE_URL>*/


// These are the extras that we need in our program that aren't out-of-the-box items.
// We're using encoding/json so we can, um, encode the json
// We're using fmt so we can print things to the screen
import (
	"fmt"
)
/* More on Imports at <RSD:STRUCT_URL>*/


// The main function. The Big Cheese, the Big Chedda. This is where the magic happens in many of the programs you'll
//  write that aren't libraries. What are libraries? Check out the tutorial at <RSD:OUR_URL>
func main(){
	// This function which we imported earlier from the "fmt" package allows us to print a string with an automatic
	//  newline at the end so we don't have to worry about it ourselves
	fmt.Println("In space no one can hear hellooooooo!")
}
