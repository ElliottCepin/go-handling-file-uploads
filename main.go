package main

import (
	"net/http"
	"fmt"
)


func serveUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 * 1000 * 1000) // 10b * 1000 = 10Kb; 10Kb * 1000 = 10Mb
	
}
