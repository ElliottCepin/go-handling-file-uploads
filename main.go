package main

import (
	"net/http"
	"fmt"
)


func serveUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 * 1000 * 1000) // 10b * 1000 = 10Kb; 10Kb * 1000 = 10Mb
	
	file, err := r.MultipartForm.Get("file").Open
	if (err != nil) {
		// do something
	}

	// to actually get something out of file, you have to do file.Read(buffer []byte),
	// but you don't really understand any of that nonsense
	// please write a blog post about io.Reader, mime/multipart's multipart.File, allocating memory in go,
	// and whatever else you need to be able to read a file and write it to disk
	// thx bestie
	// do remember that you don't have to do allocation, you can get a byteslice from io.ReadAll(file) ([]byte, err).


}
