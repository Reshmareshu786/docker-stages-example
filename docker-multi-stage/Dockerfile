################################
# BASE IMAGE
###############################

FROM ubuntu AS build

# Update package lists and install necessary packages
RUN apt-get update && apt-get install -y golang-go ca-certificates

# Ensure the CA certificates are up to date
RUN update-ca-certificates

# Turn off Go modules
ENV GO111MODULE=off

# Copy the current directory into the Docker image
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 go build -o /app ./main.go


######################################
# MULTI STAGE BUILD
######################################

FROM scratch

# Copy the compiled binary from the build stage
COPY --from=build /app /app

# Set the entrypoint for the container to run the binary
ENTRYPOINT ["/app"]

