# Use an official Golang runtime as a parent image
FROM golang:latest as build

# Set the working directory to /app
WORKDIR .

# copy cert with correct owner
COPY --chown=0:70 ./*.key .
# Copy the current directory contents into the container at /app
COPY . .

# Download and install any required dependencies
RUN go mod download

# Generate orm
RUN go generate ./database/ent

# Build the Go app
RUN go build .

FROM gcr.io/distroless/base-debian12

COPY --from=build /go/viseh-api .
COPY --from=build --chmod=640 /go/*.key .

COPY ./*.crt .
COPY ./.env .
COPY ./static ./static

# Expose port 8080 for incoming traffic
EXPOSE 8080

# Define the command to run the app when the container starts
CMD ["./viseh-api"]