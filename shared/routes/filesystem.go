package routes

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type NeuteredFileSystem struct {
	Fs http.FileSystem
}

func (nfs NeuteredFileSystem) Open(path string) (http.File, error) {
	cleanPath := filepath.Clean(path)
	fullPath := filepath.Join("./ui/static", cleanPath)

	// Enforce that the requested file is still within the static directory
	staticRoot, _ := filepath.Abs("./ui/static")
	absPath, _ := filepath.Abs(fullPath)
	if !strings.HasPrefix(absPath, staticRoot) {
		return nil, os.ErrPermission
	}

	f, err := nfs.Fs.Open(cleanPath)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, err
	}

	if s.IsDir() {
		index := filepath.Join(cleanPath, "index.html")
		if _, err := nfs.Fs.Open(index); err != nil {
			f.Close()
			return nil, err
		}
	}

	return f, nil
}
