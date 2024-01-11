# Blog app using Go, JWT Authentication with Gin, PostgreSQL, and GORM

This is a simple Go web application that demonstrates user authentication using JSON Web Tokens (JWT) with Gin, PostgreSQL, and GORM.

## Features

- User signup and login using JWT.
- Secured endpoint demonstrating JWT authentication middleware.
- Create, Get, Update, Delete Funcions

## Prerequisites

- Go installed on your machine.
- PostgreSQL database server.

## Setup

1. Clone the repository:

    ```bash
    git clone https://github.com/barkhayot/go-blog.git
    cd go-blog
    ```

2. Install dependencies:

    ```bash
    go mod download
    ```

3. Set up the PostgreSQL database:

    - Create a PostgreSQL database.
    - Update the database configuration in the `.env` file.

4. Run the application:

    ```bash
    go run main.go
    ```

## Endpoints

- **POST /signup:** Create a new user account.

- **POST /login:** Obtain a JWT token by providing valid credentials.

- **GET /posts:** Access a posts endpoint by providing a valid JWT token in the Authorization header.

- **GET /posts/:id:** Access a posts endpoint by passing ID of needed post

- **POST /posts:** Create new post by sending post details in body of request

- **PUT /posts:** Update post by ID and sending updated structure in body request

- **DELETE /posts:** Delete post by ID 

## Configuration

- The application uses environment variables for configuration. Update the `.env` file with your settings.

    ```env
    # Database Configuration
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=your_db_user
    DB_PASSWORD=your_db_password
    DB_NAME=your_db_name

    # JWT Configuration
    SECRET=your_secret_key
    ```

## Dependencies

- [Gin](https://github.com/gin-gonic/gin): HTTP web framework.
- [GORM](https://gorm.io/): ORM library for database interactions.
- [JWT-Go](https://github.com/dgrijalva/jwt-go): JWT implementation for Go.
