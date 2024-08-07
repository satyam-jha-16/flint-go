package main

import (
	"fmt"
	"log"
	"testing"
)

func TestIsDir(t *testing.T) {
	isDir, err := IsDir("new/a.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(isDir)
}
