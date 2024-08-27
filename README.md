# gt

![image-gt](https://i.postimg.cc/Dzp2zHJ7/gtshoot.png)
<br>
<br>
#### `gt`, is a command-line tool written in Go that displays a tree of files and directories with colored icons.

## Features

- Displays a tree of files and directories.
- Supports showing hidden files.
- Option to unsort files.
- Displays summary of files and directories.
- Colored icons based on file types.
- List directories only.
- Order files.
- Print full path.

## Installation

1. Ensure you have [Go](https://go.dev/dl/) installed, and also [Nerd Fonts](https://www.nerdfonts.com)

2. Clone the repository

2-1. Go
```bash
    git clone https://github.com/elbachir-one/gt
    cd gt/
    go build
    sudo cp gt /usr/local/bin/gt
    sudo cp gt.1 /usr/local/share/man/man1/ # To copy the man page.
    gt
```
2-2. Using the Makefile

``` bash
    git clone --depth 1 https://github.com/elbachir-one/gt
    cd gt/
    make
    sudo make install
```

#### Note: gt is in the [AUR](https://aur.archlinux.org/packages/gt), and a [template](https://github.com/elbachir-one/void-templates) is available for Void Linux.

- Arch Linux
```bash
    yay -S gt
```

- Void Linux

```bash
    git clone --depth=1 https://github.com/void-linux/packages
    cd void-packages/
    ./xbps-src binary-bootstrap
    mkdir srcpkgs/gt
    vim srcpkgs/gt/template # Use you own editor
```
Past the content of this
[template](https://github.com/elbachir-one/void-templates) to your local
template and save it.

```bash
    ./xbps-src pkg gt
    sudo xbps-install -R hostdir/binpkgs gt
```

## Usage

```bash
    gt
```
```bash
    gt [OPTIONS] [DIRECTORY]
```

#### Note: you can also pip `gt` to `fzf` or `less`.

```bash
    gt | less -r
```
![gt and less](https://i.postimg.cc/d3tzmKjs/2024-07-21-18-05.png)

`-r` is an option for less to repaint the screen.

```bash
    gt | fzf --ansi
```
![gt and fzf](https://i.postimg.cc/C5P9c6cj/2024-07-21-18-06.png)

Options:

| Option  | Description                                                                                                                    | example                                          |
|---------|--------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------|
| -h --help | Show help                                                                                                                      | ![-h](https://i.postimg.cc/647QG2YT/h.png)       |
| -v --version | Show version                                                                                                                   | ![-v](https://i.postimg.cc/xk0T3Ftw/v.png)       |
| -a --all | Show hidden files                                                                                                              | ![-s](https://i.postimg.cc/4YR3m3tN/s.png)       |
| -u --unsorted | Files                                                                                                                          | ![-u](https://i.postimg.cc/Bj5nDGc5/u.png)       |
| -m --summary | Summary                                                                                                                        | ![-m](https://i.postimg.cc/PNcJny5D/m.png)       |
| -d --directories | List directories only                                                                                                          | ![-d](https://i.postimg.cc/YGD99fNJ/d.png)       |
| -o  --order-by-extension | Order files based on extension                                                                                                 | ![-o](https://i.postimg.cc/VdC6ftJV/o.png)       |
| -f --full-path | Print full path prefix to each file                                                                                            | ![-f](https://i.postimg.cc/mhhknJBR/f.png)       |
| --depth | Depth to which the tree should be displaye the default is -1 `gt --depth 1` any number greater than or equal to 1 can be used. | ![--depth](https://i.postimg.cc/yg8xsrRm/dd.png) |

## Contributing

Contributions, issues, and feature requests are always welcome! Thank you.
