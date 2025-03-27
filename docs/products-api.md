### GET /products/{id}

- **Method:** GET
- **URL:** /products/{id}
- **Path Parameter:**
  - `id` (path parameter): The unique identifier of the product.
- **Description:** Retrieve detailed information about a specific product by its ID.
- **Response:**
  ```json
  {
    "id": "1",
    "name": "Slack",
    "url": "https://www.spendbase.co/discounts/slack-50-off/",
    "image_url": "https://www.spendbase.co/wp-content/uploads/2024/05/6911160.png",
    "description": "Slack is a communication platform that enhances team collaboration through chat rooms, private messages, and integrated tools to streamline work processes.",
    "price": "450",
    "rating": 5,
    "summary": "Slack is used by teams to communicate and organize work, helping to reduce email volume and increase productivity.",
    "created_at": "2025-05-01T10:00:00Z",
    "updated_at": "2025-05-10T15:00:00Z",
    "categories": [
      {
        "id": "1",
        "name": "Work Collaboration",
        "url": "https://www.spendbase.co/categories/work-collaboration",
        "created_at": "2025-05-01T10:00:00Z",
        "updated_at": "2025-05-10T15:00:00Z"
      }
    ]
  }
  ```

### GET /products?name={name}

- **Method:** GET
- **URL:** /products?name={name}
- **Query Parameters:**
  - `name` (query parameter): The name of the product to search for.
- **Description:** Retrieve detailed information about a specific product by its name.
- **Response Example:**
  ```json
  {
    "id": "1",
    "name": "Slack",
    "url": "https://www.spendbase.co/discounts/slack-50-off/",
    "image_url": "https://www.spendbase.co/wp-content/uploads/2024/05/6911160.png",
    "description": "Slack is a communication platform that enhances team collaboration through chat rooms, private messages, and integrated tools to streamline work processes.",
    "price": "450",
    "rating": 5,
    "summary": "Slack is used by teams to communicate and organize work, helping to reduce email volume and increase productivity.",
    "created_at": "2025-05-01T10:00:00Z",
    "updated_at": "2025-05-10T15:00:00Z",
    "categories": [
      {
        "id": "1",
        "name": "Work Collaboration",
        "url": "https://www.spendbase.co/categories/work-collaboration",
        "created_at": "2025-05-01T10:00:00Z",
        "updated_at": "2025-05-10T15:00:00Z"
      }
    ]
  }
  ```

### GET /products/{id}/similar

- **Method:** GET
- **URL:** /products/{id}/similar
- **Path Parameter:**
  - `id` (path parameter): The unique identifier of the product.
- **Description:** Retrieve up to 10 similar products from the same category as the given product.
- **Response Example:**
  ```json
  {
    "product": {
      "id": "1",
      "name": "Trello",
      "url": "https://www.spendbase.co/discounts/trello-50-off/",
      "image_url": "https://www.spendbase.co/wp-content/uploads/2024/05/trello.png",
      "description": "Trello is a visual collaboration tool that helps teams organize tasks and projects with boards, lists, and cards.",
      "price": "300",
      "rating": 4.8,
      "summary": "Trello allows teams to manage workflows effectively using a simple drag-and-drop interface.",
      "created_at": "2025-05-01T10:00:00Z",
      "updated_at": "2025-05-10T15:00:00Z",
      "categories": [
        {
          "id": "1",
          "name": "Work Collaboration",
          "url": "https://www.spendbase.co/categories/work-collaboration",
          "created_at": "2025-05-01T10:00:00Z",
          "updated_at": "2025-05-10T15:00:00Z"
        }
      ]
    },
    "similar_products": [
      {
        "id": "2",
        "name": "Asana",
        "url": "https://www.spendbase.co/discounts/asana-50-off/",
        "image_url": "https://www.spendbase.co/wp-content/uploads/2024/05/asana.png",
        "price": "350",
        "rating": 4.7
      },
      {
        "id": "3",
        "name": "Monday.com",
        "url": "https://www.spendbase.co/discounts/monday-50-off/",
        "image_url": "https://www.spendbase.co/wp-content/uploads/2024/05/monday.png",
        "price": "400",
        "rating": 4.6
      }
    ]
  }
  ```
