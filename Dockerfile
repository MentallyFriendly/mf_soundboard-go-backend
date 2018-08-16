FROM golang:1.9

ARG app_env

ENV APP_ENV $app_env

WORKDIR /go/src/go_apps/go_api_apps/mf_soundboard

ADD . .

RUN go install
RUN go get -u github.com/tools/godep
RUN godep restore
RUN godep go build
RUN go get -u github.com/pilu/fresh

CMD fresh