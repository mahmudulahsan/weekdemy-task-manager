# Bookstore API with GO

## 🛠️ Run/Build Instructions
### 💻 Run locally
- Clone the repository:
  - $ `git clone https://github.com/jahangir1x/weekdemy-task-manager-backend.git`

- CD into the project directory:
  - $ `cd weekdemy-task-manager-backend`

- Create `app.env` file with `app.env.example` as a template:
  - $ `cp app.env.example app.env`

- **Modify `app.env` contents as needed...**

- Run the application:
  - $ `go run main.go`

### 🐳 Run with Docker
- Clone the repository:
  - $ `git clone https://github.com/jahangir1x/weekdemy-task-manager-backend.git`
- CD into the project directory:
  - $ `cd weekdemy-task-manager-backend`
- Create `app.env` file with `app.env.example` as a template:
  - $ `cp app.env.example app.env`
- **Modify `app.env` contents as needed...**
- Run the Docker container:
  - $ `docker compose --env-file app.env up -d`
- Wait for the process to complete...

***
## 🧭 Explore with Postman: [Postman collection link](https://github.com/jahangir1x/weekdemy-task-manager-backend/blob/main/postman_collection/weekdemy-task-manager-backend.postman_collection.json)

***

## 📚 Endpoints Documentation
- JSON standard does not allow comments. The comments are used here for demonstration purposes.
- Ensure to replace `:id` in the paths with the actual ID of the book when making requests.
- Please note that **existing author ID must be provided** when creating and updating a book information.

### 🔵 Signup User
- **Path:** `/signup`
- **Method:** `POST`
- **Authorization required:** `NO`
- **Request Body:**
  ```json
  {
    "username": "username",    // required
    "password": "password",    // required
    "name":     "John Doe",    // required
    "email":    "abc@xyz.com", // required
    "address":  "123 Main St." // optional
  }
  ```
- **Example:**
    - `POST /signup`
      ```json
      {
        "username": "exampleName",
        "password": "examplePass",
        "name":     "Example User",
        "email":    "abc@xyz.com"
      }
      ```

***

### 🔵 Login User
- **Path:** `/login`
- **Method:** `POST`
- **Authorization required:** `NO`
- **Request Body:**
  ```json
  {
    "username": "username", // required
    "password": "password"  // required
  }
  ```
- **Example:**
    - `POST /login`
      ```json
      {
        "username": "exampleName",
        "password": "examplePass"
      }
      ```

***

### 🔵 Create Author
- **Path:** `/bookstore/authors`
- **Method:** `POST`
- **Authorization required:** `YES`
- **Request Body:**
  ```json
  {
    "authorName":  "John Doe",       // required
    "address":     "123 Main St.",   // optional
    "phoneNumber": "+12 34 5678 910" // optional
  }
  ```
- **Example:**
    - `POST /bookstore/authors`
      ```json
      {
        "authorName":  "Example Author",
        "address":     "Example Address",
        "phoneNumber": "+1234567890"
      }
      ```
    - `POST /bookstore/authors`
      ```json
      {
        "authorName": "Example Author"
      }
      ```

***

### 🔵 Get Authors

- **Path:** `/bookstore/authors`
- **Method:** `GET`
- **Authorization required:** `NO`
- **Example:**
    - `GET /bookstore/authors`

***

### 🔵 Get Author by ID
- **Path:** `/bookstore/authors/:id`
- **Method:** `GET`
- **Authorization required:** `NO`
- **Example:**
    - `GET /bookstore/authors/1`
    - `GET /bookstore/authors/123`

***

### 🔵 Delete Author by ID
- **Path:** `/bookstore/authors/:id`
- **Method:** `DELETE`
- **Authorization required:** `YES`
- **Example:**
    - `DELETE /bookstore/authors/1`
    - `DELETE /bookstore/authors/123`

***

### 🔵 Update Author by ID
- **Path:** `/bookstore/authors/:id`
- **Method:** `PUT`
- **Authorization required:** `YES`
- **Request Body:**
  ```json
  {
    "authorName":  "Jane Smith",     // optional
    "address":     "987 Main St.",   // optional
    "phoneNumber": "+10 98 7654 321" // optional
  }
  ```
- **Example:**
    - `PUT /bookstore/authors/1`
      ```json
      {
        "authorName": "New Name",
        "address":    "New Address"
      }
      ```
    - `PUT /bookstore/authors/123`
      ```json
      {
        "authorName": "New Name"
      }
      ```

***

### 🔵 Create Book
- **Path:** `/bookstore/books`
- **Method:** `POST`
- **Authorization required:** `YES`
- **Request Body:**
  ```json
  {
    "bookName":    "Book Name", // required
    "authorID":    99,          // required
    "publication": "Pub Name"   // optional
  }
  ```
- **Example:**
    - `POST /bookstore/books`
      ```json
      {
        "bookName":    "Example Book",
        "authorID":    42,
        "publication": "Example Publication"
      }
      ```
    - `POST /bookstore/books`
      ```json
      {
        "bookName": "Example Book",
        "authorID": 42
      }
      ```

***

### 🔵 Get Books

- **Path:** `/bookstore/books`
- **Method:** `GET`
- **Authorization required:** `NO`
- **Query Parameters:**
    - `bookName`: *(optional)* Filter books by name.
    - `id`: *(optional)* Filter books by ID.
    - `authorID`: *(optional)* Filter books by author ID.
    - `publication`: *(optional)* Filter books by publication.
- **Example:**
    - `GET /bookstore/books`
    - `GET /bookstore/books?bookName=ExampleBook&authorID=42`
    - `GET /bookstore/books?authorID=42`

***

### 🔵 Get Book by ID

- **Path:** `/bookstore/books/:id`
- **Method:** `GET`
- **Authorization required:** `NO`
- **Example:**
    - `GET /bookstore/books/1`
    - `GET /bookstore/books/123`

***

### 🔵 Delete Book by ID

- **Path:** `/bookstore/books/:id`
- **Method:** `DELETE`
- **Authorization required:** `YES`
- **Example:**
    - `DELETE /bookstore/books/1`
    - `DELETE /bookstore/books/123`

***

### 🔵 Update Book by ID

- **Path:** `/bookstore/books/:id`
- **Method:** `PUT`
- **Authorization required:** `YES`
- **Request Body:** *(all fields are optional)*
  ```json
  {
    "bookName":    "New Name", // optional
    "authorID":    99,         // optional
    "publication": "New Pub"   // optional
  }
  ```
- **Example:**
    - `PUT /bookstore/books/1`
      ```json
      {
        "bookName":    "New Name",
        "publication": "New Pub"
      }
      ```
    - `PUT /bookstore/books/123`
      ```json
      {
        "bookName": "New Name"
      }
      ```
***
