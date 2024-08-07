package main

import (
	"fmt"
)

type Compare struct {
	RemovedFile []string
	CreatedFile []string
	Updated     [][]byte
}

func CompareDir(oldRootPath string, newRootPath string) (*Compare, error) {

	compare := &Compare{}

	oldTree, err := GetTree(oldRootPath)
	if err != nil {
		return nil, err
	}

	newTree, err := GetTree(newRootPath)
	if err != nil {
		return nil, err
	}

	// checking if tyhe file in old folder exists in new foleder or not

	var sameTree = make(Tree)

	for path := range newTree {
		var _, ok = oldTree[path]
		if ok {
			sameTree[path] = ""

			delete(oldTree, path)
			delete(newTree, path)
		}

	}

	for path := range oldTree {
		compare.RemovedFile = append(compare.RemovedFile, path)
	}
	fmt.Println()

	for path := range newTree {
		compare.CreatedFile = append(compare.CreatedFile, path)
	}

	for path := range sameTree {

		pathOldFile := oldRootPath + path

		oldFile, err := GetFileInfo(pathOldFile)
		if err != nil {
			return nil, err
		}
		if oldFile.IsDir {
			break
		}

		pathNewFile := newRootPath + path

		newFile, err := GetFileInfo(pathNewFile)
		if err != nil {
			return nil, err
		}
		if newFile.IsDir {
			break
		}

		diff := Diff(pathOldFile, oldFile.Data, pathNewFile, newFile.Data)

		if diff != nil {

			compare.Updated = append(compare.Updated, diff)
		}

	}

	return compare, nil

}

func LogComparedData(compare *Compare) {
	fmt.Println(Red + "## Removed:" + Reset)

	for _, path := range compare.RemovedFile {
		fmt.Printf("- %s\n", path)
	}
	fmt.Println()
	fmt.Println(Green + "## Created:" + Reset)
	for _, path := range compare.CreatedFile {
		fmt.Printf("+ %s\n", path)
	}

	fmt.Println()
	fmt.Println(Yellow + "## Changed:" + Reset)
	for _, path := range compare.Updated {
		fmt.Printf("%s", path)
	}
}
