FROM golang:1.17
WORKDIR /go/src

ENV GO111MODULE=on
ENV SIGNINGKEY=ltasgtagjlers

COPY ../ .
EXPOSE 80
RUN go get -u github.com/cosmtrek/air && \
    go get -v github.com/rubenv/sql-migrate/...

CMD ["air"]