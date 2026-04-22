FROM docker.io/library/busybox:1.37.0@sha256:1487d0af5f52b4ba31c7e465126ee2123fe3f2305d638e7827681e7cf6c83d5e AS base

# used to avoid typing the name everywhere
ENV APP_NAME=testrepo

FROM docker.io/library/golang:1.25.7@sha256:85c0ab0b73087fda36bf8692efe2cf67c54a06d7ca3b49c489bbff98c9954d64 AS builder

WORKDIR /app

# get dependencies first
COPY go.mod go.sum ./
RUN go mod download

# only copy what is required
# and give it a predictable name
COPY main.go ./
COPY cmd/ ./cmd
COPY internal/ ./internal
RUN go build -o a.out

FROM base AS final

COPY entrypoint /entrypoint
RUN chmod +x /entrypoint

COPY --from=builder --chown=0:0 /app/a.out /usr/bin/$APP_NAME

USER nobody:nobody
ENTRYPOINT ["/entrypoint"]
