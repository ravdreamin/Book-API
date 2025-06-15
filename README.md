
# 📚 Book API in Go (net/http)

A simple RESTful API built with Go's `net/http` package for managing a collection of books. This project demonstrates basic HTTP server setup, routing, JSON encoding/decoding, and handling HTTP requests and responses.

---

## 📦 Features

- 📖 Get all books
- ➕ Add a new book
- 🔍 Fetch a book by ID (via query parameter)
- ✅ Simple in-memory data storage (slice)
- ⚙️ JSON encoding and decoding

---

## 📑 API Endpoints

### Get All Books  
**Request:**  
`GET /books`

**Response:**  
Returns a list of all books in JSON format.

---

### Add a New Book  
**Request:**  
`POST /books`  

**Body (JSON):**
```json
{
  "title": "Atomic Habits",
  "author": "James Clear"
}

Response:
Returns the added book with a generated ID.


---

Get Book by ID

Request:
GET /book?id=1

Response:
Returns the book with the specified ID, or a 404 error if not found.


---

🚀 How to Run

1. Install Go (if not installed): https://go.dev/doc/install


2. Clone the repository:

git clone https://github.com/your-username/book-api.git
cd book-api


3. Run the server:

go run main.go


4. API will be available at http://localhost:8080




---

📥 Testing with Postman

GET /books → fetch all books

POST /books → add a book (use raw JSON in request body)

GET /book?id=1 → fetch book by ID



---

📚 Technologies Used

Go (Golang)

net/http

encoding/json

Standard Library only (no external frameworks)

+--------+

✨ Author

ravdreamin

