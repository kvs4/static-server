package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kvs4/static-server/Internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIndexHandler(t *testing.T) {
	// Preparing text files
	err := os.Mkdir("./test_files", 0755)
	require.NoError(t, err, "Folder created error: %v", err)
	defer os.RemoveAll("./test_files")
	f, err := os.Create("./test_files/index.html")
	require.NoError(t, err, "File created error: %v", err)
	f.Close()

	gin.SetMode(gin.TestMode)

	fs := storage.NewFileService("./test_files")
	router := NewRouter(fs)

	req, err := http.NewRequest("GET", "/", nil)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestStaticFileHandler(t *testing.T) {
	// Preparing text files
	err := os.Mkdir("./test_files", 0755)
	require.NoError(t, err, "Folder created error: %v", err)
	defer os.RemoveAll("./test_files")
	f, err := os.Create("./test_files/test.html")
	require.NoError(t, err, "File created error: %v", err)
	f.Close()

	gin.SetMode(gin.TestMode)

	fs := storage.NewFileService("./test_files")
	router := NewRouter(fs)

	req, err := http.NewRequest("GET", "/static/test.html", nil)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestStaticFileNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	fs := storage.NewFileService("./test_files")
	router := NewRouter(fs)

	req, err := http.NewRequest("GET", "/static/missing.html", nil)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
}

func TestPageNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	fs := storage.NewFileService("./test_files")
	router := NewRouter(fs)

	req, err := http.NewRequest("GET", "/missingpage", nil)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
}

func TestIncorrectMethod(t *testing.T) {
	gin.SetMode(gin.TestMode)

	fs := storage.NewFileService("./test_files")
	router := NewRouter(fs)

	req, err := http.NewRequest("POST", "/", nil)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusMethodNotAllowed, recorder.Code)
}
