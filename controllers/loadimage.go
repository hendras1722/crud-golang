package controllers

import (
	"net/http"
	"path/filepath"
)

// StaticFileHandler menyajikan file statis dari folder uploads
func StaticFileHandler() http.Handler {
	absPath, _ := filepath.Abs("uploads")
	return http.StripPrefix("/uploads/", http.FileServer(http.Dir(absPath)))
}
