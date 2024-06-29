package robotsSituation

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
)

func Situation() {
	log.Info("creating robots test pages")

	log.Info("setting up robots page")
	http.HandleFunc("/robots.txt", robotsHandler)

	log.Info("setting up home page (allowed page)")
	http.HandleFunc("/", homeHandler)

	log.Info("setting up unallowed page (unallowed page)")
	http.HandleFunc("/unallowed", unallowedHandler)

	log.Info("done creating robots test pages")
}

func robotsHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, `User-agent: *`)
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, `Disallow: /unallowed`)
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<html><head>")

	// Not neccessary just to test youre metadata collection
	fmt.Fprintf(w, `<title> Robot Test Situation </title>`)
	fmt.Fprintf(w, `<meta property="og:description" content="Home page for robots test situation">`)
	fmt.Fprintf(w, `<meta property="og:site_name" content="Robots Test Situation">`)
	fmt.Fprintf(w, `<link rel="shortcut icon" href="https://cdn-icons-png.flaticon.com/512/2433/2433036.png">`)
	fmt.Fprintf(w, "</head>")
	// Not neccessary ends

	// You can remove the lorem ipsum text and just keep the link this is just to check if the page text collecting is working
	fmt.Fprintf(w, "<body>")
	fmt.Fprintf(w, `<h3>Lorem ipsum dolor sit amet, consectetur adipiscing elit.<br>
					Etiam id tempor nulla, eget egestas velit.<br> 
					Donec placerat eget sem sed finibus.<br>
					Nam id laoreet sapien.<br>
					Etiam condimentum rhoncus tincidunt est.<br>
					UNALLOWED PAGE LINK <a href="/unallowed"> HERE DONT CLICK </a></h3>`)
	fmt.Fprintf(w, "</body></html>")
}

func unallowedHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<html><body><h1>UNALLOWED</h1></body></html>")
}
