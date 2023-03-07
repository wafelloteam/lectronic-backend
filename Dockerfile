FROM golang:1.20.1

WORKDIR /app

COPY . .

RUN go mod tidy

ENV DB_HOST=${DB_HOST} \
    DB_PORT=${DB_PORT} \
    DB_NAME=${DB_NAME} \
    DB_USER=${DB_USER} \
    DB_PASS=${DB_PASS}

RUN go build -v -o /app/lectronic-api
EXPOSE 3000

ENTRYPOINT [ "/app/lectronic-api" ]
CMD [ "serve" ]
