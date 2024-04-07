package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

var (
	version   string = "dev"
	commitSHA string = "none"
	buildDate string = "unknown"
)

func main() {
	// Define a flag for --version
	showVersion := flag.Bool("version", false, "Prints the version information")

	flag.Parse()

	if *showVersion {
		fmt.Printf("Gaia %s, commit %s, built at %s", version, commitSHA, buildDate)
		os.Exit(0)
	}

	fmt.Println("This is a template version")
	t := T{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("--- t:\n%v\n\n", t)
}
