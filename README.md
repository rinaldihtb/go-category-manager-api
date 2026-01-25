# Category Manager API

A simple REST API for managing product categories built with Go.

## Features

- Get all categories
- Get a specific category by ID
- Create new categories
- Update existing categories
- Delete categories
- Health check endpoint

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Running the Server

```bash
go run main.go
```

The server will start on `http://localhost:8080`

You should see: `Server running di localhost:8080`

## API Endpoints

### Health Check
- **GET** `/health`
  - Returns API status

### Get All Categories
- **GET** `/categories`
  - Returns all categories

### Get Category by ID
- **GET** `/categories/{id}`
  - Returns a specific category

### Create Category
- **POST** `/categories`
  - Request body:
    ```json
    {
      "name": "Category Name",
      "description": "Category Description"
    }
    ```

### Update Category
- **PUT** `/categories/{id}`
  - Request body:
    ```json
    {
      "name": "Updated Name",
      "description": "Updated Description"
    }
    ```

### Delete Category
- **DELETE** `/categories/{id}`
  - Deletes the specified category

## Example Usage

```bash
# Get all categories
curl http://localhost:8080/categories

# Get category by ID
curl http://localhost:8080/categories/1

# Create new category
curl -X POST http://localhost:8080/categories \
  -H "Content-Type: application/json" \
  -d '{"name":"Makanan","description":"Makanan dan minuman"}'

# Update category
curl -X PUT http://localhost:8080/categories/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Buah Segar","description":"Buah-buahan segar"}'

# Delete category
curl -X DELETE http://localhost:8080/categories/1
```
