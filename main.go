package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
)

const nPages int = 20000
const serverPort string = ":8090"

var pagesHashes []string

func hashSHA256(index int) string {
	hasher := sha256.New()
	hasher.Write([]byte(fmt.Sprintf("%d", index)))
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)
}

func generatePagesHashes(nPages int) []string {
	var localPagesHashes []string

	for nPage := 0; nPage < nPages; nPage++ {
		pageHash := hashSHA256(nPage)
		localPagesHashes = append(localPagesHashes, pageHash)
	}

	return localPagesHashes
}

func createPageHandler(pageIndex int) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "<html><body>")

		fmt.Fprintf(w, "<h1>Page %d</h1>", pageIndex+1)
		fmt.Fprintf(w, "<h3>Page hash %s</h3>", pagesHashes[pageIndex])

		var nextPageIndex int = pageIndex + 1
		if nextPageIndex < nPages {
			fmt.Fprintf(w, "<a href='/%s'>Next Page</a>", pagesHashes[nextPageIndex])

			if pageIndex == 0 {
				fmt.Fprintf(w, "<a href='/%s'>2nd Next Page</a>", pagesHashes[nextPageIndex+1])
				fmt.Fprintf(w, "<a href='/%s'>3rd Next Page</a>", pagesHashes[nextPageIndex+2])
				fmt.Fprintf(w, "<a href='/%s'>3rd Next Page</a>", pagesHashes[nextPageIndex+3])
				fmt.Fprintf(w, "<a href='/%s'>3rd Next Page</a>", pagesHashes[nextPageIndex+4])

			}
		}

		fmt.Fprintf(w, "</body></html>")
	}
}

func main() {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		Level:           log.DebugLevel,
	})
	log.SetDefault(logger)

	log.Info("setting up pages hashes", "number of pages", nPages)
	pagesHashes = generatePagesHashes(nPages)

	log.Info("creating pages")
	for pageIndex, pageHash := range pagesHashes {
		log.Info("creating page", "index", pageIndex, "hash", pageHash)
		http.HandleFunc(fmt.Sprintf("/%s", pageHash), createPageHandler(pageIndex))
	}
	log.Infof("done creating %d pages", nPages)
	log.Info("first page", "hash", pagesHashes[0])

	log.Fatal(http.ListenAndServe(serverPort, nil))
}
