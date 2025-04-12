package main

import (
	cyoa "chooseYourOwnAdventure/story"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func main() {
	cyoa.CmdCYOA()
	// the bonus is here to compile.
	port := flag.Int("port", 1234, "the port to start the server")
	filename := flag.String("file", "gopher.json", "the JSON file with the story")
	flag.Parse()
	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	story, err := cyoa.JsonStory(file)
	if err != nil {
		panic(err)
	}
	// h := cyoa.NewHandler(story, cyoa.WithTemplate(nil))
	// with options:
	someOtherTmpl := template.Must(template.New("").Parse(someOtherTmpl_))
	h := cyoa.NewHandler(story, cyoa.WithTemplate(someOtherTmpl), cyoa.WithPathFunc(someOtherPathFn))
	mux := http.NewServeMux()
	mux.Handle("/story/", h)
	fmt.Printf("Stating the server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

func someOtherPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}

var someOtherTmpl_ = `
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Choose Your Own Adventure</title>
    </head>
    <body>
		<section class="page">
			<h1>{{.Title}}</h1>
			{{range .Paragraphs}}
				<p>{{.}}</p>
			{{end}}
			<ul>
				{{range .Options}}
					<li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
				{{end}}
			</ul>
		</section>
		<style>
			body {
				font-family: helvetica, ariel;
			}
			h1 {
				text-align: center;
				position: relative;
			}
			.page {
				width: 80%;
				max-width: 500px;
				margin: auto;
				margin-top: 40px;
				margin-bottom: 40px;
				padding: 80px;
				background: #FFFCF6;
				border: 1px solid #eee;
				box-shadow: 0 10px 6px -6px #777;
			}
			ul {
			border-top: 1px dotted #ccc;
			padding: 10px 0 0 0;
			-webkit-padding-start: 0;
			}
			li {
				padding-top: 10px
			}
			a,
			a:visited {
				text-decoration: none;
				color: #6295b5;
			}
			a:active,
			a:hover {
				color: #7792a2;
			}
			p {
				text-indent: 1em;
			}
				
		</style>
</body>
</html>`
