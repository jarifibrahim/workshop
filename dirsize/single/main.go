// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 247.

//!+main

// The command computes the disk usage of the files in a directory.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func main() {
	flag.Parse()
	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	now := time.Now()

	var nfiles, nbytes int64
	for _, root := range roots {
		nf, nb := walkDir(root)
		nfiles += nf
		nbytes += nb
	}

	fmt.Println("Total time taken: ", time.Since(now))
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

//!-main

//!+walkDir

// walkDir recursively walks the file tree rooted at dir
// and returns the size of each found file and the number of files.
func walkDir(dir string) (numFiles int64, size int64) {
	time.Sleep(100 * time.Millisecond)
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			nf, fs := walkDir(subdir)
			numFiles += nf
			size += fs
		} else {
			numFiles++
			size += entry.Size()
		}
	}
	return
}

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
