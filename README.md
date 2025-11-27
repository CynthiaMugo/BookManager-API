# BookManager API

A RESTful API built in **Go** using [Gin](https://github.com/gin-gonic/gin) for managing a personal library. Supports adding, listing, updating, and deleting books, with persistent storage in a JSON file.

---

## Features

- **Add books** with title and author
- **List all books**
- **Update book details** by ID
- **Delete books** by ID
- **Persistent storage** using a JSON file (`books.json`)
- Built with **Gin** for fast, simple REST API development
- **Swagger UI documentation** at `/swagger/index.html`

---

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.20+ installed
- Terminal / Command Prompt

### Clone the repository

```bash
git clone git@github.com:CynthiaMugo/BookManager-API.git
cd BookManager-API
```

### Install Dependencies
#### Gin
```
go get github.com/gin-gonic/gin
go mod tidy
```

#### Swagger(Swaggo)
```
go install github.com/swaggo/swag/cmd/swag@latest
go get github.com/swaggo/files
go get github.com/swaggo/gin-swagger
```
Generate Swagger docs
```
swag init
```

### Running the API
```
go run .
```

Your server will start on:

http://localhost:8080

Swagger UI is available at:

 http://localhost:8080/swagger/index.html

## Endpoints

| Method | Endpoint     | Description         |
| ------ | ------------ | ------------------- |
| GET    | `/books`     | List all books      |
| POST   | `/books`     | Add a new book      |
| PUT    | `/books/:id` | Update a book by ID |
| DELETE | `/books/:id` | Delete a book by ID |

### Example Requests

#### Add a book:

POST /books
Body (JSON):
```
{
  "title": "The Midnight Library",
  "author": "Matt Haig"
}
```
#### List books:

GET /books

#### Update a book:

PUT /books/1
Body (JSON):
```

{
  "title": "Updated Title",
  "author": "Updated Author"
}
```

#### Delete a book:

DELETE /books/1

## Project Structure
```
BookManager-API/
│
├── books/               # Core book logic (model + utilities)
│   └── books.go
├── docs/                # Auto-generated Swagger files
├── main.go              # API entry point + routes
├── data.go              # JSON load/save functionality
└── books.json           # Persistent data store

```

## License
This project is open-source and free to use.