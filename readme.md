# User Management System with JWT Token Authentication

This project is a User Management System built using Fiber as the web framework. It incorporates JWT token authentication for securing routes, and utilizes PostgreSQL as the database with Gorm as the ORM. Bcrypt is employed for password hashing, and Validators ensure the integrity of post submissions. The system includes four main routes: Dashboard, Login, Logout, and Signup.

## Table of Contents

1. [Installation](#installation)
2. [Project Structure](#project-structure)
3. [Routes](#routes)
4. [Middleware](#middleware)
5. [Database](#database)
6. [Dependencies](#dependencies)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/Harichandra-Prasath/User-Management-System-in-GO.git
cd User-Management-System-in-GO
```

2. Set up the .env

Populate the Following fields in .env file

1.DB_HOST
2.DB_NAME
3.DB_USER
4.DB_PASSWORD
5.DB_PORT
6.SECRET_KEY         //your secret_key for jwt

3. Run the application

```bash 
make watch
```

Reflex is integrated to watch the changes and restart

## Project Structure

The project is organized with the following directory structure:

- **main.go:** Entry point of the application.
  
- **internal/handlers/:** Contains route handlers for Dashboard, Login, Logout, and Signup.

- **internal/routes/:** Contains the routing Information

- **internal/middleware/:** Custom middleware for JWT authorization.

- **database/:** SQL scripts and database configuration.

- **internal/models/:** Gorm models for database tables.

- **Config/:** Environmental Variables Configration

## Routes

Current available routes are grouped with /api/users

### 1. Dashboard

- **Path:** `/dashboard`
- **Method:** `GET`
- **Authorization:** Requires a valid JWT token.

### 2. Login

- **Path:** `/login`
- **Methods:** `POST` (Login)
- **Authorization:** None

### 3. Signup

- **Path:** `/signup`
- **Method:** `POST`
- **Authorization:** None

### 4. Logout

- **Path:** `/logout`
- **Method:** `GET`
- **Authorization:** Requires a valid JWT token.

## MiddleWare

Custom middleware is implemented for JWT authorization. This middleware validates incoming requests for protected routes and ensures the presence of a valid JWT token.For clients that cant send cookies, sending requests with following header: AUTHORIZATION Bearer $token will work

## Database

PostgreSQL is used as the database, and Gorm serves as the ORM. SQL scripts for setting up the database can be found in the database folder.


## Dependencies

- **[github.com/gofiber/fiber/v2:](https://github.com/gofiber/fiber/v2)** Web framework.

- **[golang.org/x/crypto/bcrypt:](https://pkg.go.dev/golang.org/x/crypto/bcrypt)** Password hashing.

- **[github.com/go-playground/validator/v10:](https://github.com/go-playground/validator)** Input validation.

- **[gorm.io/gorm:](https://gorm.io/gorm)** Object-Relational Mapping for database operations.

For a complete list of dependencies and versions, refer to the `requirements.txt` file or simply run go mod download.

Feel free to explore and contribute to this User Management System!