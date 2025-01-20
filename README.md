# Todo API

A simple Todo REST API built using [Gin Web Framework](https://github.com/gin-gonic/gin) for managing to-do lists. This API supports features like creating, retrieving, updating, completing, and deleting todo items. It also includes user authentication and authorization using JWT tokens.

---

## Features

- User authentication and authorization (signup/login).
- Create, retrieve, update, and delete todos.
- Mark todos as complete.
- Middleware for validating JWT tokens and authorizing users.
- Database used was SQLite
- Server is running on port 8080

---

## Endpoints

### Public Endpoints

| Method | Endpoint       | Description                |
|--------|----------------|----------------------------|
| POST   | `/signup`      | Register a new user        |
| POST   | `/login`       | Authenticate a user        |
| GET    | `/todolist`    | Get all todos for a user   |

### Protected Endpoints

| Method | Endpoint                | Description                   |
|--------|-------------------------|-------------------------------|
| GET    | `/todolist/:id`         | Get a single todo by ID       |
| POST   | `/todolist`             | Create a new todo             |
| PUT    | `/todolist/:id`         | Update a todo by ID           |
| PUT    | `/todolist/:id/complete`| Mark a todo as complete       |
| DELETE | `/todolist/:id`         | Delete a todo by ID           |

---

## Middleware

- **Authentication Middleware:**
  Ensures that all protected routes can only be accessed by authenticated users. Unauthorized requests are rejected with a `401 Unauthorized` response.

---
