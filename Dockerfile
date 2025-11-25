FROM docker.io/library/busybox:1.37.0@sha256:e3652a00a2fabd16ce889f0aa32c38eec347b997e73bd09e69c962ec7f8732ee AS base

# used to avoid typing the name everywhere
ENV APP_NAME=testrepo

FROM docker.io/library/golang:1.25.4-trixie@sha256:a02d35efc036053fdf0da8c15919276bf777a80cbfda6a35c5e9f087e652adfc AS builder

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
