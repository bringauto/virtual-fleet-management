FROM golang:1.21-alpine
WORKDIR /home/bringauto/virtual-fleet-management

RUN mv ./tmp/virtual-fleet-app ./
RUN rm -r ./tmp
#port without ssl
EXPOSE 1883
#port with ssl
EXPOSE 8883