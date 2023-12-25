# Goque

### A filesystem-backed messaging queue written in Go

There is no purpose to this project, and is an academic excursion to learn how to build cli applications in Go. Although if you are interested, contributions are most welcome.

## Usage

I'll be adding here once I have some actual material, tests. Until then, use `go build -o bin/goque` to create a binary.

Available Commands:

```
add         Add a message to a topic
read        Read messages from a topic
remove      Remove a topic
```

Usage:

```
./bin/goque add -t topic -m msg
./bin/goque read -t topic -n number
./bin/goque remove -t topic
```

Use `./bin/goque add --help` for more details on `add`, and so on.

Currently, to use this, you'll need a linux system, and the following file structure present on your filesystem:

```bash
$HOME/.goque/
└── data/
```

`data` is an empty directory. For each topic, a subdirectory pertaining to that topic is created inside data. Each topic directory contains two files, `index` and `log`. `log` file is where incoming messages are appended, and `index` file takes care of sequential reads.

## Code structure

I have tried to follow [`qrcp`](https://github.com/claudiodangelis/qrcp)'s code template, but in a simplified manner, without subpackages and such. I have been a big fan of `qrcp` over time.

At the end, it uses [`cobra`](https://cobra.dev/) to manage the cli.

To connect with the filesystem, I am using [`os`](https://pkg.go.dev/os) package.

Feel free to point out the flaws, or something that you feel isn't idiomatic to Go.

## Next Steps

1. Make this system concurrency-safe
1. A trash-compactor implementation for maintenance, which removes all the messages that are already read
1. A clean up to provide configured for data directory and easy-install scripts
