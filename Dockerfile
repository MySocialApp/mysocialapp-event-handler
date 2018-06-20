FROM golang:1-stretch

# build image : docker build . -t msa/event-handler
# use -e argument to docker run to set CONFIG_FILE env var
# ex: docker run -v $PWD/config:/mnt -e CONFIG_FILE='/mnt/config.yaml' -p 8080:8080 msa/event-handler
# install git
RUN apt-get update && apt-get --assume-yes install bash git fontconfig libfreetype6 libjpeg62-turbo libpng16-16\
    libx11-6 libxcb1 libxext6 libxrender1 xfonts-75dpi xfonts-base xauth
RUN curl -o wkhtmltox.deb -L  https://downloads.wkhtmltopdf.org/0.12/0.12.5/wkhtmltox_0.12.5-1.stretch_amd64.deb && \
    dpkg -i wkhtmltox.deb
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