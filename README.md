# Goque

### A filesystem-backed messaging queue written in Go

There is no purpose to this project, and is an academic excursion to learn how to build cli applications in Go. Although if you are interested, contributions are most welcome.

The name comes from Go and Queue, and is a pronounced like [Goku](https://en.wikipedia.org/wiki/Goku). 
<!-- I am not a weeb though, and this is a poor pun at best. -->

## Usage

I'll be adding here once I have some actual material, tests. Until then, use `go build -o bin/goque` to create a binary, and `.bin/goque --help` should help you out.

## Code structure

I have tried to follow [`qrcp`](https://github.com/claudiodangelis/qrcp)'s code template, but in a simplified manner, without subpackages and such. I have been a big fan of `qrcp` over time.

At the end, it uses [`cobra`](https://cobra.dev/) to manage the cli.

To connect with the filesystem, I am planning to use [`io/fs`](https://pkg.go.dev/io/fs), but the final decision remains to be made.

Feel free to point out the flaws, or something that you feel isn't idiomatic to Go.

## Contributions

I have no idea what to write here.