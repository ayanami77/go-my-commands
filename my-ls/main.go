package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"time"
)

func main() {
	isListed := flag.Bool("l", false, "List files")
	flag.Parse()

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	files, _ := os.ReadDir(dir)

	if *isListed {
		listFilesWithDetails(files)
		return
	}

	listFiles(files)
}

func listFiles(files []fs.DirEntry) {
	for _, f := range files {
		fmt.Printf("%s ", f.Name())
	}
}

func listFilesWithDetails(files []fs.DirEntry) {
	for _, f := range files {
		info, err := f.Info()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s %s %d %s\n", info.Mode(), formatDate(info.ModTime()), info.Size(), f.Name())
	}
}

func formatDate(t time.Time) string {
	return t.Format("Jan 02 15:04")
}
