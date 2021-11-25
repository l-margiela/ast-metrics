# AST Metrics

This code listing is an experiment that came to my mind when I was integrating
some of my code with a metrics service and found out that a big part of the
effort was very mechanical and could be, in a way, automatised.

## TL;DR

The idea is that every code block gets its execution time measured, while
keeping the source code free of tedious metrics calls.

This is very much against Go's priciples and I wouldn't use it anywhere
near production code, although it was fun to write and reason about.


## Tests

Due to the fact it is a highly experimental code in every aspect, it has no
test coverage. Some tests could be written after some research on how to
effectively test AST transformers in Go, which would probably take twice as
much time as writing this code itself, and I don't believe in value of poor
unit tests.

## How to run it?

```bash
make run-mod
```

will generate and execute the code.

To see how it looks like without the "metrics", use:

```bash
make run
```