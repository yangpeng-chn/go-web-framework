# Start from golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Peng Yang <yangpeng.chn@gmail.com>"

# Install git, Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

ENV APP_HOME /build

RUN mkdir -p $APP_HOME

# Set the current working directory inside the container 
WORKDIR $APP_HOME

# Copy go mod and sum files from host to container
COPY go.mod go.sum ./

# RUN go mod tidy

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download

# only go.mod and go.sum in /build now, why not mount as Dockerfile.dev does?
RUN echo $(ls -1 /build)

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# everything on host will be in /build in container now
RUN echo $(ls -1 /build)

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.9.4

# metadata for better organization
LABEL app="webapi"
LABEL environment="production"

# Set workdir on current image, /app doesn't work if volume - ./:/app  set in compose. err: standard_init_linux.go:211: exec user process caused "no such file or directory"
ENV APP_HOME /app
ENV APP_USER appuser
WORKDIR $APP_HOME

# Leverage a separate non-root user for the application
RUN adduser -S -D -H -h $APP_HOME $APP_USER

# Change to a non-root user
USER $APP_USER

# Add artifact from builder stage
COPY --from=builder /build/main .
COPY --from=builder /build/.env .
COPY --from=builder /build/conf.json .

# .env conf.json main are in /app
RUN echo $(ls -1a /app)

# Expose port to host
EXPOSE 4201

# Run software with any arguments
ENTRYPOINT ["./main"]