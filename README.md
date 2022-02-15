# blaise-go-fizzbuzz

A GoLang coding kata based around the fizzbuzz problem

## Getting started

- Clone this repo!
- Make yourself a branch `git checkout -b my-branch`
- `go get ./...`

## The problem

We need a simple API that will respond appropriately when given a number.

The rules:

- If the number is divisible by `3` display `Fizz`
- If the number is divisible by `5` display `Buzz`
- If the number is divisible by both `3` and `5` display `FizzBuzz`
- If the number is not divisible by `3` or `5` display the number, e.g `2` should display `2`

To ease you into the problem the codebase has already been setup so that it compiles and gives you a basic project
structure split into an implementation function and a webserver/ controller.

The repo has some example (failing) tests to get you started, be *warned* they do not cover edge cases and you should
write some extra tests.

## Pairing? Make it a game!

To get the most out of this kata, I strongly recommend pairing, and if you are pairing you might as well make it fun.

We are going to do some "Evil TDD".

The rules:

- We will have 2 roles, the coder and the tester
- The coder - your job is to write the minimum possible implemenation take a look at the
  evil tdd [example](example/eviltdd)
- The tester - your job is to write tests that make the coder write a good implementation
  (the example tests definitely don't do this!)
- Rotate roles! Set a timer of 20-30 minutes (less if your feeling crazy) so both of you get a change to write the tests
  and be an agent of chaos

## Run the tests

You should run the tests whenever you change any code, lucky for you go runs tests nice and quickly

```sh
go test ./...
```

## Too easy?

Extend the API. Think of something new you want to do, create a new controller and get it working, obviously with tests.

**Note**: Please update this section with any ideas that made for a good extension.
