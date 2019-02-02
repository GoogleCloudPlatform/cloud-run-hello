# "Hello Cloud Run" container for quickstart

This repository contains the source code of the container image used for the Cloud Run quickstart and as a demo container in the Cloud Run UI.

A [Cloud Build trigger](https://pantheon.corp.google.com/cloud-build/triggers/93a2106a-cab6-4056-ad15-021dc5f6c1f0?project=cloudrun
has been set up to automatically build this code into the `gcr.io/cloudrun/hello:latest` image anytime the master branch of this repository is updated.

