package textDocumentSituation

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
)

func Situation() {
	log.Info("creating text test pages")

	log.Info("setting up text pages")
	http.HandleFunc("/text", textDocumentHandler)
	http.HandleFunc("/text1", textDocumentHandler)

	log.Info("done creating text test pages")
}

func textDocumentHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world blah blah 123456789!@#$%^&*()_+~")
	// Just to get sure that your crawler isnt reading it as HTML
	fmt.Fprintln(w, "<h1>HELLO</h1>")
	fmt.Fprintln(w, `<a href="/text">LINK?</a>`)
}
