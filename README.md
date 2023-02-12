# video-search

## What is video-search?
- it expose the REST API endpoints to fetch the meta-data of youtube videos using youtube data v3 api.

## How to setup locally?
- Make sure you have installed at least `go1.16` in your local machine.
- Clone this repository with `git clone`.
- Run `go get ./...` to download all the dependencies used by `video-search`.
- Expose mandatory envs `MONGO_URI`, `MONGO_DB_NAME`, `YOUTUBE_API_KEY` and to specify `VIDEO_FETCH_FROM` to start fetching videos published after time.
- Run `go run src/main.go` to  run the server.

## How to use docker image?
- Make sure you already have installed `docker` in your machine.
- Go to root directory of the repository and run `docker build -t video-search .` to build the docker image
- Run the dockerized image with `docker run -p 3000:3000 video-search:latest`. P.S. you might have to expose the envs as mentioned above.
- Checkout already built docker image on `https://hub.docker.com/r/utsavvaghani/video-search`.