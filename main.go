package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

func main() {
	baseFolderPath := "."
	fileList, err := os.ReadDir(baseFolderPath)

	if err != nil {
		panic(err)
	}
	var m map[string]string
	m = make(map[string]string)

	// loop through the fileList
	for _, file := range fileList {

		if !file.IsDir() {
			data, err := os.ReadFile(baseFolderPath + file.Name())

			if err != nil {
				panic(err)
			}

			hashed := md5.Sum(data)
			// text := string(data)
			fmt.Printf("%x %s\n", hashed, file.Name())
			hashedstring := fmt.Sprint("%x", hashed)

			i, ok := m[hashedstring]

			if ok {
				fmt.Printf("Found conflicting hash %s and %s\n", i, file.Name())
			} else {
				m[hashedstring] = file.Name()
			}
		}
	}
}
