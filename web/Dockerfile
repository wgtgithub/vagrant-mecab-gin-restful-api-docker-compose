FROM golang:1.15.0-alpine3.12
LABEL maintainer="wgtdocker"

COPY ./src/ /go/src/
COPY ./runserver.sh /

RUN apk add --update --no-cache build-base
ENV build_deps 'curl git bash file sudo openssh'
ENV dependencies 'openssl'

RUN apk add --update --no-cache ${build_deps} \
  # Install dependencies
  && apk add --update --no-cache ${dependencies}

#ENV CGO_LDFLAGS "-L/usr/local/lib -lmecab -lstdc++"
#ENV CGO_CFLAGS "-I/usr/local/include"
#ENV CGO_LDFLAGS "mecab-config --libs"
#ENV CGO_CFLAGS "mecab-config --inc-dir"

RUN ["chmod", "+x", "/runserver.sh"]

ENTRYPOINT ["/runserver.sh"]
