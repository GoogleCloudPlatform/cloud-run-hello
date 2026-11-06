# Cloud Run "Hello" container source code

This repository contains the source code of multiple sample applications.

## Hello service

A sample service implemented in Go and distributed as a public container image. It is used in the [Cloud Run quickstart](https://cloud.google.com/run/docs/quickstarts/deploy-container) and is a suggested container image in the Cloud Run UI on Cloud Console.

* **Container Image:** `us-docker.pkg.dev/cloudrun/container/hello`

<a href="https://deploy.cloud.run"><img src="https://deploy.cloud.run/button.svg" alt="Run on Google Cloud" height="40"/></a>

### Configuration Options

Set the `COLOR` environment variable to a valid CSS color to change the background color.

Set `ACK_EVENTS=false` to return HTTP 503 instead of successful POST statuses
without changing event logging or response bodies. The hello image defaults to
`ACK_EVENTS=true`. Invalid or empty values cause startup to fail.

### Response Formats

By default, the service returns answers in HTML format. You can also request responses in JSON or plain text format by passing a URL query parameter or an HTTP header:

* **URL Query Parameter:** Add `?format=json`, `?format=text` (or `plain`), or `?format=html`.
* **Standard Header:** Pass `Accept: application/json`, `Accept: text/plain`, or `Accept: text/html`.

## Hello job

A sample job implemented in Go and distributed as a public container image. It is used in the [Cloud Run quickstart](https://cloud.google.com/run/docs/quickstarts/jobs/create-execute) and is a suggested container image in the Cloud Run UI on Cloud Console.

* **Container Image:** `us-docker.pkg.dev/cloudrun/container/hello-job`
* **Source Code:** [job/](job/)

## Placeholder service

A sample service implemented in Go and distributed as a public container image. It is used to create a placeholder revision when setting up [Continuous Deployment](https://cloud.google.com/run/docs/continuous-deployment-with-cloud-build).

* **Container Image:** `us-docker.pkg.dev/cloudrun/container/placeholder`

The placeholder defaults to `ACK_EVENTS=false`, replacing POST acknowledgment
statuses (2xx and 102) with HTTP 503 so Pub/Sub and Eventarc deliveries are not
acknowledged. Handlers and GET/HEAD responses remain unchanged. Set `ACK_EVENTS=true`
to restore the previous behavior. Retries remain subject to the event source's
retry policy and retention limits.
