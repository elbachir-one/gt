package main

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
	".patch":      "\033[37m \033[0m",
	".diff":       "\033[37m \033[0m",
	".tex":        "\033[38;5;180m \033[0m",
	".ini":        "\033[33m󰘓 \033[0m",
	".zst":        "\033[35m \033[0m",
	".bash":       "\033[92m \033[0m",
	".jai":        "\033[38;5;22m \033[0m",
	".swift":      "\033[38;5;214m \033[0m",
	".hs":         "\033[38;5;135m \033[0m",
	".v":          "\033[32m \033[0m",
	".lock":       "\033[38;5;172m󱧈 \033[0m",
	".ts":         "\033[34m \033[0m",
	".log":        "\033[37m \033[0m",
	".app":        "\033[37m \033[0m",
	".bat":        "\033[38;5;208m󰭟 \033[0m",
	".7z":         "\033[90m \033[0m",
	".odt":        "\033[94m󰈬 \033[0m",
	".ods":        "\033[92m󰈛 \033[0m",
	".odp":        "\033[38;5;214m󰈧 \033[0m",
    ".R":          "\033[38;2;34;104;180m \033[0m",
    ".asm":        "\033[37m \033[0m",
    ".clj":        "\033[38;2;104;180;63m \033[0m",
    ".cr":         "\033[30m \033[0m",
    ".dart":       "\033[38;5;32m \033[0m",
    ".scala":      "\033[38;5;196m \033[0m",
    ".erl":        "\033[38;5;88m \033[0m",
    ".ex":         "\033[38;5;56m \033[0m",
    ".exs":        "\033[38;5;56m \033[0m",
    ".f90":        "\033[38;5;99m󱈚 \033[0m",
    ".fs":         "\033[38;5;72m \033[0m",
    ".gd":         "\033[38;5;74m \033[0m",
    ".groovy":     "\033[38;5;15m \033[0m",
    ".jl":         "\033[38;5;250m \033[0m",
    ".kt":         "\033[38;5;250m \033[0m",
    ".lisp":       "\033[38;5;74m󰅲 \033[0m",
    ".m":          "\033[38;5;21m \033[0m",
    ".ml":         "\033[38;5;208m \033[0m",
    ".nim":        "\033[38;5;227m \033[0m",
    ".pl":         "\033[38;5;24m \033[0m",
    ".ps1":        "\033[38;5;21m󰨊 \033[0m",
    ".sql":        "\033[38;5;250m \033[0m",
    ".ejs":        "\033[38;5;227m \033[0m",
    ".torrent":    "\033[38;5;29m<U+E371> \033[0m",
    "directory":   "\033[34;1m \033[0m",
    "other":       "\033[1m \033[0m",
    "symlink":     "\033[36m \033[0m",
    "symlink_dir": "\033[36;1m \033[0m",
}

var directoryIcons = map[string]string{
	    "default":      " ",
        "Music":        "󱍙 ",
        "Downloads":    "󰉍 ",
        "Videos":       " ",
        "Documents":    " ",
        "Pictures":     " ",
        "dotfiles":     "󱗜 ",
        "Public":       " ",
	    "src":          "󰳐 ",
	    "bin":          " ",
	    "docs":         " ",
        "lib":          " ",
	    ".github":      " ",
	    ".git":         " ",
        ".config":      " ",
        ".ssh":         "󰣀 ",
        ".gnupg":       "󰢬 ",
        ".icons":       " ",
        ".fonts":       " ",
        ".cache":       "󰃨 ",
        ".emacs.d":     " ",
        ".vim":         " ",
}

const (
	iconOther              = "\033[1m \033[0m"
	iconDirectory          = "\033[34;1m \033[0m"
	iconSymlink            = "\033[36m \033[0m"
	iconSymlinkDir         = "\033[36;1m \033[0m"
	iconExecutable         = "\033[32m \033[0m"
	innerPointer           = "├── "
	finalPointer           = "└── "
	innerPointerSpace      = "│   "
	finalPointerSpace      = "    "
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
)

const (
    reset         = "\033[0m"
    black         = "\033[30m"
    red           = "\033[31m"
    green         = "\033[32m"
    yellow        = "\033[33m"
    blue          = "\033[34m"
    magenta       = "\033[35m"
    cyan          = "\033[36m"
    white         = "\033[37m"
    gray          = "\033[90m"
    orange        = "\033[38;5;208m"
    lightRed      = "\033[91m"
    lightGreen    = "\033[92m"
    lightYellow   = "\033[93m"
    lightBlue     = "\033[94m"
    lightMagenta  = "\033[95m"
    lightCyan     = "\033[96m"
    lightWhite    = "\033[97m"
    lightGray     = "\033[37m"
    lightOrange   = "\033[38;5;214m"
    lightPink     = "\033[38;5;218m"
    lightPurple   = "\033[38;5;183m"
    lightBrown    = "\033[38;5;180m"
    lightCyanBlue = "\033[38;5;117m"
    brightOrange  = "\033[38;5;214m"
    brightPink    = "\033[38;5;213m"
    brightCyan    = "\033[38;5;51m"
    brightPurple  = "\033[38;5;135m"
    brightYellow  = "\033[38;5;226m"
    brightGreen   = "\033[38;5;46m"
    brightBlue    = "\033[38;5;33m"
    brightRed     = "\033[38;5;196m"
    brightMagenta = "\033[38;5;198m"
    darkGray      = "\033[38;5;236m"
    darkOrange    = "\033[38;5;208m"
    darkGreen     = "\033[38;5;22m"
    darkCyan      = "\033[38;5;23m"
    darkMagenta   = "\033[38;5;90m"
    darkYellow    = "\033[38;5;172m"
    darkRed       = "\033[38;5;124m"
    darkBlue      = "\033[38;5;18m"
)

var specialFileIcons = map[string]string{
        "default":          white + "󱁹 " + reset,
        "Makefile":         darkBlue + " " + reset,
        "Dockerfile":       blue + " " + reset,
        "LICENSE":          gray + " " + reset,
        "config":           lightGray + " " + reset,
        "PKGBUILD":         brightBlue + "󰣇 " + reset,
        "Gemfile":          brightRed + " " + reset,
        "template":         darkRed + " " + reset,
        "TODO":             lightCyanBlue + " " + reset,
        "CHANGELOG":        lightGray + " " + reset,
        "FAQ":              lightGray + "󰦨 " + reset,
        "LEGACY":           lightBrown + "󰦨 " + reset,
        ".gitconfig":       darkOrange + " " + reset,
        ".gitignore":       darkOrange + " " + reset,
        ".gitattributes":   darkOrange + " " + reset,
        ".xinitrc":         lightGray + " " + reset,
        ".bashrc":          lightGray + "󱆃 " + reset,
        ".bash_profile":    lightGray + "󱆃 " + reset,
        ".bash_history":    lightGray + " " + reset,
        ".zshrc":           lightGray + "󱆃 " + reset,
        ".vscode":          blue + " " + reset,
        ".vimrc":           darkGreen + " " + reset,
        ".yarnrc":          lightRed + " " + reset,
        ".npmrc":           red + " " + reset,
        ".emacs":           magenta + " " + reset,
        ".inputrc":         lightGray + " " + reset,
    }
