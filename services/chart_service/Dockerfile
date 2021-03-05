FROM golang

ADD . /app
WORKDIR /app

RUN go install

ENV HTTP_PORT=8080
EXPOSE 8080

CMD ["go", "run", "main.go"]