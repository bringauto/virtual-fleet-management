FROM golang:1.16-alpine
WORKDIR /virtual-industrial-portal
COPY . /virtual-industrial-portal/tmp
RUN chmod +x ./tmp/scripts/docker_build.sh
RUN sh ./tmp/scripts/docker_build.sh
RUN mv ./tmp/virtual-industrial-portal-app ./
RUN rm -r ./tmp
#port without ssl
EXPOSE 1883
#port with ssl
EXPOSE 8883