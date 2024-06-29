package pdfDocumentSituation

import (
	"net/http"
	"os"

	"github.com/charmbracelet/log"
)

func Situation() {
	log.Info("creating pdf test page")
	http.HandleFunc("/pdf", servePDF)
	log.Info("done creating pdf test ")
}

func servePDF(w http.ResponseWriter, r *http.Request) {
	// Set the correct Content-Type header
	w.Header().Set("Content-Type", "application/pdf")

	// Open the PDF file
	file, err := os.Open("situations/pdfDocumentSituation/test.pdf")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Copy the file contents to the response writer
	_, err = file.WriteTo(w)
	if err != nil {
		http.Error(w, "Error writing file", http.StatusInternalServerError)
		return
	}
}
