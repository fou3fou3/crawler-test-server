package speedSituation

import (
	"fmt"
	"net/http"
	"test-server/common"

	"github.com/charmbracelet/log"
)

func Situation(nPages int) {
	log.Info("setting up pages hashes", "number of pages", nPages)
	pagesHashes := common.GeneratePagesHashes(nPages)

	log.Info("creating pages")
	for pageIndex, pageHash := range pagesHashes {
		log.Info("creating page", "index", pageIndex, "hash", pageHash)
		http.HandleFunc(fmt.Sprintf("/%s", pageHash), createPageHandler(pageIndex, nPages, &pagesHashes))
	}
	log.Infof("done creating %d pages", nPages)
	log.Info("first page", "hash", pagesHashes[0])
}

func createPageHandler(pageIndex int, nPages int, pagesHashes *[]string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "<html><body>")

		fmt.Fprintf(w, "<h1>Page %d</h1>", pageIndex+1)
		fmt.Fprintf(w, "<h3>Page hash %s</h3>", (*pagesHashes)[pageIndex])

		var nextPageIndex int = pageIndex + 1
		if nextPageIndex < nPages {
			fmt.Fprintf(w, "<a href='/%s'>Next Page</a>", (*pagesHashes)[nextPageIndex])

			if pageIndex == 0 {
				fmt.Fprintf(w, "<a href='/%s'>2nd Next Page</a>", (*pagesHashes)[nextPageIndex+1])
				fmt.Fprintf(w, "<a href='/%s'>3rd Next Page</a>", (*pagesHashes)[nextPageIndex+2])
				fmt.Fprintf(w, "<a href='/%s'>3rd Next Page</a>", (*pagesHashes)[nextPageIndex+3])
				fmt.Fprintf(w, "<a href='/%s'>3rd Next Page</a>", (*pagesHashes)[nextPageIndex+4])

			}
		}

		fmt.Fprintf(w, "</body></html>")
	}
}
