# hr.go
A port of [LuRst/hr](https://github.com/LuRsT/hr) in Go.

Unclutter your terminal using this small, but helpful program. Put the `<hr />` tag in your terminal and visually seperate parts of text.

![hr.go](assests/gohr.png)

## Setup
Clone the repository
`git clone https://github.com/hasit/gohr`

Install the package (you must have GOPATH set)
`go install`

## Usage
Assuming that you have already added GOPATH/bin/ to PATH, simply run the command
`hr`
`##########`
With no arguments, the default symbol '#' will fill one row in your terminal

You can provide your own symbol as a command line argument
`hr 'o'` <hr />
`oooooooooo` <hr />
This is fill one row in your terminal with symbol '*'

You can also provide multiple patterns as arguments seperated by spaces
`hr '-' '#' '-'` <hr />
`----------` <hr />
`##########` <hr />
`----------` <hr />
<hr />
`hr '-o-' '#' '-o-'` <hr />
`-o--o--o--` <hr />
`##########` <hr />
`-o--o--o--` <hr />
<hr />
## TODO
- [ ] Make it `go get`able
- [ ] Make it available as a library
- [ ] Add proper doc and examples
