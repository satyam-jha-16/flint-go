package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestDiff(t *testing.T){
  pathOldFile := "old/a.txt"
  pathNewFile := "new/a.txt"

  oldFile, err := os.ReadFile(pathOldFile)
  if(err != nil){
    log.Fatal(err)
  }
  newFile, err := os.ReadFile(pathNewFile)
  if err != nil {
    log.Fatal(err)
  }

  compare := Diff(pathOldFile, oldFile, pathNewFile, newFile)

  fmt.Printf("%s",compare)

  
}
