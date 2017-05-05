# Hello, Space
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
