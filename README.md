# FILM INFO
FILM INFO is the spider of the douban, you can easy to get film info in this package.

# START

```
mv confg.local config.yaml // and set you own config

// then you can run:
go run film.go douban

// or run:
go build && /go/bin/film-info douban

// if you want run with docker, you can build like:
docker build -t youHubName/film-info -f Dockerfile .

// and push or run:
docker push youHubName/film-info && docker run -d youHubName/film-info

```