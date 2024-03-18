# Rest API (Go and Fiber)

## Initial Steps

```bash

# Initializing the
go mod init github/kodesenkoffie/rest-api-go-fiber

###### Adding Project Dependencies ######

# Fiber
go get github.com/gofiber/fiber/v2

# Postgres
go get -u gorm.io/driver/postgres

# UUID
go get github.com/google/uuid

# Bcrypt
go get -u golang.org/x/crypto/bcrypt

# JWT (JSON Web Token)
go get -u github.com/golang-jwt/jwt/v4

# GODOTENV (Reading files from dotenv)
go get github.com/joho/godotenv

# Validator
go get github.com/go-playground/validator/v10
```

# Executing the script

- Development - `pnpm run dev`
