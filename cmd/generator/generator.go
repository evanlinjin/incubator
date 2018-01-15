package main

import (
	"fmt"
	"github.com/kittycash/incubator"
	"math/rand"
	"os"
	"time"
	"github.com/skycoin/skycoin/src/util/file"
)

func main() {
	in := os.Args[1]

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	config, e := incubator.RandomKittyConfig(r, in)
	if e != nil {
		panic(e)
	}
	fmt.Println("[CONFIG]", config.Print(true))

	segments, e := incubator.NewKittySegments(in, config)
	if e != nil {
		panic(e)
	}

	configHashStr := config.Hash().Hex()
	fmt.Println("[HASH]", configHashStr)

	if e := segments.CompileToFile(configHashStr+".png"); e != nil {
		panic(e)
	}
	if e := file.SaveJSONSafe(configHashStr+".json", config, os.FileMode(0666)); e != nil {
		panic(e)
	}

	fmt.Println("OKAY!")
}
