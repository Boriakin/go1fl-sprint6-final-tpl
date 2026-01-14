package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dowFile, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dowFile.Close()

	data, err := io.ReadAll(dowFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	convert, err := service.ConverterMorse(string(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fileName := time.Now().UTC().String() + filepath.Ext(header.Filename)

	outFile, err := os.Create(fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer outFile.Close()

	if _, err := outFile.WriteString(convert); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	http.ServeFile(w, r, fileName)
}
