# gt

`gt`, is a command-line tool written in Go that displays a tree of files
and directories with colored icons.

![gt shot](gt.png)

## Features

- Displays a tree of files and directories.
- Supports showing hidden files.
- Option to unsort files.
- Displays summary of files and directories.
- Colored icons based on file types.
- List directories.
- Order files.
- Print full path.

## Installation

1. Ensure you have [Go](https://go.dev/dl/) installed, and also [Nerd Fonts](https://www.nerdfonts.com)

2. Clone the repository

```bash
   git clone https://github.com/elbachir-one/gt
   cd gt
   go build gt.go
   ./gt
```

3. Install

```bash
    sudo cp gt /usr/local/bin/gt
```


#### Note: gt is in the [AUR](https://aur.archlinux.org/packages/gt), and a [template](https://github.com/elbachir-one/void-templates) is available for Void Linux.

## Usage

```bash
    gt [OPTIONS] [DIRECTORY]
```

Options:

| Option | Description                             |
|--------|-----------------------------------------|
| -h     | Show help                               |
| -v     | Show version                            |
| -s     | Show hidden files                       |
| -u     | Files                                   |
| -m     | Summary                                 |
| -d     | List directories only                   |
| -o     | Order files based on extension          |
| -f     | Print full path prefix to each file     |

## Contributing

Contributions, issues, and feature requests are always welcome! Thank you.
