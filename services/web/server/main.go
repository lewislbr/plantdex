package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".web-env")

	fmt.Println("Web server ready ✅")

	port := os.Getenv("WEB_APP_PORT")
	err := http.ListenAndServe(":"+port, serveSPA("dist"))
	if err != nil {
		log.Fatal(err)
	}
}

func serveSPA(directory string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestedPath := filepath.Join(directory, filepath.Clean(r.URL.Path))

		if filepath.Clean(r.URL.Path) == "/" {
			requestedPath = requestedPath + "/index.html"
		}
		if _, err := os.Stat(requestedPath); os.IsNotExist(err) {
			requestedPath = filepath.Join(directory, "index.html")
		}

		acceptedEncodings := r.Header.Get("Accept-Encoding")
		brotli := "br"

		if strings.Contains(acceptedEncodings, brotli) {
			serveCompressedFile := func(mimeType string) {
				w.Header().Add("Content-Encoding", brotli)
				w.Header().Add("Content-Type", mimeType)
				http.ServeFile(w, r, requestedPath+".br")
			}

			switch filepath.Ext(requestedPath) {
			case ".html":
				serveCompressedFile("text/html")
			case ".css":
				serveCompressedFile("text/css")
			case ".js":
				serveCompressedFile("application/javascript")
			case ".svg":
				serveCompressedFile("image/svg+xml")
			default:
				http.ServeFile(w, r, requestedPath)
			}
		} else {
			http.ServeFile(w, r, requestedPath)
		}
	}
}