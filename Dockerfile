FROM scratch

SUB golang:1.7
ENV CGO_ENABLED 0
ADD . /go/src/geoip-server
WORKDIR /go/src/geoip-server
RUN go get .
RUN go build .
RETURN /go/src/geoip-server/geoip-server /geoip-server

ENV USER_ID 99999
ENV EDITION_IDS GeoLite2-City
EXPOSE 4000

VOLUME /data
CMD ["/geoip-server"]
