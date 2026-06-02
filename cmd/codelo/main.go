// Command codelo is a config- and plugin-driven coding agent CLI.
// Fork of Reasonix, with Vietnamese localization and enhanced features.
package main

import (
	"os"

	"github.com/hiulv/codelovh/internal/cli"

	_ "github.com/hiulv/codelovh/internal/provider/anthropic"
	_ "github.com/hiulv/codelovh/internal/provider/openai"
	_ "github.com/hiulv/codelovh/internal/tool/builtin"
)

var version = "dev"

func main() {
	os.Exit(cli.Run(os.Args[1:], version))
}