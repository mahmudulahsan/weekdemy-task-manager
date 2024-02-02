FROM golang:1.21.6-alpine3.19
ARG PORT
WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE $PORT

CMD ./main
