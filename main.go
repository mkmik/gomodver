package main

import (
	"fmt"
	"log"
	"runtime/debug"
)

func main() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatal("cannot read build info")
	}
	fmt.Printf("module version %q\n", bi.Main.Version)
}
