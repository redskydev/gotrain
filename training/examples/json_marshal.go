// This program shows how to Marshal (encode) a struct into json.

// Always gotta set a package name. For most of these tutorials we'll use main and for many small applications it's all
//  you'll need!
package main

// These are the extras that we need in our program that aren't out-of-the-box items.
// We're using encoding/json so we can, um, encode the json
// We're using fmt so we can print things to the screen
import (
	"encoding/json"
	"fmt"
)

// Creating a simple Person struct
// Remember, capital letter export. This would be more important if we had this struct in a separate file.
type Person struct {
	First string
	Last string
	Handle string
}

func main() {
	// Creating a Person struct named 'me' with the details (struct variables)
	me := Person{First:"Steven", Last:"Jackson", Handle: "jacksinn"}

	// Calling the json.Marshal function which accepts many variable types (interfaces, technically) and returns a
	//  byte array and an error
	json_struct, err := json.Marshal(me)

	// We always need to handle errors. To not handle them is blasphemy.
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Alright, lets print this bad boy! We should expect to our Person (me) printed out.
	fmt.Println(json_struct)

	// Yay...what the...that looked weird. Yeah, that's because the json.Marshal returned a byte array.
	// Let's jus go ahead and cast this byte array as a string and print it to the screen.
	fmt.Println(string(json_struct))
}