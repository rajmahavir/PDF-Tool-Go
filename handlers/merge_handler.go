package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"pdf-merger/pdf"
	"strconv"
	"time"
)

func HandleMerge(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		http.Error(w, "File too large. Maximum size is 50MB", http.StatusBadRequest)
		return
	}

	pdf1File, pdf1Header, err := r.FormFile("pdf1")
	if err != nil {
		http.Error(w, "Error retrieving first PDF", http.StatusBadRequest)
		return
	}
	defer pdf1File.Close()

	pdf2File, pdf2Header, err := r.FormFile("pdf2")
	if err != nil {
		http.Error(w, "Error retrieving second PDF", http.StatusBadRequest)
		return
	}
	defer pdf2File.Close()

	pageNumStr := r.FormValue("pageNumber")
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil || pageNum < 1 {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	if filepath.Ext(pdf1Header.Filename) != ".pdf" || filepath.Ext(pdf2Header.Filename) != ".pdf" {
		http.Error(w, "Only PDF files are allowed", http.StatusBadRequest)
		return
	}

	tempDir, err := os.MkdirTemp("", "pdfmerge-*")
	if err != nil {
		http.Error(w, "Error creating temporary directory", http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(tempDir)

	pdf1Path := filepath.Join(tempDir, "pdf1.pdf")
	pdf2Path := filepath.Join(tempDir, "pdf2.pdf")
	outputPath := filepath.Join(tempDir, "merged.pdf")

	if err := pdf.SaveFile(pdf1File, pdf1Path); err != nil {
		http.Error(w, "Error saving first PDF", http.StatusInternalServerError)
		return
	}

	if err := pdf.SaveFile(pdf2File, pdf2Path); err != nil {
		http.Error(w, "Error saving second PDF", http.StatusInternalServerError)
		return
	}

	pdf1PageCount, err := pdf.GetPageCount(pdf1Path)
	if err != nil {
		http.Error(w, "Error reading first PDF: "+err.Error(), http.StatusBadRequest)
		return
	}

	pdf2PageCount, err := pdf.GetPageCount(pdf2Path)
	if err != nil {
		http.Error(w, "Error reading second PDF: "+err.Error(), http.StatusBadRequest)
		return
	}

	if pdf1PageCount < 2 {
		http.Error(w, fmt.Sprintf("First PDF must have multiple pages (found %d page)", pdf1PageCount), http.StatusBadRequest)
		return
	}

	if pdf2PageCount < 2 {
		http.Error(w, fmt.Sprintf("Second PDF must have multiple pages (found %d page)", pdf2PageCount), http.StatusBadRequest)
		return
	}

	if pageNum > pdf1PageCount {
		http.Error(w, fmt.Sprintf("Page number %d exceeds first PDF page count (%d pages)", pageNum, pdf1PageCount), http.StatusBadRequest)
		return
	}

	if err := pdf.MergePDFs(pdf1Path, pdf2Path, outputPath, pageNum); err != nil {
		http.Error(w, "Error merging PDFs: "+err.Error(), http.StatusInternalServerError)
		return
	}

	mergedPDF, err := os.ReadFile(outputPath)
	if err != nil {
		http.Error(w, "Error reading merged PDF", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=merged_%d.pdf", time.Now().Unix()))
	w.Write(mergedPDF)
}
