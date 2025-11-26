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

### Install Gin
```
go get github.com/gin-gonic/gin
go mod tidy
```

### Running the API
```
go run .
```
The server will start on port 8080. You can now test endpoints using Postman or curl.

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
├── books/            # Core book logic
│   └── books.go
├── main.go           # API entry point
├── data.go           # JSON save/load functionality
└── books.json        # Persisted library data
```

## License
This project is open-source and free to use.