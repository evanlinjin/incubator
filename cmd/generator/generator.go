package main

import (
	"fmt"
	"github.com/kittycash/incubator"
	"math/rand"
	"os"
	"time"
)

func main() {
	in, out := os.Args[1], os.Args[2]

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	config, e := incubator.RandomKittyConfig(r, in)
	if e != nil {
		panic(e)
	}
	fmt.Println(config.Print(true))

	segments, e := incubator.NewKittySegments(in, config)
	if e != nil {
		panic(e)
	}

	if e := segments.CompileToFile(out); e != nil {
		panic(e)
	}

	fmt.Println("OKAY!")
}
