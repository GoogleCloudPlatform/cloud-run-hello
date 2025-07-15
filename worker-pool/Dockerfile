# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.23 AS builder
# Create and change to the app directory.
WORKDIR /app
# Retrieve application dependencies using go modules.
# Allows container builds to reuse downloaded dependencies.
COPY go.* ./
RUN go mod download
# Copy local code to the container image.
COPY . ./
# Build the binary.
# -mod=readonly ensures immutable go.mod and go.sum in container builds.
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o worker

# Build the runtime container image from scratch, copying what is needed from the previous stage.  
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM scratch
# Create and change to the workspace directory.
WORKDIR /workspace
# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/worker /worker
# Run the web service on container startup.
# Configuring container entrypoint as worker 
ENTRYPOINT ["/worker"]
