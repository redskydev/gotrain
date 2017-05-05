# Comments
A comment is a human readable, machine/language-ignored bit of text in a source file that gives ourselves and other developers a heads up of what is going on. 
Some people are in the school of thought that code should be self-documenting and that comments are fluff and a sign of _code smell_ but I believe in documenting/commenting just about everything because it doesn't hurt anything and GET OFF MY LAWN.

Comments start with two forward-slashes _//_ and my convention is that I typically add a space after the start of the comment because I think it looks prettier.
 Anything after the _//_ will be ignored by the compiler and if you're using an IDE you'll typically see it as greyed-out text. If you put any code after this on the same line, it will be ignored.
 You'll see lots of lazy clean-up of programs _commenting out_ old code just in case they need it again. Don't do this (do as I say not as I do). You'll do this.
 
```go
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

fmt.Prinln("Something I want printed out") //Comments can also go at the end of a bit of code on the same line.
```

 