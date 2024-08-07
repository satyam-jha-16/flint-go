# Flint - File and Directory Comparison Tool

Flint is a powerful and efficient file and directory comparison tool written in Go. It allows users to compare individual files or entire directory structures, highlighting differences, additions, and deletions.

## Features

- Compare individual files
- Compare entire directory structures
- Highlight file/directory additions, deletions, and modifications
- Color-coded output for easy reading
- Efficient diff algorithm for quick comparisons

## Installation

To install Flint, make sure you have Go installed on your system, then run:

```
go get github.com/yourusername/flint
```
Usage```
flint [path old] [path new]```
Examples:
Compare two files:```
bashCopyflint old-hello.txt new-hello.txt```
Compare files in different directories:```
bashCopyflint old/hello.txt new/hello.txt```
Compare two directories:```
bashCopyflint old new```
Compare subdirectories:```
bashCopyflint old/abc new/abc```

Output
Flint provides a color-coded output for easy interpretation:

Red: Removed files/directories
Green: Created files/directories
Yellow: Modified files
Magenta: Diff headers

How It Works

Flint first determines if the provided paths are files or directories.
For file comparisons, it uses an efficient diff algorithm to highlight changes.
For directory comparisons, it:

Creates a tree structure of both directories
Identifies added and removed files/directories
Compares contents of files that exist in both directories



Contributing
Contributions to Flint are welcome! Please feel free to submit a Pull Request.
License
This project is licensed under the MIT License.
Acknowledgements

The diff algorithm is based on the work of Thomas G. Szymanski, as described in "A Special Case of the Maximal Common Subsequence Problem," Princeton TR #170 (January 1975).
