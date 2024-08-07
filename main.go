package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {

	flag.Parse()

	oldRootPath := flag.Arg(0)

	if oldRootPath == "" {
		fmt.Println(helpMessage)
		return
	}
	isOldRootPathDir, err := IsDir(oldRootPath)
	if err != nil {
		log.Fatal(err)
	}

	newRootPath := flag.Arg(1)

	if newRootPath == "" {
		fmt.Println(helpMessage)
		return
	}
	isNewRootPathDir, err := IsDir(newRootPath)
	if err != nil {
		log.Fatal(err)
	}

	if isOldRootPathDir && isNewRootPathDir {

		compare, err := CompareDir(oldRootPath, newRootPath)
		if err != nil {
			log.Fatal(err)
		}
		LogComparedData(compare)
	}

	if !isOldRootPathDir && !isNewRootPathDir {
		compare, err := CompareFiles(oldRootPath, newRootPath)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(Yellow, "## Changed: ", Reset)
		fmt.Printf("%s\n", compare)
	}

	if (isOldRootPathDir && !isNewRootPathDir) || (!isOldRootPathDir && isNewRootPathDir) {
		fmt.Println(Red, "Cannot compare a file with a directory", Reset)
	}
}
