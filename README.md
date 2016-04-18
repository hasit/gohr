# gohr

A port of [LuRst/hr](https://github.com/LuRsT/hr) in Go.

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hasit/gohr/master/LICENSE.txt)

## About

Unclutter your terminal using this small, but helpful program. Put the `<hr />` tag in your terminal and visually seperate parts of text.

![hr.go](assests/gohr.png)

## Installing

To start using gohr, run `go get`:

```terminal
$ go get github.com/hasit/gohr/hr
```

This will download and install `gohr` library and `hr` executable.

## Usage

`gohr` can be used as a library as well as a standalone application.

To use it as a library, import `gohr` in your program and call `DrawHr` function:

```go
package main

import (
	"os"

	"github.com/hasit/gohr"
)

func getArgs() []string {
	return os.Args[1:]
}

func main() {
	args := getArgs()
	gohr.DrawHr(args)
}
```

You can find this example in `cmd/hr/main.go` file.

To use it as a standalone application, run `hr` from your terminal:

```
$ hr
##########
```

You can provide a symbol of your choice of symbol as a command line argument

```
$ hr 'o'
oooooooooo
```

You can also provide multiple patterns as arguments seperated by spaces

```
$ hr '-' '#' '-'
----------
##########
----------
```

```
$ hr '-o-' '#' '-o-'
-o--o--o--
##########
-o--o--o--
```

## TODO

- [x] Make it `go get`able
- [x] Make it available as a library
- [x] Write `hr.go` using that library
- [ ] Add proper doc
- [x] Add examples
