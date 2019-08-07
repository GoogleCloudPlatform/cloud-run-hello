# Cloud Run Quickstart

This repository contains the source code of a sample Go application that is
distributed as the public container image (`gcr.io/cloudrun/hello`) used in the
[Cloud Run quickstart](https://cloud.google.com/run/docs/quickstarts/) and as
the suggested container image  in the Cloud Run UI on Cloud Console.

[![Run on Google Cloud](https://storage.googleapis.com/cloudrun/button.svg)](https://console.cloud.google.com/cloudshell/editor?shellonly=true&cloudshell_image=gcr.io/cloudrun/button&cloudshell_git_repo=https://github.com/GoogleCloudPlatform/cloud-run-hello.git)


[Add a Cloud Build trigger](https://cloud.google.com/cloud-build/docs/running-builds/automate-builds) for the github repository, with Cloud Build configuration file location set to "/cloudbuild.yaml". The tigger watches any commit from any branch within the repository, then build & deploy to Cloud Run accordingly. 