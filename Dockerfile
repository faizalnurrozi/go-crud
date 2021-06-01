FROM golang:alpine
#RUN apt-get update && apt-get install -y
ENV GO111MODULE=on
ENV PKG_NAME=github.com/faizalnurrozi/go-crud/
ENV PKG_PATH=$GOPATH/src/$PKG_NAME
RUN apk update && apk upgrade
RUN apk add --no-cache git
RUN git config --global url."https://it.shoesmart:47Pax8bptr7jN7Zeiny5@gitlab.com".insteadOf "https://gitlab.com"
WORKDIR $PKG_PATH/
COPY . $PKG_PATH/
RUN echo $PWD
RUN go mod vendor
WORKDIR $PKG_PATH/server/http
RUN echo $PWD
RUN go build main.go
#EXPOSE 80
EXPOSE 4000
CMD ["sh", "-c", "./main"]