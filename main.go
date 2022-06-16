package main

import (
	"catalog-deployer/api"
	"log"
	"os"
)

func main() {
	//os.Setenv(api.RootDirKey, "./tests/test_dir")
	_, present := os.LookupEnv(api.RootDirKey)
	if !present {
		log.Fatalf("Storage filepath should be set under env variable %s\n", api.RootDirKey)
	}
	api.RunApi()
}
