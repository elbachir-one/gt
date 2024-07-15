# gt

`gt`, short for `go tree`, is a command-line tool written in Go that
displays a tree of files and directories with colored icons.

![gt shot](gt.png)

## Features

- Displays a tree of files and directories.
- Supports showing hidden files.
- Option to unsort files.
- Displays summary of files and directories.
- Colored icons based on file types.

## Installation

1. Ensure you have Go installed. You can download it from
   [golang.org](https://golang.org/).

2. Clone the repository:

```bash
   git clone https://github.com/yourusername/gt.git
   cd gt
   go build gt.go
   ./gt
```

3. Install

```bash
    sudo cp gt /usr/local/bin/gt
```

## Usage

```bash
    gt [OPTIONS] [DIRECTORY]
```

Options:

* -h --help          show help
* -v --version       show version
* -s --show-hidden   show hidden files
* -u --unsort unsort files
* -m --summary show  summary
