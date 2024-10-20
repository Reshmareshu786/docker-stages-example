################################
# BASE IMAGE FOR BUILD STAGE
################################

FROM ubuntu AS build

# Install dependencies
RUN apt-get update && apt-get install -y golang-go

# Set environment variable to disable module support
ENV GO111MODULE=off

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 go build -o myapp .

######################################
# FINAL STAGE
######################################

FROM scratch

# Copy the compiled binary from the build stage
COPY --from=build /app/myapp /app/myapp

# Set the entrypoint for the container to run the binary
ENTRYPOINT ["/app/myapp"]


