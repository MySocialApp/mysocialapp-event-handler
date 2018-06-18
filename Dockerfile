FROM golang:1-alpine

# build image : docker build . -t msa/event-handler
# use -e argument to docker run to set CONFIG_FILE env var
# ex: sudo docker run -v $PWD/config:/mnt -e CONFIG_FILE='/mnt/config.yaml' -p 8080:8080 msa/event-handler
# install git
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN mkdir /go/src/app
ADD . /go/src/app
WORKDIR /go/src/app
RUN go get
RUN go test
RUN go build -o event-handler .

ENTRYPOINT /go/src/app/event-handler -config=$CONFIG_FILE