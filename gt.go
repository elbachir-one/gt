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
	helpDescription        = "Help"
	versionDescription     = "Show version"
	showHiddenDescription  = "Show hidden files"
	unsortDescription      = "Unsort files"
	summaryDescription     = "Show summary"
	dirsOnlyDescription    = "List directories only"
	fullPathDescription    = "Print full path prefix to each file"
	orderByExtDescription  = "Order files based on extension"
	depthDescription       = "Depth to which the tree should be displayed"
	defaultVersion         = "gt: v1.2.0"
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
  ".ts":         "\033[38;2;49;120;198m󰛦 \033[0m",
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
	flag.BoolVar(&args.Help, "h", false, helpDescription)
	flag.BoolVar(&args.Version, "v", false, versionDescription)
	flag.BoolVar(&args.ShowHidden, "s", false, showHiddenDescription)
	flag.BoolVar(&args.Unsort, "u", false, unsortDescription)
	flag.BoolVar(&args.Summary, "m", false, summaryDescription)
	flag.BoolVar(&args.DirsOnly, "d", false, dirsOnlyDescription)
	flag.BoolVar(&args.FullPath, "f", false, fullPathDescription)
	flag.BoolVar(&args.OrderByExt, "o", false, orderByExtDescription)
	flag.IntVar(&args.Depth, "depth", -1, depthDescription)

	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--") && len(arg) > 2 {
			for _, char := range arg[1:] {
				switch char {
				case 'h':
					args.Help = true
				case 'v':
					args.Version = true
				case 's':
					args.ShowHidden = true
				case 'u':
					args.Unsort = true
				case 'm':
					args.Summary = true
				case 'd':
					args.DirsOnly = true
				case 'f':
					args.FullPath = true
				case 'o':
					args.OrderByExt = true
				default:
					fmt.Printf("flag provided but not defined: -%c\n", char)
					os.Exit(1)
				}
			}
			os.Args = append(os.Args[:1], os.Args[2:]...)
		}
	}

	flag.Parse()

	if flag.NArg() > 0 {
		args.Dir = flag.Arg(0)
	} else {
		args.Dir = defaultDirectory
	}

	return args
}

func walk(directory, prefix string, depth int, args Args) error {
	if depth == 0 {
		return nil
	}

	entries, err := os.ReadDir(directory)
	if err != nil {
		return err
	}

	entries = filterEntries(entries, args)

	if !args.Unsort {
		sort.Slice(entries, func(i, j int) bool {
			if args.OrderByExt {
				return strings.ToLower(filepath.Ext(entries[i].Name())) < strings.ToLower(filepath.Ext(entries[j].Name()))
			}
			return entries[i].Name() < entries[j].Name()
		})
	}

	for i, entry := range entries {
		printEntry(entry, prefix, i == len(entries)-1, args)

		if entry.IsDir() {
			dirs++
			subPrefix := finalPointerSpace
			if i != len(entries)-1 {
				subPrefix = innerPointerSpace
			}

			err := walk(filepath.Join(directory, entry.Name()), prefix+subPrefix, depth-1, args)
			if err != nil {
				return err
			}
	} else {
		files++
	}
	}
	return nil
}

func filterEntries(entries []os.DirEntry, args Args) []os.DirEntry {
	filtered := make([]os.DirEntry, 0, len(entries))

	for _, entry := range entries {
		if args.DirsOnly && !entry.IsDir() {
			continue
		}

		if !args.ShowHidden && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		filtered = append(filtered, entry)
	}

	return filtered
}

func printEntry(entry os.DirEntry, prefix string, isLast bool, args Args) {
	name := entry.Name()

	if !args.FullPath {
		name = filepath.Base(name)
	} else {
		absPath, err := filepath.Abs(name)
		if err != nil {
			absPath = name
		}
		name = absPath
	}

	icon := getIcon(entry, name)
	pointer := innerPointer
	if isLast {
		pointer = finalPointer
	}

	fmt.Printf("%s%s%s%s\n", prefix, pointer, icon, name)
}

func getIcon(entry os.DirEntry, name string) string {
	if entry.Type()&os.ModeSymlink != 0 {
		return iconSymlink
	}

	if entry.IsDir() {
		return iconDirectory
	}

	if entry.Type()&0111 != 0 {
		return iconExecutable
	}

	ext := filepath.Ext(name)
	if icon, ok := icons[ext]; ok {
		return icon
	}

	return iconOther
}

func printSummary() {
	fmt.Printf("\n\033[32m%d directories, %d files\033[0m\n", dirs, files)
}

func main() {
	args := parseArgs()

	if args.Help {
		fmt.Printf("Usage: gt [OPTION]... [DIRECTORY]...\n")
		flag.PrintDefaults()
		return
	}

	if args.Version {
		fmt.Printf("%s\n", defaultVersion)
		return
	}

	err := walk(args.Dir, "", args.Depth, args)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	if args.Summary {
		printSummary()
	}
}
