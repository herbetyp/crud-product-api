![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
<!-- ![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white) -->


# Crud Products API

> API for product inventory management.

### Developed with:
- [Golang](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)
- [Gin Web Framwork](https://github.com/gin-gonic/gin)
<!-- - [GORM](https://gorm.io/index.html)
- [Swagger](https://swagger.io/) (documentation)
- [Heroku](https://www.heroku.com/) (cloud platform) -->


### API Features:

> Auth
- [x] OAuth Authentication Endpoint (*_grant_type=client_credentials_*)
- [x] Authorization from **JWT Bearer Token**
> Admin
- [x] Create User
- [ ] Get User
- [ ] List Users
- [ ] Delete User
> User / Admin
- [x] Reset Password
> Products
- [x] Create Product
- [x] Get Product
- [x] List Products
- [x] Update Product
- [x] Delete Product
> Documentation
- [ ] Documentation with **Swagger/OpenAPI**
> Application systems
- [ ] Log System
- [ ] Cache System
- [ ] Auto Migrations System
---

### Run localy application:
Docker *__with docker compose module__*

```bash
docker compose up -d # exposed in port 3000
```
Runner Local Server *__Gin__*
```bash
export GINMODE="local"
docker compose up -d api_db
go run cmd/main.go # exposed in port 5000
```