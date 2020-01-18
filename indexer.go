package main

import (
	"code.sajari.com/docconv"
	"fmt"
	//"log"
	"os"
	"path/filepath"
)

var allowedFileTypes = map[string]int {
	".pdf": 1,
}

func filesInDirectory(path string) []string {
	var files []string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if _, ok := allowedFileTypes[filepath.Ext(path)]; ok {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	return files
}

func convertToText(file string) *docconv.Response {
	res, err := docconv.ConvertPath(files[0])

	if err != nil {
		fmt.Println(err)
	}

	return res
}

func main() {
	files := filesInDirectory("/mnt/manuals/Brother/Brother Parts List 10-25-2019")
	indexFiles(files)
}
