# Arrays
I like to think of arrays as storage containers - like those little pill containers that people can store doses in for the week.
So, in this example, we have an entire item (the pill container), sections of that container, and the pills that are placed inside them.

Let's put this into some Computer Science-y speak now (yay).

An array, in Go, is a collection of N-elements of type T. Whaaatt? Let's just think of it concretely for now as a collection of three names.

Let's look at code.

```go
	var names [3]string
	names[0] = "Steven"
	names[1] = "Christopher"
	names[2] = "Jackson"

    fmt.Println(names[0])
	fmt.Println(names)
```

Ok, so we see that we're declaring a variable like we've done before but this time there's what looks like a box with a number 3 in it.

All that first line is doing is declaring an array named 'names' is 3 strings-worth of memeroy long and that contains up to 3 strings.

More on Memory and why it matters later.

We could have done something like: `var names [100]string` and that would have allocated space for 100 names/strings.

The next three lines are how we open the lid of a particular day in the pill box example.
Because we're doing Computer Science things, we start counting with 0. 0 is the new 1. So cool.

So at index 0 (ok, so whenever we access a particular array item we say that it is at index 0,1,N) we are storing the string "Steven"
`names[0] = "Steven"`. We do the same for index 1 and index 2 and put different strings in those.

Finally we're just printing to the screen. If the first print, we're choosing the string at index 0 which will be (drumroll, please) "Steven".
On the final line we're printing out the entire array of strings.

If we didn't know how many items we wanted to store we would use a slice. Slices are first-class citizens in Go.

[Learn about Slices](/training/slices.md)