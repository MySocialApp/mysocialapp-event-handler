FROM golang:1-stretch

# build image : docker build . -t msa/event-handler
# use -e argument to docker run to set CONFIG_FILE env var
# ex: sudo docker run -v $PWD/config:/mnt -e CONFIG_FILE='/mnt/config.yaml' -p 8080:8080 msa/event-handler
# install git
RUN apt-get update && apt-get --assume-yes install bash git wkhtmltopdf
RUN apt-get install --assume-yes -qqy x11-apps
ENV DISPLAY :0
CMD xeyes
ENV XSOCK /tmp/.X11-unix
ENV XAUTH /tmp/.docker.xauth
RUN xauth nlist :0 | sed -e 's/^..../ffff/' | xauth -f $XAUTH nmerge -

RUN mkdir /go/src/app
ADD . /go/src/app
WORKDIR /go/src/app
RUN go get
RUN go test
RUN go build -o event-handler .

ENTRYPOINT /go/src/app/event-handler -config=$CONFIG_FILE