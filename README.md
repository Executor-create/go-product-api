# Description
A Go-based API that provides API operations for managing products.
- Technologies: Go, Fiber (web framework), GORM(ORM)

# Features
- Clean separation of concerns (Handler, Service).
- RESTful API with endpoints to interact with products.

# Installation
1 Clone the repository
   ```bash
  git clone git@github.com:Executor-create/go-product-api.git
  cd go-product-api
   ```
2 Install dependencies
   ```bash
   go mod tidy
  ```
3 Set up your database connection:
- Create .env file and configure the database connection string.

4. Run the application:
  ```bash
    go run main.go
  ```
The application will start running on http://localhost:3000 (or the port specified in your environment configuration).

# API Endpoints
### GET /products/:id
Fetch a product by its ID.
- Response: Returns a JSON object with product details or an error message if not found.
### GET /products?name=:name
Fetch a product by its name.
- Response: Returns a JSON object with product details or an error message if not found.
### GET /products/:id/similar
Fetch similar products based on the category of a product.
- Response: Returns a list of similar products in the same category.
