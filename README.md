# gohr

[![Go Report Card](https://goreportcard.com/badge/github.com/hasit/gohr)](https://goreportcard.com/report/github.com/hasit/gohr)

Unclutter your terminal using this small, but helpful Go package. `gohr` puts horizontal rulers in your terminal so that you can visually seperate parts of text.

`hasit/gohr` comes with:
1. `gohr` - An importable package to use in your own Go programs.
2. `hr` - A CLI tool.

`hasit/gohr` is a port of [LuRst/hr](https://github.com/LuRsT/hr).

[![asciicast of gohr](https://asciinema.org/a/124619.png)](https://asciinema.org/a/124619)

## Installation

```
$ go get -u github.com/hasit/gohr/cmd/hr
```

This will download and install `gohr` package and `hr` executable.

## Usage

`gohr` can be used as a package and as a standalone CLI tool.

### As a package

To use it in your own Go program, import `gohr` and call `gohr.Draw(args ...string)` function.

```go
package main

import "github.com/hasit/gohr"

func main() {
	gohr.DrawHr("-0-")
	gohr.DrawHr("-", "#")
	gohr.DrawHr(".", "\\", "/")
}
```

### As a CLI tool

To use it as a standalone CLI tool, run `hr` from your terminal.

```
$ hr
##########
```

You can provide a symbol of your choice as a command line argument.

```
$ hr 'o'
oooooooooo
```

You can also provide multiple patterns as arguments seperated by spaces.

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

## Contribution

Feel free to ask questions, post issues and open pull requests.
