# 💼 Project Challenge: Team & Task Management API (Go)

A realistic backend project designed to simulate a production-level REST API system — ideal for showcasing in job applications.

---

## 🧩 Objective

Build a scalable, secure **Task & Team Management Web API** in Go, supporting user accounts, projects, roles, tasks, and authorization. Think of it as a simplified version of **Trello** or **Jira**.

---

## ⚙️ Tech Stack

| Layer             | Technology                         |
|------------------|-------------------------------------|
| Language         | Go (Golang)                         |
| Web Framework    | Gin or Echo                         |
| Database         | PostgreSQL                          |
| ORM              | GORM                                |
| Auth             | JWT (access + refresh tokens)       |
| Migrations       | golang-migrate                      |
| Validation       | go-playground/validator             |
| Docker           | Docker Compose                      |
| Docs             | Swagger (swaggo)                    |
| Testing          | testify or gomock                   |

---

## 🧑‍💼 Core Features

### 🔐 Authentication
- Register, login
- Password hashing (bcrypt)
- JWT access and refresh token system
- Logout & token invalidation

### 👥 User & Role Management
- User CRUD
- Project-specific roles: `Admin`, `Manager`, `Member`
- Invite users to projects (via ID or email)

### 📁 Projects
- Create, update, delete projects
- Assign users to a project
- Role-based access per project

### ✅ Tasks
- Create/update/delete tasks per project
- Task attributes: `title`, `description`, `status`, `assignee`, `dueDate`
- Task status enum: `TODO`, `IN_PROGRESS`, `DONE`
- Assign tasks to users

### 🔍 Filtering & Search
- Filter tasks by status, project, assignee, due date

### 📊 Dashboard Endpoint
- Aggregated data: tasks per project, user activity
- Recent task changes per user/project

---

## 🧰 Dev Features

- Environment-based configuration
- Swagger UI API docs
- Middleware for logging, validation, auth
- Modular folder structure (`internal/user`, `internal/task`, etc.)
- Docker + PostgreSQL setup
- Seed/demo data script
- Role-based middleware (`adminOnly`, `projectMemberOnly`, etc.)

---

## 🧪 Testing

- Unit tests for service layer
- Integration tests for handlers
- Mock database using `testify/mock` or `sqlmock`

---

## 🧠 Bonus (Stretch Goals)

- WebSocket for live task updates
- gRPC internal service (e.g. notification)
- GitHub OAuth login
- Rate limiting middleware
- Email invite system
- CLI tool for admin user setup
- CI/CD with GitHub Actions

---

## 🗂️ Suggested Folder Structure

