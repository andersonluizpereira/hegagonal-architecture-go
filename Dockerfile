FROM golang:1.16

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
#ENV CGO_ENABLED=0

RUN go install github.com/spf13/cobra/cobra@latest && \
    go install github.com/golang/mock/mockgen@v1.5.0


#RUN go mod download

RUN apt-get update && apt-get install sqlite3 -y

RUN touch sqlite.db
#CMD go run main.go
CMD ["tail", "-f", "/dev/null"]
