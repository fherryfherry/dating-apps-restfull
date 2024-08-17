# Apps Dating API Restfull
by Ferry Ariawan

### Tech Stack:
- Language: Golang 1.22
- DB: MySQL

### Dependencies:
- github.com/go-playground/validator v9.31.0+incompatible
- github.com/golang-jwt/jwt/v5 v5.0.0
- github.com/labstack/echo-jwt/v4 v4.2.0
- github.com/labstack/echo/v4 v4.12.0
- github.com/spf13/viper v1.19.0
- github.com/stretchr/testify v1.9.0
- golang.org/x/crypto v0.24.0
- gorm.io/driver/mysql v1.5.7
- gorm.io/gorm v1.25.10

I assume you have GO and mySQL server environment on your computer.

### Installation
1. Import the database named `dating_apps.sql` into your mySQL local server
1. Config the database credential on `config.yaml` at the root directory
1. Update dependencies by run command `go mod tidy`

### Run Project
```bash
go run main.go
```

### Run Unit Test
```bash
go test -v ./...
```

### Postman Collection
Import postman file named `Dating Apps.postman_collection.json`

# Project Structure
## assets
Contains a collection of asset files or file storage that has been uploaded. Which can be accessed by public users.
## commons
Contains a collection of function files that can be used reusably or what we usually know as a kind of Helper function.
## domains
Contains a collection of database tables used in the project
## handlers
Contains a collection of restful API modules and is separated based on folders according to the intended API name.
## config.yaml
Contains the basic configuration used by this project.
## main.go
The main file of this project system. Contains routing and function initiations.