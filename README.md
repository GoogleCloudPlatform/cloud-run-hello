# Cloud Run "Hello" container source code

This repository contains the source code of multiple sample applications.

## Hello service

A sample service implemented in Go and distributed as a public container image. It is used in the [Cloud Run quickstart](https://cloud.google.com/run/docs/quickstarts/deploy-container) and is a suggested container image in the Cloud Run UI on Cloud Console.

* **Container Image:** `us-docker.pkg.dev/cloudrun/container/hello`

### Configuration Options

Set the `COLOR` environment variable to a valid CSS color to change the background color.

<a href="https://deploy.cloud.run"><img src="https://deploy.cloud.run/button.svg" alt="Run on Google Cloud" height="40"/></a>

## Hello job

A sample job implemented in Go and distributed as a public container image. It is used in the [Cloud Run quickstart](https://cloud.google.com/run/docs/quickstarts/jobs/create-execute) and is a suggested container image in the Cloud Run UI on Cloud Console.

* **Container Image:** `us-docker.pkg.dev/cloudrun/container/hello-job`
* **Source Code:** [job/](job/)

## Placeholder service

A sample service implemented in Go and distributed as a public container image. It is used to create a placeholder revision when setting up [Continuous Deployment](https://cloud.google.com/run/docs/continuous-deployment-with-cloud-build).

* **Container Image:** `us-docker.pkg.dev/cloudrun/container/placeholder`
