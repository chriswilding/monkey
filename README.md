# Monkey

Monkey is the programming language from [Writing An Interpreter In Go by Thorsten Ball](https://interpreterbook.com).

This repository contains my implementation of the monkey programming language written while following the book.

My implementation of [A Macro System For Monkey](https://interpreterbook.com/lost) can be found on the [a-macro-system-for-monkey](https://github.com/ChrisWilding/monkey/tree/a-macro-system-for-monkey) branch.

## Prerequisites

1. [Go 1.16](https://golang.org/doc/install)

## Build

```sh
$ go build
```

## Run

```sh
$ ./monkey
Hello USERNAME! This is the Monkey programming language!
Feel free to type in commands
>> let people = [{"name": "Erik", "age": 24}, {"name": "Astrid", "age": 28}];
>> people[0]["name"];
Erik
>> people[1]["age"];
28
>> people[1]["age"] + people[0]["age"];
52
>> let getName = fn(person) { person["name"]; };
>> getName(people[0]);
Erik
>> getName(people[1]);
Astrid
```
