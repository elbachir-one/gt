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
	Help       bool
	Version    bool
	ShowHidden bool
	Unsort     bool
	Summary    bool
	Dir        string
}

var version = "gt: v0.1.0"

var icons = map[string]string{
	".c":       "\033[34m\033[0m",
	".cpp":     "\033[34m\033[0m",
	".html":    "\033[33m\033[0m",
	".css":     "\033[34m\033[0m",
	".js":      "\033[33m\033[0m",
	".py":      "\033[36m\033[0m",
	".sh":      "\033[32m\033[0m",
	".java":    "\033[31m\033[0m",
	".ino":     "\033[34m\033[0m",
	".rs":      "\033[33m\033[0m",
	".go":      "\033[36m\033[0m",
	".txt":     "\033[37m\033[0m",
	".png":     "\033[35m\033[0m",
	".jpg":     "\033[35m\033[0m",
	".jpeg":    "\033[35m\033[0m",
	".gif":     "\033[35m󰵸\033[0m",
	".mp4":     "\033[31m\033[0m",
	".mov":     "\033[31m\033[0m",
	".mp3":     "\033[31m\033[0m",
	".xfc":     "\033[34m\033[0m",
	".zip":     "\033[31m󰛫\033[0m",
	".gz":      "\033[31m󰛫\033[0m",
	".o":       "\033[33m󰆧\033[0m",
	".obj":     "\033[33m󰆧\033[0m",
	".out":     "\033[32m\033[0m",
	"":         "\033[32m\033[0m",
	".bin":     "\033[32m\033[0m",
	".h":       "\033[35m\033[0m",
	".hpp":     "\033[35m\033[0m",
	"directory": "\033[34;1m\033[0m",
	"other":    "\033[1m\033[0m",
	"symlink":  "\033[36m\033[0m",
}

var dirs, files int

var innerPointers = []string{"├── ", "│   "}
var finalPointers = []string{"└── ", "    "}

func parseArgs() Args {
	var args Args
	flag.BoolVar(&args.Help, "h", false, "show help")
	flag.BoolVar(&args.Help, "help", false, "show help")
	flag.BoolVar(&args.Version, "v", false, "show version")
	flag.BoolVar(&args.Version, "version", false, "show version")
	flag.BoolVar(&args.ShowHidden, "s", false, "show hidden files")
	flag.BoolVar(&args.ShowHidden, "show-hidden", false, "show hidden files")
	flag.BoolVar(&args.Unsort, "u", false, "unsort files")
	flag.BoolVar(&args.Unsort, "unsort", false, "unsort files")
	flag.BoolVar(&args.Summary, "m", false, "show summary")
	flag.BoolVar(&args.Summary, "summary", false, "show summary")

	flag.Parse()

	if len(flag.Args()) > 0 {
		args.Dir = flag.Args()[0]
	} else {
		args.Dir = "."
	}

	return args
}

func walk(directory, prefix string, args Args) error {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return err
	}

	var filesList []os.DirEntry

	for _, entry := range entries {
		if args.ShowHidden || entry.Name()[0] != '.' {
			filesList = append(filesList, entry)
		}
	}

	if !args.Unsort {
		sort.Slice(filesList, func(i, j int) bool {
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

		fmt.Print(prefix + pointers[0])
		icon := icons["other"]

		if entry.IsDir() {
			icon = icons["directory"]
		} else if entry.Type()&os.ModeSymlink != 0 {
			icon = icons["symlink"]
		} else {
			ext := filepath.Ext(entry.Name())
			if ic, found := icons[ext]; found {
				icon = ic
			}
		}

		fmt.Print(icon + " " + entry.Name() + "\n")

		if entry.IsDir() {
			dirs++
			walk(filepath.Join(directory, entry.Name()), prefix+pointers[1], args)
		} else {
			files++
		}
	}

	return nil
}

func main() {
	args := parseArgs()

	if args.Version {
		fmt.Println(version)
		return
	}

	if args.Help {
		fmt.Println(version)
		fmt.Println("Shows a tree of files.")
		fmt.Println("Arguments:")
		flag.PrintDefaults()
		return
	}

	fmt.Println(args.Dir)
	err := walk(args.Dir, "", args)
	if err != nil {
		fmt.Println("Error accessing directory:", err)
	}

	if args.Summary {
		fmt.Printf("\n%d directories, %d files\n", dirs, files)
	}
}
