# Monkey Programming Language Interpreter
This repo was made in process of reading this amazing book "[Writing An Interpreter In Go](https://interpreterbook.com)"

## Running the Interpreter
To run the interpreter you need to run
```
go run ./cmd/main.go
```

## Features
Even tho this interpreter isn't 200k loc it has quite a lot of interesting features.

```
// variable declaration

let a = 5;
let b = "foo";
let c = false;

// you can also assign functions to variables
let d = fn() { return 69; }
d(); // returns 69

// if statement

if(5 > 6) {
    puts("foo");
} else {
    puts("bar");
}

// simple mafs also works

let e = 5 * 5 / 25 - 1 + 4;

// you can pass functions as parameters and do all sorts of crazy things

let do_smth = fn(other_fn, a, b) {
    other_fn(a,b);
}

do_smth(
    fn(b, c) {
        return b + c;
    },
    1,
    2
) // returns 3

```
