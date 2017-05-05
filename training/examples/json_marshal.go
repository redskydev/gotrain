// This program shows how to Marshal (encode) a struct into json.

// Always gotta set a package name. For most of these tutorials we'll use main and for many small applications it's all
//  you'll need!
package main
/* More on Packages at <RSD:PACKAGE_URL>*/

// These are the extras that we need in our program that aren't out-of-the-box items.
// We're using encoding/json so we can, um, encode the json
// We're using fmt so we can print things to the screen
import (
	"encoding/json"
	"fmt"
)
/* More on Imports at <RSD:STRUCT_URL>*/

// Creating a simple Person struct
// Remember, capital letter export. This would be more important if we had this struct in a separate file.
type Person struct {
	First string
	Last string
	Handle string
}
/* More on Structs at <RSD:STRUCT_URL>*/

// The main function. The Big Cheese, the Big Chedda. This is where the magic happens in many of the programs you'll
//  write that aren't libraries. What are libraries? Check out the tutorial at <RSD:OUR_URL>
func main() {
	// Creating a Person struct named 'me' with the details (struct variables).
	me := Person{First:"Steven", Last:"Jackson", Handle: "jacksinn"}
	/* More on Structs at <RSD:STRUCT_URL>*/

	// Calling the json.Marshal function which accepts many variable types (interfaces, technically) and returns a
	//  byte array and an error
	json_struct, err := json.Marshal(me)

	// We always need to handle errors. Let's start by using the function we created earlier!
	HandleError(err)
	/* More on Errors at <RSD:ERROR_URL>*/
	/* More on Imports at <RSD:STRUCT_URL>*/

	// Alright, lets print this bad boy!
	//   We should expect to see our Person (me) printed out.
	fmt.Println(json_struct)

	// Yay...what the...that looked weird. Yeah, that's because the json.Marshal returned a byte array.
	//  Let's jus go ahead and cast this byte array as a string and print it to the screen.
	fmt.Println(string(json_struct))
	// More on Casting at <RSD:CASTING_URL>
}

// Rather than retyping the same code over and over, we can create a function to do that for us and we get to type less!
// This function handles errors (go figure!) and if there is one, it prints to the screen.
//  HandleError accepts a type error (named err here) and does not return anything.
func HandleError(err error) {
	// Let's see if the err is not nil. If it is, we're good. If not, print that error to the screen.
	if err != nil {
		//Let's print the error out with a newline to make it pretty.
		fmt.Println(err)
	}
	//This function doesn't return anything so we don't need to say `return something` anywhere in this function.
}