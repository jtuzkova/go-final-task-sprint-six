package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	s "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func MainHandle(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Сервер не поддерживает %s запросы", req.Method)
        return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, req, "../index.html")
}

func UploadHandle(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := req.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusInternalServerError)
		return
	}

	file, handler, err := req.FormFile("myFile")
    if err != nil {
        http.Error(w, "ошибка при получении файла", http.StatusBadRequest)
        return
    }
    defer file.Close()

	buf := make([]byte, handler.Size)
    _, err = file.Read(buf)
    if err != nil {
        http.Error(w, "Ошибка при чтении файла", http.StatusInternalServerError)
        return
    }

	input := string(buf)
	result, err := s.Convert(input)
	if err != nil {
		http.Error(w, fmt.Sprintf("Conversion error: %v", err), http.StatusInternalServerError)
		return
	}

	timestamp := time.Now().UTC().Format("2006-01-02_15-04-05")
	originalExt := filepath.Ext(handler.Filename)
	outputFilename := fmt.Sprintf("result_%s%s", timestamp, originalExt)
	
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		http.Error(w, "Unable to create result file", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString(result)
	if err != nil {
		http.Error(w, "Unable to write result to file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Conversion successful!\n\nInput: %s\n\nOutput: %s\n\n", 
		input, result)

}
