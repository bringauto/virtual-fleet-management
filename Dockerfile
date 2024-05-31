FROM golang:1.21-alpine

RUN apk add --no-cache bash
WORKDIR /home/bringauto/virtual-fleet-management
COPY . /home/bringauto/virtual-fleet-management/tmp
RUN #chmod +x ./tmp/scripts/docker_build.sh
RUN bash ./tmp/build.sh
RUN mv ./tmp/virtual-fleet-management ./
RUN mkdir -p config && cp ./tmp/resources/config/for_docker.json ./config/config.json
RUN rm -r ./tmp
