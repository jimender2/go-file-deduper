package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"os"

	"github.com/cheggaaa/pb/v3"
)

func main() {

	baseFolderPath := ""
	flag.StringVar(&baseFolderPath, "folder", ".\\", "Specifies a folder to find duplicate files in")
	flag.Parse()

	fileList := findFilePaths(baseFolderPath)

	bar := pb.StartNew(len(fileList))

	var m map[string]string
	m = make(map[string]string)

	var duplicateFiles []string // an empty list

	// loop through the fileList
	for _, file := range fileList {

		bar.Increment()

		data, err := os.ReadFile(file)

		if err != nil {
			panic(err)
		}

		hashed := md5.Sum(data)
		hashedstring := fmt.Sprint("%x", hashed)

		i, ok := m[hashedstring]

		if ok {
			duplicateFiles = append(duplicateFiles, i+" "+file+"\n")
		} else {
			m[hashedstring] = file
		}

	}

	bar.Finish()

	for _, i := range duplicateFiles {
		fmt.Printf("Found Conflicting hash %s\n", i)
	}

	fmt.Printf("%i", len(duplicateFiles))
}

func findFilePaths(path string) []string {
	var files []string

	fileList, err := os.ReadDir(path)

	if err != nil {
		panic(err)
	}

	if path[len(path)-1:] != "\\" {
		path = path + "\\"
	}

	for _, file := range fileList {

		if file.IsDir() {
			files = append(files, findFilePaths(path+file.Name())...)
		} else {
			files = append(files, path+file.Name())
		}
	}

	return files

}
