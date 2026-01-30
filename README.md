# Category API

A simple REST API for managing categories built with Go.

## Course Reference

This project is part of the "Bootcamp Jago Golang - Code With Umam" course on [CodeWithUmam - Course Online](https://docs.kodingworks.io/s/01e57b74-74e6-44df-ac02-7e30a2478528) and [CodeWithUmam - Youtube](https://www.youtube.com/watch?v=HL1JU206V-4).

This repo is submission for Modul 01 Task on Week 01 [CodeWithUmam - Course Task #01](https://docs.kodingworks.io/s/17137b9a-ed7a-4950-ba9e-eb11299531c2).

## How to Access Publicly

You can access the running application publicly at the following URL:

```
https://fendi-modul-01-task.up.railway.app
```

## How to Use Locally

### Prerequisites
- Go 1.x or higher installed on your system

### Running the Application

1. Navigate to the project directory:
```bash
cd /<your-path-project>/modul-01-task
```

2. Run the application:
```bash
go run main.go
```

3. The server will start on `http://localhost:6969`

You should see:
```
Server is up and running
http://localhost:6969
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Health check |
| GET | `/categories` | Get all categories |
| POST | `/categories` | Create a new category |
| GET | `/categories/{id}` | Get a specific category |
| PUT | `/categories/{id}` | Update a category |
| DELETE | `/categories/{id}` | Delete a category |

## API Usage with cURL

### 1. Health Check
Check if the server is running.

```bash
curl -X GET https://fendi-modul-01-task.up.railway.app/
```

**Response:**
```json
{
  "code": 200,
  "status": "OK"
}
```

---

### 2. Get All Categories
Retrieve all categories.

```bash
curl -X GET https://fendi-modul-01-task.up.railway.app/categories
```

**Response:**
```json
{
  "code": 200,
  "message": "OK",
  "data": {
    "categories": [
      {
        "id": "c231e125",
        "name": "Food",
        "description": "Food and beverages"
      }
    ]
  }
}
```

---

### 3. Create a New Category
Add a new category to the system.

```bash
curl -X POST https://fendi-modul-01-task.up.railway.app/categories \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Electronics",
    "description": "Electronic devices and accessories"
  }'
```

**Response:**
```json
{
  "code": 201,
  "status": "Created",
  "data": {
    "category": {
      "id": "a1b2c3d4",
      "name": "Electronics",
      "description": "Electronic devices and accessories"
    }
  }
}
```

---

### 4. Get Category by ID
Retrieve a specific category by its ID.

```bash
curl -X GET https://fendi-modul-01-task.up.railway.app/categories/c231e125
```

**Response:**
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "category": {
      "id": "c231e125",
      "name": "Food",
      "description": "Food and beverages"
    }
  }
}
```

**Error Response (Not Found):**
```
Not Found
```

---

### 5. Update a Category
Update an existing category by its ID.

```bash
curl -X PUT https://fendi-modul-01-task.up.railway.app/categories/c231e125 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Food & Beverages",
    "description": "All types of food and beverages"
  }'
```

**Response:**
```json
{
  "code": 200,
  "status": "OK"
}
```

**Error Response (Not Found):**
```
Not Found
```

---

### 6. Delete a Category
Remove a category from the system.

```bash
curl -X DELETE https://fendi-modul-01-task.up.railway.app/categories/c231e125
```

**Response:**
```json
{
  "code": 200,
  "status": "OK"
}
```

**Error Response (Not Found):**
```
Not Found
```

---

## Data Structure

### Category
```json
{
  "id": "string (auto-generated)",
  "name": "string",
  "description": "string"
}
```

## Notes
- The ID field is automatically generated when creating a new category
- IDs are 8-character hexadecimal strings
- All responses are in JSON format
- The application uses in-memory storage (data is lost when the server restarts)
