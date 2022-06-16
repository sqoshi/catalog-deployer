package api

import (
	"catalog-deployer/system"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const RootDirKey = "ROOT_DIR"

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func displayEntityByPath(c *gin.Context) {
	path := strings.Replace(c.Param("path"), "<slash>", "/", -1)
	abspath := filepath.Join(os.Getenv(RootDirKey), path)
	log.Println(abspath)
	if exists(abspath) {
		entity := system.GetEntityInfo(abspath)
		c.IndentedJSON(http.StatusOK, entity)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Path does not exists."})
}

func RunApi() {
	router := gin.Default()
	router.GET("/:path", displayEntityByPath)
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
		return
	}
}
