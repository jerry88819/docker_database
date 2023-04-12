FROM golang:alpine3.17
WORKDIR /docker_database
ADD . /docker_database
RUN cd /docker_database && go build
ENTRYPOINT ["./docker_database"]