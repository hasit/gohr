# gohr
A port of [LuRst/hr](https://github.com/LuRsT/hr) in Go.

## About
Unclutter your terminal using this small, but helpful program. Put the `<hr />` tag in your terminal and visually seperate parts of text.

![hr.go](assests/gohr.png)

## Setup
	go get github.com/hasit/gohr/hr

## Usage
Assuming that you have already added `$GOPATH/bin/` to `$PATH`, simply run the command `hr`. With no arguments, the default symbol '#' will fill one row in your terminal.

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
- [ ] Make it available as a library
- [ ] Write `hr.go` using that library 
- [ ] Add proper doc and examples
