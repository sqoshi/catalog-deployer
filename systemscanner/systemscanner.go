package systemscanner

import (
	"fmt"
	"os"
)

/*
This function checks if given path is a directory.
*/
func isDir(path string) (bool, error) {
	//asd asda asd
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}

func listFiles(path string) ([]string, error) {
	files, err := os.ReadDir(path)

	if err != nil {
		return []string{}, err
	}

	var fileNames []string
	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}
	fmt.Println(fileNames)
	return fileNames, err

}

// tests t21321
func ReadEntry(path string) string {
	decision, err := isDir(path)
	if decision {
		listFiles(path)
	}
	fmt.Println(err)
	return ""
}
