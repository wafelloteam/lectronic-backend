# Lectronic

Lectronic is a marketplace for selling electronic products, such as phones, laptops, headset, and other electronic goods. This RESTful API built using Golang, PostgreSQL, GORM, and dependency injection as modularization. Continous deployment with Docker Compose, Github Action, and Docker Hub.

## 🔥 Showcase

- [Web Client Repository](https://github.com/wafellofazztrack/lectronic-frontend)
- [Docker Image](https://hub.docker.com/r/rfauzi/lectronic-api)
- [Postman Docs](https://documenter.getpostman.com/view/25042327/2s93JtQPYk)

## ⚡ Features

- Authentication & Authorization using JWT
- CRUD for all modules
- Manage Product for admin role
- Checkout & Payment for user role
- Get History order
- Database migrate and seed
- Continous deployment with Github Action and Docker Hub

## 💻 Built with

- [Gorilla MUX](https://github.com/gorilla/mux) for handling HTTP requests and responses
- [GORM](https://github.com/go-gorm/gorm) for ORM library
- [JWT](https://github.com/golang-jwt/jwt) for authentication and authorization
- [Postgres](https://github.com/postgres/postgres) for DBMS
- [Sendinblue](https://github.com/sendinblue/APIv3-go-library) for sending emails
- [Docker](https://github.com/docker) for deployment

## 🛠️ Installation Steps

1. Clone the repository

```bash
git clone https://github.com/wafellofazztrack/lectronic-backend
```

2. Install dependencies

```bash
go get -u ./...
```

3. Migrate up

```bash
go run . migrate -u
```

4. Seed up

```bash
go run . seed -u
```

5. Run the app

```bash
go run . serve
```

🌟 You are all set!

## Contributor

- [Rizaldi Fauzi](https://github.com/rfauzi44): as Project Manager
- [Ahmad Saifudin Ardhiansyah](https://github.com/ardhisaif): as Developer
- [Muhammad Angga Ardhinata](https://github.com/AnggaArdhinata): as Developer
