FROM golang:1.14

RUN go get github.com/oxequa/realize

# mapped volume in docker-compose.yml
ENV APP_HOME /app

# create working dir
RUN mkdir -p $APP_HOME

WORKDIR $APP_HOME

EXPOSE 4201

CMD [ "realize", "start", "--run" ]