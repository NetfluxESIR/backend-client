# Backend client

Backend client is a small client to interact with the backend server. It is used inside the pipeline to send the data to the backend server
related to the pipeline execution.

## Installation

If you want to use the backend client follow the instructions below.

### Requirements

- [Docker](https://docs.docker.com/get-docker/)

### Run with Docker

```bash
docker run ghcr.io/netfluxesir/backend-client:latest [command]
```

## Usage

```bash
backend-client is a client for the backend

Usage:
  backend-client [flags]
  backend-client [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  status      status the backend
  step        notify the backend of a step

Flags:
  -e, --email string      email
  -h, --help              help for backend-client
  -i, --id string         video id
  -p, --password string   password
  -r, --role string       role
  -U, --url string        url
```

### status

```bash
status the backend

Usage:
  backend-client status [flags]

Flags:
  -h, --help            help for status
  -s, --status string   status

Global Flags:
  -e, --email string      email
  -i, --id string         video id
  -p, --password string   password
  -r, --role string       role
  -U, --url string        url
```

### step

```bash
notify the backend of a step

Usage:
  backend-client step [flags]

Flags:
  -c, --current-step string           current step
  -S, --current-step-status string    current step status
  -h, --help                          help for step
  -P, --previous-step string          previous step
  -l, --previous-step-log string      previous step log
  -s, --previous-step-status string   previous step status

Global Flags:
  -e, --email string      email
  -i, --id string         video id
  -p, --password string   password
  -r, --role string       role
  -U, --url string        url
```

## Development

### Requirements

- [Git](https://git-scm.com/downloads)
- [Go](https://golang.org/doc/install)

### Run from source

#### Build the client

```bash
git clone https://github.com/NetfluxESIR/backend-client.git
cd backend-client
go build -o backend-client main.go
```

#### Run the client

```bash
./backend-client [command]
```
