package server

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/kvs4/static-server/Internal/storage"
)

func NewRouter(fs *storage.FileServer) *gin.Engine {

	router := gin.Default()
	router.HandleMethodNotAllowed = true

	router.GET("/", func(c *gin.Context) {
		IndexHandler(c, fs)
	})

	router.StaticFS("/static", http.Dir(fs.BaseDir))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
	})

	return router
}

func IndexHandler(c *gin.Context, fs *storage.FileServer) {
	if fs.FileExists("index.html") {
		c.File(path.Join(fs.BaseDir, "index.html"))
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "index.html not found"})
	}
}
