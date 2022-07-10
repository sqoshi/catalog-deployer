package main

import (
	"catalog-deployer/api"
	"catalog-deployer/system"
	"log"
	"os"
)

// createStorageIfNotExists creates storage directory if not exists
func createStorageIfNotExists(path string) {
	if !system.Exists(path) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			log.Fatal(err)
		}
		log.Println("Successfully created storage.")
	} else {
		log.Println("Storage already exists.")
	}
}

func main() {
	//os.Setenv(api.RootDirKey, "./tests/test_dir")
	storage, present := os.LookupEnv(api.RootDirKey)
	if !present {
		log.Fatalf("Storage filepath should be set under env variable %s\n", api.RootDirKey)
	}
	createStorageIfNotExists(storage)
	api.RunAPI()
}
