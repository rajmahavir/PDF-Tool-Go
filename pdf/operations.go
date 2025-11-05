package pdf

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

// SaveFile saves an uploaded file to the specified destination
func SaveFile(src io.Reader, dst string) error {
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

// GetPageCount returns the number of pages in a PDF
func GetPageCount(pdfPath string) (int, error) {
	ctx, err := api.ReadContextFile(pdfPath)
	if err != nil {
		return 0, err
	}
	return ctx.PageCount, nil
}

// MergePDFs merges two PDFs by inserting pdf2 into pdf1 at the specified page
func MergePDFs(pdf1Path, pdf2Path, outputPath string, insertAfterPage int) error {
	ctx1, err := api.ReadContextFile(pdf1Path)
	if err != nil {
		return fmt.Errorf("failed to read first PDF: %w", err)
	}

	ctx2, err := api.ReadContextFile(pdf2Path)
	if err != nil {
		return fmt.Errorf("failed to read second PDF: %w", err)
	}

	tempDir := filepath.Dir(pdf1Path)
	part1Path := filepath.Join(tempDir, "part1.pdf")
	part2Path := filepath.Join(tempDir, "part2.pdf")

	if insertAfterPage > 0 {
		pageRange := fmt.Sprintf("1-%d", insertAfterPage)
		if err := api.TrimFile(pdf1Path, part1Path, []string{pageRange}, nil); err != nil {
			return fmt.Errorf("failed to create first part: %w", err)
		}
	}

	if insertAfterPage < ctx1.PageCount {
		pageRange := fmt.Sprintf("%d-%d", insertAfterPage+1, ctx1.PageCount)
		if err := api.TrimFile(pdf1Path, part2Path, []string{pageRange}, nil); err != nil {
			return fmt.Errorf("failed to create second part: %w", err)
		}
	}

	var filesToMerge []string
	if insertAfterPage > 0 {
		filesToMerge = append(filesToMerge, part1Path)
	}
	filesToMerge = append(filesToMerge, pdf2Path)
	if insertAfterPage < ctx1.PageCount {
		filesToMerge = append(filesToMerge, part2Path)
	}

	if err := api.MergeCreateFile(filesToMerge, outputPath, false, nil); err != nil {
		return fmt.Errorf("failed to merge PDFs: %w", err)
	}

	if err := api.ValidateFile(outputPath, model.NewDefaultConfiguration()); err != nil {
		return fmt.Errorf("merged PDF validation failed: %w", err)
	}

	_ = ctx1
	_ = ctx2

	return nil
}

// ParsePageNumbers parses a comma-separated string of page numbers
func ParsePageNumbers(pageStr string) []int {
	var pages []int
	parts := strings.Split(pageStr, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		pageNum, err := strconv.Atoi(part)
		if err == nil && pageNum > 0 {
			pages = append(pages, pageNum)
		}
	}

	return pages
}

// GetPagesToKeep returns the pages to keep given total pages and pages to remove
func GetPagesToKeep(totalPages int, pagesToRemove []int) []int {
	removeMap := make(map[int]bool)
	for _, page := range pagesToRemove {
		removeMap[page] = true
	}

	var pagesToKeep []int
	for i := 1; i <= totalPages; i++ {
		if !removeMap[i] {
			pagesToKeep = append(pagesToKeep, i)
		}
	}

	return pagesToKeep
}

// CreatePageRanges converts a slice of page numbers into page range strings
func CreatePageRanges(pages []int) []string {
	if len(pages) == 0 {
		return nil
	}

	var ranges []string
	start := pages[0]
	end := pages[0]

	for i := 1; i < len(pages); i++ {
		if pages[i] == end+1 {
			end = pages[i]
		} else {
			if start == end {
				ranges = append(ranges, fmt.Sprintf("%d", start))
			} else {
				ranges = append(ranges, fmt.Sprintf("%d-%d", start, end))
			}
			start = pages[i]
			end = pages[i]
		}
	}

	if start == end {
		ranges = append(ranges, fmt.Sprintf("%d", start))
	} else {
		ranges = append(ranges, fmt.Sprintf("%d-%d", start, end))
	}

	return ranges
}

// RemovePages removes specified pages from a PDF
func RemovePages(inputPath, outputPath string, pagesToRemove []int) error {
	pageCount, err := GetPageCount(inputPath)
	if err != nil {
		return fmt.Errorf("error reading PDF: %w", err)
	}

	if len(pagesToRemove) >= pageCount {
		return fmt.Errorf("cannot remove all pages, at least one page must remain")
	}

	pagesToKeep := GetPagesToKeep(pageCount, pagesToRemove)
	if len(pagesToKeep) == 0 {
		return fmt.Errorf("no pages would remain after removal")
	}

	pageRanges := CreatePageRanges(pagesToKeep)
	if err := api.TrimFile(inputPath, outputPath, pageRanges, nil); err != nil {
		return fmt.Errorf("error removing pages: %w", err)
	}

	return nil
}
