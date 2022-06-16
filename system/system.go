package system

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Entity struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

func isDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return fileInfo.IsDir()
}

func getFileType(filename string) string {
	sep := "."
	if !strings.Contains(filename, sep) {
		return "unknown"
	}
	filenameParts := strings.Split(filename, sep)
	return strings.ToLower(filenameParts[len(filenameParts)-1])
}

func listEntities(dirpath string) []string {
	var entities []string
	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		entities = append(entities, filepath.Join(dirpath, file.Name()))
	}
	return entities
}

func GetEntityInfo(path string) Entity {
	var (
		eType   string
		content string
	)
	name := filepath.Base(path)
	if isDir(path) {
		eType = "directory"
		content = strings.Join(listEntities(path), ",")
	} else {
		fileContent, err := ioutil.ReadFile(path)
		if err != nil {

		}
		eType = getFileType(name)
		content = string(fileContent)
	}

	return Entity{
		Name:    name,
		Type:    eType,
		Content: content,
	}
}
