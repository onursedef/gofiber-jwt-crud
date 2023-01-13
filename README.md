# CRUD API using Gofiber, JWT, GORM, and Swagger

This API is built using Gofiber, JWT, GORM, and Swagger. It allows for full CRUD functionality, including creating, reading, updating, and deleting resources.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go version 1.14 or higher
- PostgreSQL 

### Installation

1. Clone the repository 
```sh
$ git clone https://github.com/YOUR-USERNAME/YOUR-REPOSITORY
```

2. Set up the environment files:
 - Rename .env.example to .env.development and .env.production in the root directory.
 - Fill in the environment variables in the respective files.

3. Install the dependencies
```sh
$ go mod download
```

4. Start the server
```sh
$ go run main.go
```

### File Structure
- `src`: contains all the source code for the application
    - `controllers`: contains the handlers for the routes
    - `database`: contains the db connection and gorm related stuff
    - `middleware`: contains middleware functions
    - `models`: contains the db schema models
    - `repository`: contains the repositories which interact with the database
    - `routes`: contains the routing information
    - `utils`: contains utility functions

### Major models
- `User`: represents a user in the application
- `Post`: represents a post in the application
- `Login` and `Register`: structs used for authentication and registration

### Note
- In the `models/user/User.go` and `models/post/Post.go` files, there are some commented swagger annotations. Please change `gorm.DeletedAt` to `time.Time` and `datatypes.JSON` to `[]string` before running 
    ```sh
    swag init
    ```

### Swagger
Swagger UI is available at `http://127.0.0.1:3001/swagger/index.html#/`


### Documentation Resources

- [Gofiber](https://github.com/gofiber/fiber)
- [Gofiber/swagger](https://github.com/gofiber/swagger)
- [Swaggo/swag](https://github.com/swaggo/swag)
- [Gorm](https://gorm.io/)
- [Gorm/driver/postgresql](https://gorm.io/docs/connecting_to_the_database.html)
- [dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)
- [golang-jwt/jwt/v4](https://github.com/golang-jwt/jwt)
- [crypto/bcrypt](https://golang.org/pkg/crypto/bcrypt/)
- [godotenv](https://github.com/joho/godotenv)
