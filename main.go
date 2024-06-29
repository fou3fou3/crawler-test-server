package main

import (
	"net/http"
	"os"

	// "test-server/situations/robotsSituation"
	// "test-server/situations/speedSituation"

	"test-server/situations/textDocumentSituation"

	"github.com/charmbracelet/log"
)

// speed situation constants
// const nPages int = 20000

const serverPort string = ":8090"

var pagesHashes []string

func main() {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		Level:           log.DebugLevel,
	})
	log.SetDefault(logger)

	// Tests how fast is your crawler change the nPages constant at top for specefic number of pages to test on
	// speedSituation.Situation()

	// Test if your crawler is fetching and using robots correctly
	//robotsSituation.Situation()

	textDocumentSituation.Situation()

	log.Fatal(http.ListenAndServe(serverPort, nil))
}
