FROM golang:1.14

RUN go get github.com/oxequa/realize

EXPOSE 4201

CMD [ "realize", "start", "--run" ]