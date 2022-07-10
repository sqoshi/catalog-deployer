package system

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestShouldGetAppropriateFileTypeFromFilename(t *testing.T) {
	testCases := []struct {
		Input    string
		Expected string
	}{
		{"f1.csv", "csv"},
		{"tasd/ad.asd.txt", "txt"},
		{"f1csv", "unknown"},
	}

	for _, test := range testCases {
		if getFileType(test.Input) == test.Expected {
			t.Fatal("Strings not equal")
		}
	}
}

func TestShouldCheckPathExistence(t *testing.T) {
	if Exists("/a/b/c/d") {
		t.Fatal()
	}
	if !Exists(".") {
		t.Fatal()
	}
}

func TestShouldGetDefaultValueForVariableOrFromEnv(t *testing.T) {
	testCases := []struct {
		Name       string
		Value      string
		Default    string
		Expected   string
		ExistInEnv bool
	}{
		{"VAR0", "1", "2", "2", false},
		{"VAR1", "1", "2", "1", true},
		{"VAR2", "", "123", "123", true},
	}
	for _, test := range testCases {
		if test.ExistInEnv {
			t.Setenv(test.Name, test.Value)
		}
		if !(GetEnvOrFallback(test.Name, test.Default) == test.Expected) {
			t.Fatal("Unexpected value under environment variable")
		}
	}
}

func TestShouldRecognizeDirectory(t *testing.T) {
	sysDir, _ := os.Getwd()
	if !isDir(sysDir) {
		t.Fatal("Dir not recognized")
	}
	if isDir(filepath.Join(sysDir, "system.go")) {
		t.Fatal("Dir not recognized")
	}
}

func TestShouldListEntitiesInWd(t *testing.T) {
	sysDir, _ := os.Getwd()
	expected := []string{filepath.Join(sysDir, "system.go"), filepath.Join(sysDir, "system_test.go")}

	if !reflect.DeepEqual(expected, listEntities(sysDir)) {
		t.Fatal("Arrays not equal")
	}
}
