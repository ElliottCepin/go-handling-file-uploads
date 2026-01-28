package main

import (
	"net/http"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)


func serveUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 * 1000 * 1000) // 10b * 1000 = 10Kb; 10Kb * 1000 = 10Mb
	
	header := r.MultipartForm.File["file"]
	if (header[0] == nil) {
		w.WriteHeader(400)
		fmt.Fprint(w, "Bad request\n")
		return
	}

	file, err := header[0].Open()

	if (err != nil) {
		w.WriteHeader(422)
		fmt.Fprint(w, "Unprossessable entity\n")
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if (err != nil) {
		w.WriteHeader(422)
		fmt.Fprint(w, "Unprossessable entity\n")
		return
	}

	path, err := os.Getwd()

	if (err != nil) {
		w.WriteHeader(500)
		fmt.Fprint(w, "Internal server error - os.Getwd\n")
		return
	}

	dir := filepath.Join(path, "/upload/" + time.Now().String()[0:20] + ".file")

	f, err := os.Create(dir)

	if (err != nil) {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Internal server error - os.Create\n%w\n", err)
		return
	}
	
	defer f.Close()

	_, err = f.Write(data)

	if (err != nil) {
		w.WriteHeader(500)
		fmt.Fprint(w, "Internal server error")
		return
	}
}

func main() {
	http.HandleFunc("/upload", serveUpload)

	http.ListenAndServe(":8080", nil)
}
