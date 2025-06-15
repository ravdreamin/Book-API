# 📚 Book API in Go

A simple RESTful API built with Go's `net/http` package for managing a list of books. Supports adding, listing, and fetching books by ID using JSON over HTTP.

## 📑 Endpoints

- `GET /books` → List all books  
- `POST /books` → Add a new book (JSON body)  
- `GET /book?id=1` → Get a book by ID  

## 🚀 Run

```bash
go run main.go