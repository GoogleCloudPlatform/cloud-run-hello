package main

import (
        "fmt"
        "log"
        "net/http"
        "html/template"
        "os"
)

type Data struct {
	Service string
	Revision string
}

func main() {
        tmpl := template.Must(template.ParseFiles("index.html"))

		data := Data{
			Service: os.Getenv("K_SERVICE"),
			Revision: os.Getenv("K_REVISION"),
		}

        http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
            tmpl.Execute(w, data)
        })

        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
        }

        log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
        log.Print("Hello from Cloud Run!")
}