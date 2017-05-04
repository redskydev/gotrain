# Go Training Overview

[Training](/training/index.md)
# Level

## Starting Out
### Get Go
Go to [https://golang.org/](https://golang.org/) and follow the directions to set it up on your system. 
If the documentation isn't clear to you (I still think it relies on a bit of insider know-how) hit us up on Github and we'll look into making our own documentation.

### Editors
* VimGo
* Gogland
* Whatever

### WTF is a workspace?
Yeah so if you're coming from a web programming background this idea of workspaces, GOROOT, etc may leave in a bit of a WTFM8 moment. 
No worries, we'll get you there.

### Hello, Space
If you followed the directions after downloading and installing Go, you may have already done a Hello, World program. 

Whatever. This one is way better.

Create a file called main.go (say it out loud, it sounds like mango mmmmm), put this deliciousness and then do _go run main.go_
```go
package main

import "fmt"

func main(){
	fmt.Println("In space no one can hear hellooooooo!")
}
```

```
# In the directory where main.go is, run this
go run main.go
# You should see output like this:
In space no one can hear hellooooooo!
```
Sooo what's going on here? We'll go line-by-line and briefly cover the topics and link to further details about the topics.
* _package main_ tells Go that this is the main package. Most programs will need this. For now, assume you need this but later on we'll discuss when to use it or now in the [Packages](#) Section.
* _import "fmt"_ says that we're going to use something from the _fmt_ package. I pronounce it _fumt_ but think of it as _format_ because it is.
* _func main(){_ ok so this is the main function and it's a special function. For now we're going to write everything in the main function but if you want to learn more, jump ahead to the [Functions](#) Section.
* _fmt.Println("In space no one can hear hellooooooo!")_ This is the goods of the program, the meat (or veggie alternative). 
We're saying, hey, we're going to use _Println_ from the _fmt_ package and print whatever we have in quotes. In our case it's the _string_ "In space no one can hear hellooooooo!"
* _}_ we're just closing out the main() function here. All functions need to have a beginning and ending curly brace.

Pretty much from now on we're going to document our code using comments and a lot of the tutorial will come from reading the comments and then looking at the code.
The reasoning for this is two-fold: 1) because code should be documented and you should get used to reading comments (and writing them, Phu >:( ) and 2) because I don't feel like writing separate paragraphs and it's my party.

### Comments
A comment is a human readable, machine/language-ignored bit of text in a source file that gives ourselves and other developers a heads up of what is going on. 
Some people are in the school of thought that code should be self-documenting and that comments are fluff and a sign of _code smell_ but I believe in documenting/commenting just about everything because it doesn't hurt anything and GET OFF MY LAWN.

Comments start with two forward-slashes _//_ and my convention is that I typically add a space after the start of the comment because I think it looks prettier.
 Anything after the _//_ will be ignored by the compiler and if you're using an IDE you'll typically see it as greyed-out text. If you put any code after this on the same line, it will be ignored.
 You'll see lots of lazy clean-up of programs _commenting out_ old code just in case they need it again. Don't do this (do as I say not as I do). You'll do this.
 
### Variables
There are a couple different ways in Go to create variables. 
We can go the explicit route of creating a variable where we then assign a value later (this is useful at the top of functions to give the user an idea of what may be coming up)
We can default assign a value at declaration time (variable initialization), or
We can use short declaration so we don't have to type out the data type. To do this we use the _:=_ operator.

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
### Data Types
#### Integers
An integer is a counting number - like 1,2,3,100, and even 0 and negative numbers -1, -42 etc.

Integers are declared without quotes
#### Floating Point
#### Strings

### Arrays

### Slices

### Functions

### Structs
#### Creation
#### Methods

### Gofmt

### Go Get

## Doing More
### Go Routines
### Channels

## I Am A Code God
### Packaging
### Kits

# Topics
## Command Line Apps

## Web Development

## Microservices

## Tooling

# Tools
## Cobra
## Buffalo

# Philosophies
## Packaging
## Stack vs. Heap

# Examples
## Basic App
## Cobra App
## Buffalo App
## Daemon
