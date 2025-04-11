package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	var walkDir func(string, string) error
	walkDir = func(current string, prefix string) error {
		inthatdir, err := os.ReadDir(current)
		if err != nil {
			return err
		}
		if !printFiles {
			var dirs []os.DirEntry
			for _, entry := range inthatdir {
				if entry.IsDir() {
					dirs = append(dirs, entry)
				}
			}
			inthatdir = dirs
		}
		//проход по элементам папки
		for i, entry := range inthatdir {
			islast := false
			if i == len(inthatdir)-1 {
				islast = true
			}

			var newprefix string
			if islast {
				newprefix = prefix + "└───"
			} else {
				newprefix = prefix + "├───"
			}

			if entry.IsDir() {
				fmt.Fprintf(out, "%s%s\n", newprefix, entry.Name())
				var childPrefix string
				if islast {
					childPrefix = prefix + "        "
				} else {
					childPrefix = prefix + "│       "
				}
				if err := walkDir(filepath.Join(current, entry.Name()), childPrefix); err != nil {
					return err
				}
			} else {
				var size string
				info, _ := entry.Info()
				if info.Size() == 0 {
					size = "empty"
				} else {
					size = fmt.Sprintf("%db", info.Size())
				}
				fmt.Fprintf(out, "%s%s (%s)\n", newprefix, entry.Name(), size)

			}
		}

		return nil
	}
	return walkDir(path, "")
}
