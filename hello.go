// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Data struct {
	Service            string
	Revision           string
	Project            string
	NetworkEgressError bool
	ProjectFound       bool
}

func main() {
	tmpl := template.Must(template.ParseFiles("index.html"))

	// Get project ID from metadata server
	project := "???"
	projectFound := false
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
		projectFound = true
	}

	// Check for network egress configuration (CR-GKE)
	networkEgressError := false
	if projectFound == false {
		// Check to see if we can reach something off the cluster e.g. www.google.com
		req, _ = http.NewRequest("HEAD", "https://www.google.com", nil)
		res, err = client.Do(req)
		if err == nil && res.StatusCode >= 200 && res.StatusCode <= 299 {
			// egress worked successfully
			log.Print("Verified that network egress is working as expected.")
		} else {
			log.Print("Network egress appears to be blocked. Unable to access https://www.google.com.")
			networkEgressError = true
		}
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
		Service:            service,
		Revision:           revision,
		Project:            project,
		NetworkEgressError: networkEgressError,
		ProjectFound:       projectFound,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Print("Hello from Cloud Run! The container started successfully and is listening for HTTP requests on $PORT.")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
