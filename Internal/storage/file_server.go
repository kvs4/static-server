package storage

import (
	"os"
	"path/filepath"
)

type FileServer struct {
	BaseDir string
}

func NewFileService(baseDir string) *FileServer {
	return &FileServer{BaseDir: baseDir}
}

func (fs *FileServer) FileExists(fileName string) bool {
	filePath := filepath.Join(fs.BaseDir, fileName)
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
