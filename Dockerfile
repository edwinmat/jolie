# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:latest

#Cache dependencies
RUN go get "github.com/streadway/amqp"
RUN go get "github.com/gorilla/mux"
RUN go get "github.com/edwinmat/server-shared"
RUN go get "gopkg.in/mgo.v2"
RUN go get "gopkg.in/mgo.v2/bson"
RUN go get "github.com/influxdb/influxdb/client"

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/edwinmat/jolie
WORKDIR /go/src/github.com/edwinmat/jolie

RUN go build .

CMD ["./jolie"]
