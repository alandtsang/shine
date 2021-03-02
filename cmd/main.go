package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/alandtsang/shine/internal/project"
	"github.com/alandtsang/shine/logs"
)

func main() {
	args := parseFlags()
	projectPath, err := project.InitProject(args)
	if err != nil {
		logs.Fatal(err)
	}

	logs.Infof("project path: %s", projectPath)
}

func parseFlags() []string {
	pkgName := flag.String("pkg", "myproject", "package name")
	author := flag.String("author", "YourName", "author name")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	return []string{*pkgName, *author}
}
