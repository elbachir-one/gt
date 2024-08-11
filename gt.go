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

var dirs, files int

func parseArgs() Args {
	var args Args
	flag.BoolVar(&args.Help, "h", false, helpDescription)
	flag.BoolVar(&args.Help, "help", false, helpDescription)
	flag.BoolVar(&args.Version, "v", false, versionDescription)
	flag.BoolVar(&args.Version, "version", false, versionDescription)
	flag.BoolVar(&args.ShowHidden, "a", false, showHiddenDescription)
	flag.BoolVar(&args.ShowHidden, "all", false, showHiddenDescription)
	flag.BoolVar(&args.Unsort, "u", false, unsortDescription)
	flag.BoolVar(&args.Unsort, "unsorted", false, unsortDescription)
	flag.BoolVar(&args.Summary, "m", false, summaryDescription)
	flag.BoolVar(&args.Summary, "summary", false, summaryDescription)
	flag.BoolVar(&args.DirsOnly, "d", false, dirsOnlyDescription)
	flag.BoolVar(&args.DirsOnly, "directories", false, dirsOnlyDescription)
	flag.BoolVar(&args.FullPath, "f", false, fullPathDescription)
	flag.BoolVar(&args.FullPath, "full-path", false, fullPathDescription)
	flag.BoolVar(&args.OrderByExt, "o", false, orderByExtDescription)
	flag.BoolVar(&args.OrderByExt, "order-by-extension", false, orderByExtDescription)
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

    fmt.Printf("%s%s%s %s\n", prefix, pointer, icon, name)
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

func getDirectoryIcon(directory string) string {
	for dirType, icon := range directoryIcons {
		if filepath.Base(directory) == dirType {
			return icon
		}
	}
	return directoryIcons["default"]
}
