# syntax = docker/dockerfile:1.2

FROM golang:1.18-alpine
ENV CGO_ENABLED=0

# # define build args and environment variables
ARG PORT
ENV PORT $PORT

ENV VERSION 1.0.0

# mount env file - Render cloud stores .env secrets in this location hence the need to mount unto the docker image
RUN --mount=type=secret,id=_env,dst=/etc/secrets/.env cat /etc/secrets/.env

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o bin/app -ldflags "-X main.Version=$VERSION" app/main.go

EXPOSE $PORT

CMD [ "bin/app" ]