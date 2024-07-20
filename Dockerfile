FROM golang:1.22.5-alpine3.20 AS build

WORKDIR /app

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
COPY ./internal ./internal
COPY ./pkg ./pkg
COPY ./Taskfile.yaml ./Taskfile.yaml

## Installing dependencies
RUN --mount=type=cache,target=/root/.cache/go-build go mod download 

## Installing task
RUN --mount=type=cache,target=/root/.cache/go-build go install github.com/go-task/task/v3/cmd/task@latest

## Building project
RUN --mount=type=cache,target=/root/.cache/go-build task build


### Production image
FROM scratch AS prod

WORKDIR /app

COPY --from=build /app/build/bin/main .

EXPOSE 3000
ENTRYPOINT ["./main"]
