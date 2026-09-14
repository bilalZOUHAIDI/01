# Exercise 1: Debug Log & AI Sharpening

## Step 1: Manual Bug Identification & Fix (No AI)

### My hypothesis before fixing

Before checking the code too much, I thought the problem was that the parser was only checking normal spaces `' '`.
So it was probably not detecting other spaces like tabs `\t` and new lines `\n`, even if they are also whitespace in the exercise.

### What was the prblm

In the original `ex01_buggy.go`, the code was only checking:

```go
if c != ' '
```

So when the input had a tab or a new line, the program was treating them like normal chars. They was then added inside the word instead of separating the words.

So for example, something like `hello\tworld` was not correctly split into 2 words.

=================================================

## Step 2: Going Deeper with AI

### Prompt I sent to AI

> Here is my fixed function:
>
> ```go
> func isWhiteSpace(c rune) bool {
>     return c == ' ' || c == '\t' || c == '\n'
> }
>
> func SplitWhiteSpaces(s string) []string {
>     var res []string
>     word := ""
>     inWord := false
>
>     for _, c := range s {
>         if !isWhiteSpace(c) {
>             word += string(c)
>             inWord = true
>         } else {
>             if inWord {
>                 res = append(res, word)
>                 word = ""
>                 inWord = false
>             }
>         }
>     }
>
>     if inWord {
>         res = append(res, word)
>     }
>
>     return res
> }
> ```
> Is there a cleaner way to do it?

### What AI answered


First, `\r` can be a prblm. For exmpl with windows line endings `\r\n`, the `\r` would stay attached to the word bec my func doesnt chck for it.

It also mentioned other Unicode whitespace characters, like non-breaking spaces, `\f` and `\v`. My func doesnt detect those either.

The AI also said that my `word += string(c)` isn't the best way to build the string bec it can create a lot of allocations when the word is long.

It suggested using indexes/slicing instead, or using `strings.Fields()`.

==================================================

## Evaluation & Reflection

### Did AI point out something useful that I didn't think about?

Yeah, it did. The `\r` case was something I didn't consider at first. For exmpl with `\r\n`, my func would keep the `\r` at the end of the word.

The AI also explained that doing `word += string(c)` many times isnt very efficient. Using string indexes and slicing would be better.

### Checking the AI suggs

I checked the suggs to see if they actually work for this exercise.

The point about `\r` is correct and can be useful, especially for windows line endings.

The point about string concatenation is also correct. It could be improved by using indexes instead of creating a new string for each character.

But for `strings.Fields()`, I can't use it here. In the Piscine, the `strings` package is not allowed for this exercise. So even if it's a more idiomatic Go solution, using it would make the solution invalid for the Piscine tests/audit.

So I didn't blindly use the AI answer. I kept my own implementation and only used the suggestions that actually fit the exercise.
