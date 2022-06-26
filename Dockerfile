FROM golang:1.17
WORKDIR /go/src
ENV GO111MODULE=on
ENV SIGNINGKEY=ltasgtagjlers
COPY go.mod ./
RUN go mod download
COPY . .
EXPOSE 80
CMD ["go", "run", "main.go", "stats.go"]