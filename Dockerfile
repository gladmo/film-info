FROM golang:alpine

WORKDIR /go/src/github.com/gladmo/film-info
COPY . .

RUN go-wrapper install && \
	ln -s /go/src/github.com/gladmo/film-info/conf /go/bin/conf

CMD ["film-info", "douban"] # ["app"]