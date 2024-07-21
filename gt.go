package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Args struct {
	Help         bool
	Version      bool
	ShowHidden   bool
	Unsort       bool
	Summary      bool
	DirsOnly     bool
	FullPath     bool
	OrderByExt   bool
	Dir          string
	Depth        int
}

var version = "gt: v0.2.1"

var icons = map[string]string{
	".c":          "\033[34m \033[0m",
	".cpp":        "\033[34m \033[0m",
	".cs":         "\033[34m󰌛 \033[0m",
	".css":        "\033[34m \033[0m",
	".js":         "\033[33m󰌞 \033[0m",
	".json":       "\033[93m \033[0m",
	".php":        "\033[38;5;39m \033[0m",
	".sqlite":     "\033[38;5;22m \033[0m",
	".sh":         "\033[32m \033[0m",
	".iso":        "\033[37m \033[0m",
	".java":       "\033[31m \033[0m",
	".ino":        "\033[34m \033[0m",
	".rs":         "\033[90m \033[0m",
	".go":         "\033[36m \033[0m",
	".txt":        "\033[37m \033[0m",
	".png":        "\033[35m \033[0m",
	".jpg":        "\033[35m󰈥 \033[0m",
	".jpeg":       "\033[35m \033[0m",
	".webp":       "\033[35m \033[0m",
	".mov":        "\033[31m󰨜 \033[0m",
	".webm":       "\033[35m󰑈 \033[0m",
	".m4a":        "\033[94m \033[0m",
	".flac":       "\033[31m󱀞 \033[0m",
	".wav":        "\033[31m󰗅 \033[0m",
	".xcf":        "\033[34m \033[0m",
	".tar":        "\033[31m󰞹 \033[0m",
	".zip":        "\033[31m󰛫 \033[0m",
	".gz":         "\033[31m󰛫 \033[0m",
	".o":          "\033[33m󰆧 \033[0m",
	".obj":        "\033[33m󰆧 \033[0m",
	".out":        "\033[32m \033[0m",
	"":            "\033[32m \033[0m",
	".bin":        "\033[32m \033[0m",
	".h":          "\033[35m \033[0m",
	".hpp":        "\033[35m \033[0m",
	".deb":        "\033[90m \033[0m",
	".yml":        "\033[31m \033[0m",
	".yaml":       "\033[31m \033[0m",
	".html":       "\033[38;5;208m \033[0m",
	".xml":        "\033[38;5;208m󰗀 \033[0m",
	".py":         "\033[38;5;172m \033[0m",
	".mp4":        "\033[38;5;198m󰕧 \033[0m",
	".mp3":        "\033[38;5;39m \033[0m",
	".gif":        "\033[38;5;198m󰵸 \033[0m",
	".toml":       "\033[38;5;208m \033[0m",
	".zig":        "\033[38;5;208m \033[0m",
	".xbps":       "\033[38;5;22m \033[0m",
	".svg":        "\033[35m󰜡 \033[0m",
	".conf":       "\033[37m \033[0m",
	".gitignore":  "\033[38;5;208m󰊢 \033[0m",
	".md":         "\033[34m \033[0m",
	".rb":         "\033[31m󰴭 \033[0m",
	".pdf":        "\033[38;5;196m󰈦 \033[0m",
	".el":         "\033[38;5;125m \033[0m",
	".org":        "\033[38;5;125m \033[0m",
	".vim":        "\033[32m \033[0m",
	".epub":       "\033[94m󰂺 \033[0m",
	".ttf":        "\033[97m \033[0m",
	".otf":        "\033[97m󰛖 \033[0m",
	".db":         "\033[97m󰆼 \033[0m",
	"directory":   "\033[34;1m \033[0m",
	"other":       "\033[1m \033[0m",
	"symlink":     "\033[36m \033[0m",
}

var dirs, files int

var innerPointers = []string{"├── ", "│   "}
var finalPointers = []string{"└── ", "    "}

func parseArgs() Args {
	var args Args
	flag.BoolVar(&args.Help, "h", false, "show help")
	flag.BoolVar(&args.Version, "v", false, "show version")
	flag.BoolVar(&args.ShowHidden, "s", false, "show hidden files")
	flag.BoolVar(&args.Unsort, "u", false, "unsort files")
	flag.BoolVar(&args.Summary, "m", false, "show summary")
	flag.BoolVar(&args.DirsOnly, "d", false, "list directories only")
	flag.BoolVar(&args.FullPath, "f", false, "print full path prefix to each file")
	flag.BoolVar(&args.OrderByExt, "o", false, "order files based on extension")
	flag.IntVar(&args.Depth, "depth", -1, "depth to which the tree should be displayed")

	flag.Parse()

	if len(flag.Args()) > 0 {
		args.Dir = flag.Args()[0]
	} else {
		args.Dir = "."
	}

	return args
}

func walk(directory, prefix string, args Args, currentDepth int) error {
	if args.Depth != -1 && currentDepth > args.Depth {
		return nil
	}

	entries, err := os.ReadDir(directory)
	if err != nil {
		return err
	}

	var filesList []os.DirEntry

	for _, entry := range entries {
		if args.ShowHidden || entry.Name()[0] != '.' {
			if args.DirsOnly && !entry.IsDir() {
				continue
			}
			filesList = append(filesList, entry)
		}
	}

	if !args.Unsort {
		sort.Slice(filesList, func(i, j int) bool {
			if args.OrderByExt {
				extI := filepath.Ext(filesList[i].Name())
				extJ := filepath.Ext(filesList[j].Name())
				if extI == extJ {
					return strings.ToLower(filesList[i].Name()) < strings.ToLower(filesList[j].Name())
				}
				return extI < extJ
			}
			return strings.ToLower(filesList[i].Name()) < strings.ToLower(filesList[j].Name())
		})
	}

	for i, entry := range filesList {
		var pointers []string
		if i == len(filesList)-1 {
			pointers = finalPointers
		} else {
			pointers = innerPointers
		}

		fullPath := entry.Name()
		if args.FullPath {
			fullPath = filepath.Join(directory, entry.Name())
		}

		fmt.Print(prefix + pointers[0])
		icon := icons["other"]

		if entry.IsDir() {
			fmt.Print(icons["directory"])
			dirs++
		} else if entry.Type()&os.ModeSymlink != 0 {
			fmt.Print(icons["symlink"])
		} else {
			ext := filepath.Ext(entry.Name())
			if val, ok := icons[ext]; ok {
				fmt.Print(val)
			} else {
				fmt.Print(icon)
			}
			files++
		}
		fmt.Println(fullPath)

		if entry.IsDir() {
			walk(filepath.Join(directory, entry.Name()), prefix+pointers[1], args, currentDepth+1)
		}
	}

	return nil
}

func main() {
	args := parseArgs()

	if args.Help {
		flag.Usage()
		return
	}

	if args.Version {
		fmt.Println(version)
		return
	}

	err := walk(args.Dir, "", args, 1)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if args.Summary {
		fmt.Printf("\n%d directories, %d files\n", dirs, files)
	}
}
