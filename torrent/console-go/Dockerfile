FROM golang:1.15.5 AS build
WORKDIR /go/src
COPY src .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o console-go .

FROM ubuntu:18.04 as runtime

RUN apt-get update \
 && apt-get install -y default-jre-headless libjna-java python3 python3-pip mediainfo libchromaprint-tools unrar p7zip-full p7zip-rar mkvtoolnix mp4v2-utils gnupg curl file inotify-tools \
 && rm -rvf /var/lib/apt/lists/*

RUN apt-key adv --fetch-keys https://raw.githubusercontent.com/filebot/plugins/master/gpg/maintainer.pub  \
 && echo "deb [arch=all] https://get.filebot.net/deb/ universal main" > /etc/apt/sources.list.d/filebot.list \
 && apt-get update \
 && apt-get install -y --no-install-recommends filebot \
 && rm -rvf /var/lib/apt/lists/*

RUN pip3 install twilio

VOLUME /console-home
VOLUME /mount

COPY bin /usr/bin
COPY --from=build /go/src/console-go /usr/bin

ENV HOME /console-home
ENV LANG C.UTF-8
ENV FILEBOT_OPTS "-Dapplication.deployment=docker -Duser.home=$HOME"

EXPOSE 80

WORKDIR /console-home

ENTRYPOINT ["/usr/bin/console-go"]
