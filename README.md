# Go Training Overview

[Training](/training/index.md)
# Level

## Starting Out
### Get Go
Go to 
### Editors
* VimGo
* Gogland
* Whatever

### WTF is a workspace?
Yeah so if you're coming from a web programming background this idea of workspaces, GOROOT, etc may leave in a bit of a WTFM8 moment. No worries, we'll get you there.

### Hello, Space
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

### Variables
There are a couple different ways in Go to create variables. 
We can go the explicit route of creating a variable where we then assign a value later (this is useful at the top of functions to give the user an idea of what may be coming up)
We can default assign a value at declaration time (variable initialization), or
We can use short declaration so we don't have to type out the data type. To do this we use the _:=_ operator.

```go
// Basic Variable Declaration
var meaningOfLife int
meaningOfLife = 42

// Variable initialization
var testScore = 98

// Short Declaration
daysUntilGophercon := 67
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
