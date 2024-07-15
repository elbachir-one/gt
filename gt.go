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

var version = "ctree: v0.1.0"

var icons = map[string]string{
	".c":       "\033[34m\033[0m", // Blue
	".cpp":     "\033[34m\033[0m", // Blue
	".html":    "\033[33m\033[0m", // Yellow
	".css":     "\033[34m\033[0m", // Blue
	".js":      "\033[33m\033[0m", // Yellow
	".py":      "\033[36m\033[0m", // Cyan
	".sh":      "\033[32m\033[0m", // Green
	".java":    "\033[31m\033[0m", // Red
	".ino":     "\033[34m\033[0m", // Blue
	".rs":      "\033[33m\033[0m", // Yellow
	".go":      "\033[36m\033[0m", // Cyan
	".txt":     "\033[37m\033[0m", // White
	".png":     "\033[35m\033[0m", // Magenta
	".jpg":     "\033[35m\033[0m", // Magenta
	".jpeg":    "\033[35m\033[0m", // Magenta
	".gif":     "\033[35m󰵸\033[0m", // Magenta
	".mp4":     "\033[31m\033[0m", // Red
	".mov":     "\033[31m\033[0m", // Red
	".mp3":     "\033[31m\033[0m", // Red
	".xfc":     "\033[34m\033[0m", // Blue
	".zip":     "\033[31m󰛫\033[0m", // Red
	".gz":      "\033[31m󰛫\033[0m", // Red
	".o":       "\033[33m󰆧\033[0m", // Yellow
	".obj":     "\033[33m󰆧\033[0m", // Yellow
	".out":     "\033[32m\033[0m", // Green
	"":         "\033[32m\033[0m", // Green
	".bin":     "\033[32m\033[0m", // Green
	".h":       "\033[35m\033[0m", // Magenta
	".hpp":     "\033[35m\033[0m", // Magenta
	"directory": "\033[34;1m\033[0m", // Blue Bold
	"other":    "\033[1m\033[0m",  // Bold
	"symlink":  "\033[36m\033[0m", // Cyan
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
