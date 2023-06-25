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

	fileList, err := os.ReadDir(baseFolderPath)

	bar := pb.StartNew(len(fileList))

	if err != nil {
		panic(err)
	}
	var m map[string]string
	m = make(map[string]string)

	var duplicateFiles []string // an empty list

	// loop through the fileList
	for _, file := range fileList {

		bar.Increment()

		if !file.IsDir() {
			data, err := os.ReadFile(baseFolderPath + file.Name())

			if err != nil {
				panic(err)
			}

			hashed := md5.Sum(data)
			// text := string(data)
			// fmt.Printf("%x %s\n", hashed, file.Name())
			hashedstring := fmt.Sprint("%x", hashed)

			i, ok := m[hashedstring]

			if ok {
				fmt.Printf("Found conflicting hash %s and %s\n", i, file.Name())
				duplicateFiles = append(duplicateFiles, i+" "+file.Name()+"\n")
			} else {
				m[hashedstring] = file.Name()
			}
		}
	}

	bar.Finish()
	fmt.Printf("%d channels", duplicateFiles)
}

func findFilePaths(path string) []string {
	var files []string

	fileList, err := os.ReadDir(path)

	if err != nil {
		panic(err)
	}
	for _, file := range fileList {

		if file.IsDir() {
			files = append(files, findFilePaths(path+"/"+file.Name())...)
		} else {
			files = append(files, path+"/"+file.Name())
		}
	}

	return files

}
