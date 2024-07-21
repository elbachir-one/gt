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

- Arch Linux
```bash
yay -Sy gt
```

- Void Linux

```bash
git clone --depth=1 https://github.com/void-linux/packages
cd void-packages/
./xbps-install binary-bootstrap
mkdir srcpkgs/gt
vim srcpkgs/gt/template
```
Past the content of this
[template](https://github.com/elbachir-one/void-templates) to your local
template and save it.

```bash
./xbps-install pkg gt
sudo xbps-install -R hostdir/binpkgs gt
```

## Usage

```bash
    gt [OPTIONS] [DIRECTORY]
```

#### Note: you can also pip `gt` to `fzf` or `less`.

```bash
gt | less -r
```
`-r` is an option for less to repaint the screen.

```bash
gt | fzf --ansi
```

Options:

| Option  | Description                                                                                                                                      |
|---------|--------------------------------------------------------------------------------------------------------------------------------------------------|
| -h      | Show help                                                                                                                                        |
| -v      | Show version                                                                                                                                     |
| -s      | Show hidden files                                                                                                                                |
| -u      | Files                                                                                                                                            |
| -m      | Summary                                                                                                                                          |
| -d      | List directories only                                                                                                                            |
| -o      | Order files based on extension                                                                                                                   |
| -f      | Print full path prefix to each file                                                                                                              |
| --depth | Depth to which the tree should be displaye the default is -1 `gt --depth 1` any number greater than or equal to 1 can be used.                   |

## Contributing

Contributions, issues, and feature requests are always welcome! Thank you.
