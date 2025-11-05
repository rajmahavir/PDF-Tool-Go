package main

import (
	"fmt"
	"log"
	"net/http"
	"pdf-merger/handlers"
)

func main() {
	// Register all HTTP handlers
	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/merge", handlers.HandleMergePage)
	http.HandleFunc("/remove", handlers.HandleRemovePage)
	http.HandleFunc("/credits", handlers.HandleCredits)
	http.HandleFunc("/merge-pdfs", handlers.HandleMerge)
	http.HandleFunc("/remove-pages", handlers.HandleRemovePages)
	http.HandleFunc("/pdfinfo", handlers.HandlePDFInfo)

	port := "8080"

	fmt.Printf("Server starting on:\n")
	fmt.Printf("  Local:   http://localhost:%s\n", port)
	fmt.Printf("  Network: http://[YOUR-IP-ADDRESS]:%s\n", port)
	fmt.Printf("\nReplace [YOUR-IP-ADDRESS] with your computer's IP address\n")
	fmt.Printf("To find it on Windows, run: ipconfig\n\n")

	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
