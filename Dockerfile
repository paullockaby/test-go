FROM docker.io/library/busybox:1.37.0@sha256:d82f458899c9696cb26a7c02d5568f81c8c8223f8661bb2a7988b269c8b9051e AS base

# used to avoid typing the name everywhere
ENV APP_NAME=testrepo

FROM docker.io/library/golang:1.25.3-trixie@sha256:88ba6af1b83398ee93c4753b7f746bd979ce369d759a058c80667c46db6341ff AS builder

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
