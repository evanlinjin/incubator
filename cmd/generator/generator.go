package main

import (
	"os"
	"github.com/kittycash/incubator"
	"fmt"
)

func main() {
	in := os.Args[1]
	out := os.Args[2]

	config := &incubator.KittyConfig{
		KittyID:  "2",
		HasBrows: true,
		HasCap:   true,
		BodyID:   "02",
		BrowsID:  "00",
		CapID:    "00",
		EarsID:   "01",
		EyesID:   "04",
		HeadID:   "01",
		NoseID:   "00",
		TailID:   "01",
	}

	segments, e := incubator.NewKittySegments(in, config)
	if e != nil {
		panic(e)
	}

	if e := segments.CompileToFile(out); e != nil {
		panic(e)
	}

	fmt.Println("OKAY!")
}
