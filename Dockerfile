# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang as builder

# Copy local code to the container image.
WORKDIR /go/src/cloudrun/hello
COPY . .

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN CGO_ENABLED=0 GOOS=linux go build -v -o hello

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
# Use Google managed base image
# https://cloud.google.com/container-registry/docs/managed-base-images
FROM marketplace.gcr.io/google/ubuntu1804:latest

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/src/cloudrun/hello/hello /hello

# Copy template
COPY index.html /index.html

# Run the web service on container startup.
CMD ["/hello"]