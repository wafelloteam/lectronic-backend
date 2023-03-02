FROM golang:1.20.1

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -v -o /app/lectronic-backend

EXPOSE 3000
ENTRYPOINT [ "/app/lectronic"]
CMD [ "serve" ]