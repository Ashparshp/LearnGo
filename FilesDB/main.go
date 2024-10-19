package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(":8080", nil)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Unable to get file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		f, err := os.Create("uploaded.txt")
		if err != nil {
			http.Error(w, "Unable to create file", http.StatusInternalServerError)
			return
		}

		defer f.Close()

		if _, err := io.Copy(f, file); err != nil {
			http.Error(w, "Unable to save file", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "File uploaded successfully")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
