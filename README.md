# hr.go
A port of [LuRst/hr](https://github.com/LuRsT/hr) in Go.

Unclutter your terminal using this small, but helpful program. Put the `<hr />` tag in your terminal and visually seperate parts of text.

![hr.go](assests/gohr.png)

## Setup
Clone the repository <br />
`git clone https://github.com/hasit/gohr`

Install the package (you must have GOPATH set) <br />
`go install`

## Usage
Assuming that you have already added GOPATH/bin/ to PATH, simply run the command <br />
`hr` <br />
`##########` <br />
With no arguments, the default symbol '#' will fill one row in your terminal

You can provide your own symbol as a command line argument <br />
`hr 'o'` <br />
`oooooooooo` <br />
This is fill one row in your terminal with symbol 'o'

You can also provide multiple patterns as arguments seperated by spaces <br />
`hr '-' '#' '-'` <br />
`----------` <br />
`##########` <br />
`----------` <br />
<br />
`hr '-o-' '#' '-o-'` <br />
`-o--o--o--` <br />
`##########` <br />
`-o--o--o--` <br />
<br />
## TODO
- [ ] Make it `go get`able
- [ ] Make it available as a library
- [ ] Add proper doc and examples
