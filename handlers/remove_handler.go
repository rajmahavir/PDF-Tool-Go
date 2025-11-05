package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"pdf-merger/pdf"
	"time"
)

func HandleRemovePages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		http.Error(w, "File too large. Maximum size is 50MB", http.StatusBadRequest)
		return
	}

	pdfFile, pdfHeader, err := r.FormFile("pdf")
	if err != nil {
		http.Error(w, "Error retrieving PDF", http.StatusBadRequest)
		return
	}
	defer pdfFile.Close()

	pagesToRemoveStr := r.FormValue("pagesToRemove")
	if pagesToRemoveStr == "" {
		http.Error(w, "No pages specified for removal", http.StatusBadRequest)
		return
	}

	if filepath.Ext(pdfHeader.Filename) != ".pdf" {
		http.Error(w, "Only PDF files are allowed", http.StatusBadRequest)
		return
	}

	tempDir, err := os.MkdirTemp("", "pdfremove-*")
	if err != nil {
		http.Error(w, "Error creating temporary directory", http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(tempDir)

	inputPath := filepath.Join(tempDir, "input.pdf")
	outputPath := filepath.Join(tempDir, "output.pdf")

	if err := pdf.SaveFile(pdfFile, inputPath); err != nil {
		http.Error(w, "Error saving PDF", http.StatusInternalServerError)
		return
	}

	pageCount, err := pdf.GetPageCount(inputPath)
	if err != nil {
		http.Error(w, "Error reading PDF: "+err.Error(), http.StatusBadRequest)
		return
	}

	pagesToRemove := pdf.ParsePageNumbers(pagesToRemoveStr)
	if len(pagesToRemove) == 0 {
		http.Error(w, "Invalid page numbers", http.StatusBadRequest)
		return
	}

	if len(pagesToRemove) >= pageCount {
		http.Error(w, "Cannot remove all pages. At least one page must remain.", http.StatusBadRequest)
		return
	}

	if err := pdf.RemovePages(inputPath, outputPath, pagesToRemove); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	modifiedPDF, err := os.ReadFile(outputPath)
	if err != nil {
		http.Error(w, "Error reading modified PDF", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=modified_%d.pdf", time.Now().Unix()))
	w.Write(modifiedPDF)
}
