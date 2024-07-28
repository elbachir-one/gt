package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	helpDescription        = "show help"
	versionDescription     = "show version"
	showHiddenDescription  = "show hidden files"
	unsortDescription      = "unsort files"
	summaryDescription     = "show summary"
	dirsOnlyDescription    = "list directories only"
	fullPathDescription    = "print full path prefix to each file"
	orderByExtDescription  = "order files based on extension"
	depthDescription       = "depth to which the tree should be displayed"
	defaultVersion         = "gt: v0.4.1"
	defaultDirectory       = "."
	iconOther              = "\033[1m \033[0m"
	iconDirectory          = "\033[34;1m \033[0m"
	iconSymlink            = "\033[36m \033[0m"
	iconExecutable         = "\033[32m \033[0m"
	innerPointer           = "├── "
	finalPointer           = "└── "
	innerPointerSpace      = "│   "
	finalPointerSpace      = "    "
)

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
	".iso":        "\033[37m󰨣 \033[0m",
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
	".mkv":        "\033[33m󰃽 \033[0m",
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
	".exe":        "\033[34m \033[0m",
	"directory":   "\033[34;1m \033[0m",
	"other":       "\033[1m \033[0m",
	"symlink":     "\033[36m \033[0m",
}

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

var dirs, files int

func parseArgs() Args {
	var args Args
	flag.BoolVar(&args.Help, "?", false, helpDescription)
	flag.BoolVar(&args.Version, "v", false, versionDescription)
	flag.BoolVar(&args.ShowHidden, "s", false, showHiddenDescription)
	flag.BoolVar(&args.Unsort, "u", false, unsortDescription)
	flag.BoolVar(&args.Summary, "m", false, summaryDescription)
	flag.BoolVar(&args.DirsOnly, "d", false, dirsOnlyDescription)
	flag.BoolVar(&args.FullPath, "f", false, fullPathDescription)
	flag.BoolVar(&args.OrderByExt, "o", false, orderByExtDescription)
	flag.IntVar(&args.Depth, "t", -1, depthDescription)

	flag.Parse()

	if len(flag.Args()) > 0 {
		args.Dir = flag.Args()[0]
	} else {
		args.Dir = defaultDirectory
	}

	return args
}

func walk(directory, prefix string, args Args, currentDepth int) error {
	if args.Depth != -1 && currentDepth > args.Depth {
		return nil
	}

	entries, err := os.ReadDir(directory)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %w", directory, err)
	}

	filesList := filterEntries(entries, args)
	sortEntries(filesList, args)

	for i, entry := range filesList {
		printEntry(entry, prefix, i == len(filesList)-1, directory, args)
		if entry.IsDir() {
			if err := walk(filepath.Join(directory, entry.Name()), prefix+getNextPrefix(i, len(filesList)), args, currentDepth+1); err != nil {
				return err
			}
		}
	}

	return nil
}

func filterEntries(entries []os.DirEntry, args Args) []os.DirEntry {
	var filteredEntries []os.DirEntry
	for _, entry := range entries {
		if args.ShowHidden || entry.Name()[0] != '.' {
			if args.DirsOnly && !entry.IsDir() {
				continue
			}
			filteredEntries = append(filteredEntries, entry)
		}
	}
	return filteredEntries
}

func sortEntries(entries []os.DirEntry, args Args) {
	if !args.Unsort {
		sort.Slice(entries, func(i, j int) bool {
			if args.OrderByExt {
				extI := filepath.Ext(entries[i].Name())
				extJ := filepath.Ext(entries[j].Name())
				if extI == extJ {
					return strings.ToLower(entries[i].Name()) < strings.ToLower(entries[j].Name())
				}
				return extI < extJ
			}
			return strings.ToLower(entries[i].Name()) < strings.ToLower(entries[j].Name())
		})
	}
}

func printEntry(entry os.DirEntry, prefix string, isLast bool, directory string, args Args) {
	icon := getIcon(entry)
	fullPath := getFullPath(entry, directory, args)

	pointer := getPointer(isLast)
	fmt.Printf("%s%s%s%s\n", prefix, pointer, icon, fullPath)

	if entry.IsDir() {
		dirs++
	} else {
		files++
	}
}

func getIcon(entry os.DirEntry) string {
	if entry.IsDir() {
		return iconDirectory
	} else if entry.Type()&os.ModeSymlink != 0 {
		return iconSymlink
	} else {
		ext := filepath.Ext(entry.Name())
		if icon, ok := icons[ext]; ok {
			return icon
		}
		if entry.Type().Perm()&0111 != 0 {
			return iconExecutable
		}
		return iconOther
	}
}

func getFullPath(entry os.DirEntry, directory string, args Args) string {
	if args.FullPath {
		return filepath.Join(directory, entry.Name())
	}
	return entry.Name()
}

func getPointer(isLast bool) string {
	if isLast {
		return finalPointer
	}
	return innerPointer
}

func getNextPrefix(index, length int) string {
	if index == length-1 {
		return finalPointerSpace
	}
	return innerPointerSpace
}

func main() {
	args := parseArgs()

	if args.Help {
		flag.Usage()
		return
	}

	if args.Version {
		fmt.Println(defaultVersion)
		return
	}

	if err := walk(args.Dir, "", args, 1); err != nil {
		fmt.Println("Error:", err)
	}

	if args.Summary {
		fmt.Printf("\n%d directories, %d files\n", dirs, files)
	}
}
