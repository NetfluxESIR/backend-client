FROM golang:1.20-alpine3.18 AS builder
WORKDIR /go/src/app
ARG ssh_prv_key
RUN apk add --no-cache git openssh-client
RUN mkdir -p ~/.ssh && umask 0077 && echo "${ssh_prv_key}" > ~/.ssh/id_rsa
RUN ssh-keyscan github.com >> ~/.ssh/known_hosts
COPY . .
RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/
RUN export GOPRIVATE=github.com/NetfluxESIR && go get -d -v ./...
RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian11
COPY --from=builder /go/bin/app /
ENTRYPOINT ["/app"]
