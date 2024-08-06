package main

import (
	"io/fs"
	"path/filepath"
	"strings"
)

type Tree map[string]string

func GetTree(rootPath string) (Tree, error) {

	var tree = make(Tree)

	err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if path == rootPath {
			return nil
		}

		path = strings.TrimPrefix(path, rootPath)
		tree[path] = ""
		return nil
	})

	return tree, err
}
