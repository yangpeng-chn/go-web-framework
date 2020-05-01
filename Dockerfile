# Start from golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Peng Yang <yangpeng.chn@gmail.com>"

# Install git, Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

ENV APP_HOME /app

# Set the current working directory inside the container 
WORKDIR $APP_HOME

# Copy go mod and sum files 
COPY go.mod go.sum ./

# RUN go mod tidy

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.9.4

# metadata for better organization
LABEL app="webapi"
LABEL environment="production"

# Set workdir on current image, /app doesn't work. err: standard_init_linux.go:211: exec user process caused "no such file or directory"
ENV APP_HOME /appdir
ENV APP_USER appuser
WORKDIR $APP_HOME

# Leverage a separate non-root user for the application
RUN adduser -S -D -H -h $APP_HOME $APP_USER

# Change to a non-root user
USER $APP_USER

# Add artifact from builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
COPY --from=builder /app/conf.json .

# Expose port to host
EXPOSE 4201

# Run software with any arguments
ENTRYPOINT ["./main"]