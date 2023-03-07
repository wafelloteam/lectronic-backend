FROM golang:1.20.1

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -v -o /app/lectronic-api

EXPOSE 3000
ENTRYPOINT [ "/app/lectronic-api"]
CMD [ "serve" ]
