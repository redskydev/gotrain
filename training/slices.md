# Slices
Slices are like if we had that pill container from the Array tutorial but sometimes we only needed a week's worth because we're at home, whatever, and sometimes we need a month's worth because we're living it up in Cabo.
Well I suppose we could pack four or five of the pill thingies but it'd be great if we could have one that holds 30. Or 28. or 31.

With slices, we can do that. We just say we want a slice (like in Brooklyn). 
For the pill container example its like if we could just snap on additional sections if we needed them, like a toy-brick (can I say Lego?) chain.

Enough speak, more example.

```go
	var names []string
	names = append(names, "Steven")
	names = append(names, "Christopher")
	names = append(names, "Jackson")

	fmt.Println(names[0])
	fmt.Println(names)
```

Alrighty so it looks just like an array at first and we jsut didn't have to say how long it is. Sweet.
But next we see if we want to add an item onto the slice we have to use `append()` because we don't know how long the slice is (we'll get to that later).
No biggie, I like `append()` it's pretty explicit, IMHO, as to what's going on.

So, here `names = append(names, "Steven")`, we're saying append "Steven" to the slice named 'names'. 
So Go will do that but it doesn't mutate (nice word, we'll talk about it) the slice so we just set the same slice, 'names', to be what `append()` returns which is the original slice plus this new fella.

We add a couple more names but then, oh boy, we can still access the elements of the slice the same way as an array (because a slice is backed by an array).

Soooooooo, what if we have a ton of elements in our slice (or array) wouldn't it take foreever to print each one out individually?

Yes, if you dont embrace [Loops](/training/loops.md).