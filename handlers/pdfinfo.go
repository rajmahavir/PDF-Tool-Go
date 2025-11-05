package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"pdf-merger/pdf"
)

const maxUploadSize = 50 * 1024 * 1024 // 50 MB

// HandlePDFInfo returns information about an uploaded PDF
func HandlePDFInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"error": "File too large"}`))
		return
	}

	pdfFile, _, err := r.FormFile("pdf")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"error": "Error retrieving PDF"}`))
		return
	}
	defer pdfFile.Close()

	tempFile, err := os.CreateTemp("", "pdfinfo-*.pdf")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"error": "Error creating temp file"}`))
		return
	}
	tempPath := tempFile.Name()
	defer os.Remove(tempPath)

	if _, err := io.Copy(tempFile, pdfFile); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"error": "Error saving PDF"}`))
		return
	}
	tempFile.Close()

	pageCount, err := pdf.GetPageCount(tempPath)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"error": "Error reading PDF"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := fmt.Sprintf(`{"pageCount": %d}`, pageCount)
	w.Write([]byte(response))
}
