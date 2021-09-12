# Monkey

Monkey is the programming language from [Writing An Interpreter In Go](https://interpreterbook.com) and [Writing A Compiler In Go](https://compilerbook.com) by [Thorsten Ball](https://thorstenball.com).

This repository contains my implementation of the monkey programming language written while following the books.

My implementation of a compiler and a virtual machine for monkey from [Writing A Compiler In Go](https://compilerbook.com) can be found on the `main` branch.

My implementation of a tree walking interpreter for monkey from [Writing An Interpreter In Go](https://interpreterbook.com) can be found on the [writing-an-interpreter-go](https://github.com/ChrisWilding/monkey/tree/writing-an-interpreter-go) branch.

My implementation of [A Macro System For Monkey](https://interpreterbook.com/lost) can be found on the [a-macro-system-for-monkey](https://github.com/ChrisWilding/monkey/tree/a-macro-system-for-monkey) branch.

## Prerequisites

1. [Go 1.17](https://golang.org/doc/install)

## Build

```sh
$ go build
```

## Run

```sh
$ ./monkey
Hello USERNAME! This is the Monkey programming language!
Feel free to type in commands
>> let unless = macro(condition, consequence, alternative) { quote(if (!(unquote(condition))) { unquote(consequence); } else { unquote(alternative); }); };
>> unless(10 > 5, puts("not greater"), puts("greater"));
greater
```
