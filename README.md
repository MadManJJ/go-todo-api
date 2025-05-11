# Go Todo API

A simple RESTful API for managing Todo tasks, built with **Go** and the **Gin** framework. The API supports CRUD operations and user authentication via JWT (JSON Web Token).

This project also includes a **Docker** setup for running PostgreSQL as the database.

## Features

- **Create, Read, Update, Delete (CRUD)** operations for Todo tasks.
- **User Authentication** using JWT.
- **Database** interaction via **PostgreSQL** (managed through Docker).
- Pagination support for fetching Todos.
- Secure API endpoints with JWT middleware.

## Technologies

- **Go** (Golang)
- **Gin** (Web Framework)
- **PostgreSQL** (Database)
- **JWT** (User Authentication)
- **Docker** (For PostgreSQL)
- **GORM** (ORM for Go)

## Database Configuration
```DB_HOST=localhost
DB_PORT=5432
DB_NAME=todos
DB_USER=postgres
DB_PASSWORD=password

# PGAdmin Configuration (for PostgreSQL management)
PGADMIN_DEFAULT_EMAIL=admin@admin.com
PGADMIN_DEFAULT_PASSWORD=admin

# JWT Secret Key for Authentication
JWT_SECRET_KEY=your_secret_key_here
```
