package main

import (
	"flag"
	"os"

	"github.com/kaneshin/identicon"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	id := args[0]

	out, err := os.Create(id + ".png")
	if err != nil {
		panic(err)
	}

	d := identicon.NewDataString(id)
	if err := d.Encode(out); err != nil {
		panic(err)
	}
}
