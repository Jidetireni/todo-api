# Todo API

A simple Todo API built using [Gin Web Framework](https://github.com/gin-gonic/gin) for managing to-do lists. This API supports features like creating, retrieving, updating, completing, and deleting todo items. It also includes user authentication using JWT tokens.

## Table of Contents

- [Features](#features)
- [Technologies](#technologies)
- [Installation](#installation)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Middleware](#middleware)
- [Contributing](#contributing)
- [License](#license)

---

## Features

- User authentication and authorization (signup/login).
- Create, retrieve, update, and delete todos.
- Mark todos as complete.
- Middleware for validating JWT tokens and authorizing users.

---

## Technologies

- **Backend:** [Gin Framework](https://github.com/gin-gonic/gin)
- **Database:** (Assumes a database; details can be customized for your setup)
- **Authentication:** JWT (JSON Web Tokens)

---

## Installation

### Prerequisites
- Go (version 1.18 or later)
- Database (e.g., MySQL, PostgreSQL, or SQLite)
- Git

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/Jidetireni/todo-api.git
   cd todo-api
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Configure environment variables:
   - Create a `.env` file in the root directory.
   - Add the following variables:
     ```env
     DB_HOST=<your-database-host>
     DB_PORT=<your-database-port>
     DB_USER=<your-database-username>
     DB_PASSWORD=<your-database-password>
     DB_NAME=<your-database-name>
     JWT_SECRET=<your-jwt-secret-key>
     ```

4. Run database migrations (if applicable):
   ```bash
   go run migrations/migrate.go
   ```

5. Start the server:
   ```bash
   go run main.go
   ```

---

## Usage

### Running the Server
Once the server is running, it listens on `http://localhost:8080` by default. You can use tools like Postman, Curl, or any HTTP client to interact with the API.

---

## Endpoints

### Public Endpoints

| Method | Endpoint       | Description                |
|--------|----------------|----------------------------|
| POST   | `/signup`      | Register a new user        |
| POST   | `/login`       | Authenticate a user        |

### Protected Endpoints

| Method | Endpoint                | Description                   |
|--------|-------------------------|-------------------------------|
| GET    | `/todolist`             | Get all todos for a user      |
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

## Contributing

1. Fork the repository.
2. Create a new branch for your feature/fix:
   ```bash
   git checkout -b feature-name
   ```
3. Commit your changes and push to your forked repository:
   ```bash
   git commit -m "Add feature-name"
   git push origin feature-name
   ```
4. Open a pull request to the `main` branch of this repository.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Let me know if you'd like additional customization or sections for your `README.md`!
