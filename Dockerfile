FROM alpine
MAINTAINER Thinh Tran <duythinht@gmail.com>
ADD ./dist /dist
EXPOSE 2435/udp
ENTRYPOINT ["/dist/chaika"]
