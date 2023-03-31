# Testing for Cloud Run Samples

A Google Cloud Project is required in order to run the tests in the Cloud Run Samples. The project should have the following API's enabled:

* Cloud Run
* Cloud Build
* Container Registry

## Test Project Setup

* Set up billing for your project
* Cloud Build must be given access to deploy Cloud Run services (see [Deploying to Cloud Run][access]).
* Cloud Build GitHub App needs to be installed and connected to the repository. More info can be found in [Installing the Cloud Build app][app].

[access]: https://cloud.google.com/build/docs/deploying-builds/deploy-cloud-run#continuous-iam
[app]: https://cloud.google.com/build/docs/automating-builds/build-repos-from-github#installing_gcb_app

## Cloud Build Triggers

Each sample has Cloud Build triggers:

* A **Pull Request trigger** which checks incoming changes.
* A **Merge trigger** which builds and pushes new container images.
* A **Nightly trigger** which checks the affects of product changes, environment changes, and flakiness.

The trigger configs are defined in `testing/triggers` and can be imported via:

```sh
gcloud beta builds triggers import --source=testing/triggers/jobs.<TYPE>.yaml
```

## Manually Start Cloud Builds

To manually trigger a Cloud Run (fully managed) build via CLI:

```sh
gcloud builds submit \
  --config "testing/$SAMPLE.pr.cloudbuild.yaml" \
  --substitutions "SHORT_SHA=manual"
```

## Manually Provide Pack CLI Cloud Builders

Follow [instructions][pack] for creating your own Cloud Builder with the `pack` CLI.

[pack]: https://github.com/GoogleCloudPlatform/cloud-builders-community/tree/master/pack