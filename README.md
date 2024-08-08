# mini-cms-api

REST API Application for managing contents. Written in Go with Echo Framework

## Tech Stack

- Programming language: [Go](https://go.dev/)
- Database: [MySQL](https://www.mysql.com/)
- Web framework: [Echo](https://echo.labstack.com/)
- ORM library: [GORM](https://gorm.io/)
- Request validation: [Validator](https://github.com/go-playground/validator)
- Application configuration: [Viper](https://github.com/spf13/viper)

## How to Use

1. Clone this repository.

2. Copy the `.env` file for database configuration. Then, configure the database connection inside that file.

```sh
cp .env.example .env
```

3. Create a new database.

```sql
CREATE DATABASE minicmsdb;
```

4. Run the application.

```sh
go run main.go
```
