package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zollidan/tornet/torrentfile"
)

func main() {
	inPath := os.Args[1]

	tf, err := torrentfile.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(tf.Name)
}

