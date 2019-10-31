package main

import (
	"flag"
	"fmt"
)

var gopherType string

func init() {
	const (
		defaultGopher = "pocket"
		usage         = "the variety of gopher"
	)
	flag.StringVar(&gopherType, "g", defaultGopher, usage)
}

func main() {
	flag.Parse()
	fmt.Println("gopherType =", gopherType)
}
