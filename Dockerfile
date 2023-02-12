FROM golang:alpine3.17 AS BuildStage

WORKDIR /src

COPY . .

RUN go get ./...

RUN go build -o main src/main.go

FROM alpine:latest

WORKDIR /

COPY --from=BuildStage /src/main .

EXPOSE 3000

ENV MONGO_URI=
ENV MONGO_DB_NAME=
ENV YOUTUBE_API_KEY=
ENV SEARCH_CATEGORY=
ENV VIDEO_FETCH_INTERVAL=
ENV VIDEO_FETCH_FROM=

ENTRYPOINT ["/main"]