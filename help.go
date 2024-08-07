package main

var (
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Reset   = "\033[0m"
	Magenta = "\033[35m"
)

const helpMess = `flint- Compare Folder and Files
Usage:
$ flint [path old] [path new]

Examples:
$ flint old-hello.txt new-hello.txt
$ flint old/hello.txt new/hello.txt
$ flint old new
$ flint old/abc new/abc
`
