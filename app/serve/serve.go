// Package serve contains two implementations of a simple HTTP web server.
// The first just says hello world, and the second uses HTML templating.
package serve

import (
	"bytes"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/alecthomas/template"
	"github.com/gorilla/mux"
)

//PageData contains the body and title of the webpage.
type PageData struct {
	Title string
	Body  string
}

// PartOne runs a web server without any templating.
func PartOne() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello 世界 from %q", html.EscapeString(r.URL.Path))
	})

	log.Println("PART ONE: Starting server of localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// PartTwo runs a web server with templating.
func PartTwo() {
	rt := mux.NewRouter().StrictSlash(true)
	rt.HandleFunc("/", Index)
	log.Println("PART TWO: Starting server of localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", rt))
}

//Index is the handler for root path.
func Index(w http.ResponseWriter, r *http.Request) {
	pd := PageData{
		Title: "Index Page",
		Body:  "This is the body of the page",
	}

	tmpl, err := htmlTemplate(pd)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(tmpl))
}

func htmlTemplate(pd PageData) (string, error) {
	html := `<HTML>
<head>
	<title>
		{{.Title}}
	</title>
</head>
<body>
	{{.Body}}
</body>
`

	tmpl, err := template.New("index").Parse(html)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer

	if err := tmpl.Execute(&out, pd); err != nil {
		return "", err
	}

	return out.String(), nil
}
