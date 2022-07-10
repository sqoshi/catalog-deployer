package system

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Entity is a directory or File representation
type Entity struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

// isDir checks if path is a directory
func isDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return fileInfo.IsDir()
}

// getFileType gouges file type, after dot suffix
func getFileType(filename string) string {
	sep := "."
	if !strings.Contains(filename, sep) {
		return "unknown"
	}
	filenameParts := strings.Split(filename, sep)
	return strings.ToLower(filenameParts[len(filenameParts)-1])
}

// listEntities lists files and directories under given path
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

// GetEntityInfo creates entity and fills its contents depending on it is a directory or file
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
			log.Fatal(err)
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

// Exists checks if given path exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
