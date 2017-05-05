// This program shows how to use comments.
// If you did the main.go (Hello World) you'll have already seen all this plain english.
//   Comments are intended to give clarity to humans who are reading the program so it's a lot less daunting to learn
//   about what is going on.

// Comments can start with two forward-slashes and anything else after those slashes on the line will not be interpreted
//   as code. e.g. fmt.Println("Cannot print from comments") will not be output

/* Comments can also start with a forward-slash and an asterisk. These types of comments must end with an asterisk then
     the forward slash. This allows us to write
     multi-line
     comments
     without needing to do the double forward slash. Anything between the beginning na end is considered a comment.

   These types of comments are generally for longer documentation.
*/

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
