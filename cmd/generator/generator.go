package main

import (
	"fmt"
	"github.com/kittycash/incubator"
	"os"
)

func main() {
	in := os.Args[1]
	out := os.Args[2]

	config := &incubator.KittyConfig{
		KittyID:  4,
		BodyID:   2,
		BrowsID:  -1,
		CapID:    -1,
		EarsID:   1,
		EyesID:   1,
		HeadID:   0,
		NoseID:   1,
		TailID:   2,
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
