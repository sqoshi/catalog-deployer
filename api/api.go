package api

import (
	"catalog-deployer/system"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const RootDirKey = "ROOT_DIR"

// displayEntityByPath displays contents of given path
func displayEntityByPath(c *gin.Context) {
	path := strings.Replace(c.Param("path"), "<slash>", "/", -1)
	abspath := filepath.Join(os.Getenv(RootDirKey), path)
	log.Println(abspath)
	if system.Exists(abspath) {
		entity := system.GetEntityInfo(abspath)
		c.IndentedJSON(http.StatusOK, entity)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Path does not exists."})
}

// displayRootDir displays content of root directory, must be in different endpoint because of not allowed empty path
func displayRootDir(c *gin.Context) {
	abspath := filepath.Join(os.Getenv(RootDirKey))
	log.Println(abspath)
	if system.Exists(abspath) {
		entity := system.GetEntityInfo(abspath)
		c.IndentedJSON(http.StatusOK, entity)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Something went wrong restart service please .."})
}

func GetEnvOrFallback(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

// RunApi deploys api with endpoints
func RunApi() {
	router := gin.Default()
	router.GET("/", displayRootDir)
	router.GET("/:path", displayEntityByPath)
	addr := fmt.Sprintf("%s:%s", GetEnvOrFallback("DEPLOY_HOST", "0.0.0.0"), GetEnvOrFallback("DEPLOY_PORT", "8080"))
	err := router.Run(addr)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Successfully deployed on %s", addr)
}
