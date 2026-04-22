FROM docker.io/library/busybox:1.37.0@sha256:e3652a00a2fabd16ce889f0aa32c38eec347b997e73bd09e69c962ec7f8732ee AS base

# used to avoid typing the name everywhere
ENV APP_NAME=testrepo

FROM docker.io/library/golang:1.26.2@sha256:5f3787b7f902c07c7ec4f3aa91a301a3eda8133aa32661a3b3a3a86ab3a68a36 AS builder

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
