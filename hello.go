package main

import (
        "fmt"
        "log"
        "net/http"
        "html/template"
        "os"
        "io/ioutil"
)

type Data struct {
	Service string
	Revision string
    Project string
}

func main() {
        tmpl := template.Must(template.ParseFiles("index.html"))

        // Get project ID from metadata server
        project := "???"
        client := &http.Client{}
        req, _ := http.NewRequest("GET", "http://metadata.google.internal/computeMetadata/v1/project/project-id", nil)
        req.Header.Set("Metadata-Flavor", "Google")
        res, err := client.Do(req)
        if err == nil {
            defer res.Body.Close()
            responseBody, err := ioutil.ReadAll(res.Body)
            if err != nil {
                log.Fatal(err)
            }
            project = string(responseBody)
        }

        service := os.Getenv("K_SERVICE")
        if service == "" {
                service = "???"
        }

        revision := os.Getenv("K_REVISION")
        if revision == "" {
                revision = "???"
        }

        data := Data{
			Service: service,
			Revision: revision,
            Project: project,
		}

        http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
            tmpl.Execute(w, data)
        })

        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
        }

        log.Print("Hello from Cloud Run! The container started successfully and is listening for HTTP requests on $PORT.")
        log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}