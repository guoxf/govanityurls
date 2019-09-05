package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"gopkg.in/yaml.v2"
)

var host string

var m map[string]struct {
	Repo    string `yaml:"repo,omitempty"`
	Display string `yaml:"display,omitempty"`
}

func init() {
	flag.StringVar(&host, "host", "", "custom domain name, e.g. tonybai.com")

	vanity, err := ioutil.ReadFile("./vanity.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal(vanity, &m); err != nil {
		log.Fatal(err)
	}
	for _, e := range m {
		if e.Display != "" {
			continue
		}
		if strings.Contains(e.Repo, "github.com") {
			e.Display = fmt.Sprintf("%v %v/tree/master{/dir} %v/blob/master{/dir}/{file}#L{line}", e.Repo, e.Repo, e.Repo)
		}
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	current := r.URL.Path
	p, ok := m[current]
	if !ok {
		http.NotFound(w, r)
		return
	}

	if err := vanityTmpl.Execute(w, struct {
		Import  string
		Repo    string
		Display string
	}{
		Import:  host + current,
		Repo:    p.Repo,
		Display: p.Display,
	}); err != nil {
		http.Error(w, "cannot render the page", http.StatusInternalServerError)
	}
}

var vanityTmpl, _ = template.New("vanity").Parse(`<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="{{.Import}} git {{.Repo}}">
<meta name="go-source" content="{{.Import}} {{.Display}}">
<meta http-equiv="refresh" content="0; url=https://godoc.org/{{.Import}}">
</head>
<body>
Nothing to see here; <a href="https://godoc.org/{{.Import}}">see the package on godoc</a>.
</body>
</html>`)

func usage() {
	fmt.Println("govanityurls is a service that allows you to set custom import paths for your go packages\n")
	fmt.Println("Usage:")
	fmt.Println("\t govanityurls -host [HOST_NAME]\n")
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if host == "" {
		usage()
		return
	}

	http.Handle("/", http.HandlerFunc(handle))
	log.Fatalln(http.ListenAndServe("0.0.0.0:8088", nil))
}
