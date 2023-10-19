FROM golang:1.20

RUN apt-get update -y
RUN apt-get install -y wait-for-it

COPY . /src
WORKDIR /src
RUN go build -o /srv/aurora-test

EXPOSE 3000
