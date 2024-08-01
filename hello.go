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
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	cloudeventsClient "github.com/cloudevents/sdk-go/v2/client"
)

type Data struct {
	Service            string
	Revision           string
	Project            string
	Region             string
	AuthenticatedEmail string
	Color              string
}

func handleReceivedEvent(ctx context.Context, event cloudevents.Event) {
	type LoggedEvent struct {
		Severity  string            `json:"severity"`
		EventType string            `json:"eventType"`
		Message   string            `json:"message"`
		Event     cloudevents.Event `json:"event"`
	}
	type PubSubMessage struct {
		Data string `json:"data"`
	}
	type PubSubEvent struct {
		Message PubSubMessage `json:"message"`
	}

	dataMessage := event.Data()

	// In case of PubSub event, decode its payload to be printed in the message as-is.
	if event.Type() == "google.cloud.pubsub.topic.v1.messagePublished" {
		obj := &PubSubEvent{}
		if err := event.DataAs(obj); err != nil {
			fmt.Printf("Unable to decode PubSub data: %s\n", err.Error())
		}
		decodedMessage, decodingError := base64.StdEncoding.DecodeString(obj.Message.Data)
		if decodingError != nil {
			fmt.Printf("Unable to decode PubSub message: %s\n", decodingError.Error())
		} else {
			dataMessage = decodedMessage
		}
	}

	loggedEvent := LoggedEvent{
		Severity:  "INFO",
		EventType: event.Type(),
		Message:   fmt.Sprintf("Received event of type %s. Event data: %s", event.Type(), dataMessage),
		Event:     event, // Always log full event data
	}
	jsonLog, err := json.Marshal(loggedEvent)
	if err != nil {
		fmt.Printf("Unable to log event to JSON: %s\n", err.Error())
	} else {
		fmt.Printf("%s\n", jsonLog)
	}
}

func getEventsHandler() *cloudeventsClient.EventReceiver {
	ctx := context.Background()
	p, err := cloudevents.NewHTTP()
	if err != nil {
		log.Fatalf("failed to create http listener for receiving events: %s", err.Error())
	}
	h, err := cloudevents.NewHTTPReceiveHandler(ctx, p, handleReceivedEvent)
	if err != nil {
		log.Fatalf("failed to create handler for receiving events: %s", err.Error())
	}
	return h
}

func main() {
	tmpl := template.Must(template.ParseFiles("./index.html"))

	// Get project ID from metadata server
	project := ""
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://metadata.google.internal/computeMetadata/v1/project/project-id", nil)
	req.Header.Set("Metadata-Flavor", "Google")
	res, err := client.Do(req)
	if err == nil {
		defer res.Body.Close()
		if res.StatusCode == 200 {
			responseBody, err := io.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}
			project = string(responseBody)
		}
	}

	// Get region from metadata server
	region := ""
	req, _ = http.NewRequest("GET", "http://metadata.google.internal/computeMetadata/v1/instance/region", nil)
	req.Header.Set("Metadata-Flavor", "Google")
	res, err = client.Do(req)
	if err == nil {
		defer res.Body.Close()
		if res.StatusCode == 200 {
			responseBody, err := io.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}
			region = regexp.MustCompile(`projects/[^/]*/regions/`).ReplaceAllString(string(responseBody), "")
		}
	}
	if region == "" {
		// Fallback: get "zone" from metadata server (running on VM e.g. Cloud Run for Anthos)
		req, _ = http.NewRequest("GET", "http://metadata.google.internal/computeMetadata/v1/instance/zone", nil)
		req.Header.Set("Metadata-Flavor", "Google")
		res, err = client.Do(req)
		if err == nil {
			defer res.Body.Close()
			if res.StatusCode == 200 {
				responseBody, err := io.ReadAll(res.Body)
				if err != nil {
					log.Fatal(err)
				}
				region = regexp.MustCompile(`projects/[^/]*/zones/`).ReplaceAllString(string(responseBody), "")
			}
		}
	}

	service := os.Getenv("K_SERVICE")
	revision := os.Getenv("K_REVISION")

	color := os.Getenv("COLOR")

	data := Data{
		Service:  service,
		Revision: revision,
		Project:  project,
		Region:   region,
		Color:    color,
	}

	eventsHandler := getEventsHandler()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.Header.Get("ce-type") != "" {
			// Handle cloud events.
			eventsHandler.ServeHTTP(w, r)
			return
		}
		// Default handler (hello page).
		data.AuthenticatedEmail = r.Header.Get("X-Goog-Authenticated-User-Email") // set when behind IAP
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "User-agent: *\nDisallow: /\n")
	})

	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	address := fmt.Sprintf(":%s", port)

	log.Printf("Hello from Cloud Run! The container started successfully and is listening for HTTP requests on port %s", port)
	log.Fatal(http.ListenAndServe(address, nil))
}
