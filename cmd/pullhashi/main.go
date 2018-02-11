package main

import (
	"flag"
	"os"
	"runtime"

	"github.com/flyinprogrammer/pullhashi"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	var osName string
	var archName string
	var bindir string
	flag.StringVar(&osName, "os", runtime.GOOS, "the os to filter packages on")
	flag.StringVar(&archName, "arch", runtime.GOARCH, "the arch to filter packages on")
	flag.StringVar(&bindir, "bindir", pullhashi.UserBinDir(), "download the binaries to a specific folder")
	flag.Parse()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	pullhashi.DownloadAll(osName, archName, bindir)

}
