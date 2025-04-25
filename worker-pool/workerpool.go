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
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Create channel to listen for signals.
var signalChan chan (os.Signal) = make(chan os.Signal, 1)

func main() {
    // SIGINT handles Ctrl+C locally.
    // SIGTERM handles Cloud Run termination signal.
    signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Output until interrupt signal
	go func() {
		for {
			fmt.Println("Cloud Run Worker Pool instance is running!")
			time.Sleep(60 * time.Second)
		}
	}()

	// Receive output from signalChan.
	sig := <-signalChan
	log.Printf("%s signal caught", sig)
	
	// Add extra handling here to clean up resources, such as flushing logs and
	// closing any database or Redis connections.

	log.Print("Worker Pool instance exited")

}

