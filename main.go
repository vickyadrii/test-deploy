package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // 10 MB

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		http.Error(w, "Unable to retrieve file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileName := filepath.Base(handler.Filename)
	savePath := filepath.Join("/Users/vickyadrii/Documents/go/src/test-go/test-uploads", fileName)

	dst, err := os.Create(savePath)
	if err != nil {
		fmt.Println("Error creating the file")
		fmt.Println(err)
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		fmt.Println("Error saving the file")
		fmt.Println(err)
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	// Path ke direktori tempat file-file diunggah
	filePath := filepath.Join("/Users/vickyadrii/Documents/go/src/test-go/test-uploads", filepath.Base(r.URL.Path))

	http.ServeFile(w, r, filePath)
}

func main() {
	// Route untuk mengunggah file
	http.HandleFunc("/upload", uploadFile)

	// Route untuk menyajikan file
	http.HandleFunc("/files/", serveFile)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
