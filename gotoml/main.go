package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/cage1016/conf2toml"
	"github.com/ota42y/gotoml"
)

func main() {
	var name = flag.String("name", "Foo", "the name of the struct")
	var pkg = flag.String("pkg", "main", "the name of the package for the generated code")
	var pick = flag.String("pick", "", "the name of toml property want to generate")

	flag.Parse()

	var input io.Reader = os.Stdin
	output, err := gotoml.Generate(conf2toml.NormalizationReader(input), *name, *pkg, *pick)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error parsing", err)
		os.Exit(1)
	} else {
		fmt.Print(string(output))
	}
}
