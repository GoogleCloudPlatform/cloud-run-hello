# Cloud Run "Hello" container

This repository contains the source code of a sample Go application that is
distributed as the public container image (`us-docker.pkg.dev/cloudrun/container/hello`) used in the
[Cloud Run quickstart](https://cloud.google.com/run/docs/quickstarts/deploy-container) and as
the suggested container image  in the Cloud Run UI on Cloud Console.

It also contains the source code of a placeholder public container
(`us-docker.pkg.dev/cloudrun/container/placeholder`) used to create a placeholder revision when setting up 
[Continuous Deployment](https://cloud.google.com/run/docs/continuous-deployment-with-cloud-build).

Set a `COLOR` environment variable to a valid CSS color to change the background color.

[![Run on Google Cloud](https://deploy.cloud.run/button.svg)](https://deploy.cloud.run)
