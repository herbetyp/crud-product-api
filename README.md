![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
<!-- ![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white) -->


# Crud Product API - Golang

>API for product inventory management.

### Developed with:
- [Golang](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)
- [Gin Web Framwork](https://github.com/gin-gonic/gin)
<!-- - [GORM](https://gorm.io/index.html) -->
<!-- - [Swagger](https://swagger.io/) -->


### API Features:
- [x] Create User Enpoint
- [ ] Reset Password Endpoint
- [x] OAuth Authentication Endpoint (*_grant_type=client_credentials_*)
- [x] Authorization from **JWT Bearer Token**
- [x] Create Product
- [x] List All Products
- [x] List One Product
- [x] Update Product
- [x] Delete Product
- [x] Use [Viper](https://github.com/spf13/viper) library for app config and **.env** file for docker-compose
- [ ] Use [GORM](https://gorm.io/index.html) library for manager database
- [ ] Publish API in [Heroku](https://www.heroku.com/) Platform
- [ ] Documentation with [Swagger](https://swagger.io/)

### Run localy application:
[Docker](https://www.docker.com/) **with docker compose module** is required

```bash
docker compose up -d # exposed in port 3000
```
