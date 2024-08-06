package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {

	flag.Parse()

	oldRootPath := flag.Arg(0)
	newRootPath := flag.Arg(1)

	oldTree, err := GetTree(oldRootPath)
	if err != nil {
		log.Fatalf("Error getting tree: %v", err)
	}

	newTree, err := GetTree(newRootPath)
	if err != nil {
		log.Fatalf("Error getting tree: %v", err)
	}

	fmt.Println(oldTree)
	fmt.Println(newTree)
}
