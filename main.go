package main

import (
	"flag"
	"log"
)

func main() {

	flag.Parse()

	oldRootPath := flag.Arg(0)
	newRootPath := flag.Arg(1)

	compare, err := CompareDir(oldRootPath, newRootPath)
	if err != nil {
		log.Fatal(err)
	}
	LogComparedData(compare)

}
