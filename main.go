package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
)

// Version can be set at link time to override debug.BuildInfo.Main.Version,
// which is "(devel)" when building from within the module. See
// golang.org/issue/29814 and golang.org/issue/29228.
var Version string

func main() {
	var versionFlag = flag.Bool("version", false, "")

	flag.Parse()

	if *versionFlag {
		fmt.Println(getVersion())

		return
	}

	fmt.Println("Hello, world!")
}

func getVersion() string {
	nameBin := getNameBin()
	verBin := "(unknown)"

	if Version != "" {
		verBin = Version
	} else if buildInfo, ok := debug.ReadBuildInfo(); ok {
		verBin = buildInfo.Main.Version
	}

	return fmt.Sprintf("%s %s", nameBin, verBin)
}

func getNameBin() string {
	pathBin, err := os.Executable()
	if err != nil {
		pathBin = os.Args[0]
	}

	return filepath.Base(pathBin)
}
