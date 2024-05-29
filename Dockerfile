FROM golang:1.21-alpine

RUN apk add --no-cache bash
WORKDIR /home/bringauto/virtual-fleet-management
COPY . /home/bringauto/virtual-fleet-management/tmp
RUN #chmod +x ./tmp/scripts/docker_build.sh
RUN bash ./tmp/build.sh
RUN mv ./tmp/virtual-fleet-management ./
RUN rm -r ./tmp
