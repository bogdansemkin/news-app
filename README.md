# News App

News App - это приложение для управления постами, написанное на Go. Приложение предоставляет REST API для выполнения операций CRUD (создание, чтение, обновление, удаление) над постами.

## Структура проекта

```plaintext
.
|-- Makefile
|-- cmd
|   `-- main.go
|-- gen
|-- go.mod
|-- go.sum
|-- internal
|   |-- api
|   |   |-- api.go
|   |   |-- common.go
|   |   |-- post_create.go
|   |   |-- post_delete.go
|   |   |-- post_list.go
|   |   |-- post_update.go
|   |   |-- post_via_id.go
|   |   `-- routes.go
|   |-- config
|   |   `-- config.go
|   |-- repository
|   |   |-- db
|   |   |   `-- postgres
|   |   |       `-- postgres.go
|   |   `-- post
|   |       |-- post.go
|   |       `-- repository.go
|   `-- service
|       |-- post.go
|       `-- service.go
|-- myapp
|-- openapi.yaml
|-- pkg
|   `-- model
|       `-- model.go
`-- schema
    |-- 000001_init.down.sql
    `-- 000001_init.up.sql
```

## Installation and Running

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/news-app.git
    cd news-app
    ```

2. Install dependencies:

    ```sh
    go mod download
    ```

3. Set up the database by applying migrations:

    ```sh
    make migrate-up
    ```

4. Run the application:

    ```sh
    go run cmd/main.go
    ```

## API Documentation

The `openapi.yaml` file contains the description of all available endpoints. You can use it with any OpenAPI tool, such as Swagger.

### Endpoint Examples:

- Get all posts:
## Get All Posts

```http
GET /posts

curl --location --request GET 'http://localhost:8092/posts'
```

## Get a Post by ID

```http
GET /posts/{id}

curl --location --globoff --request GET 'http://localhost:8092/posts/{id}
```

## Create a New Post

```http
POST /posts

curl --location 'http://localhost:8092/posts' \
--header 'Content-Type: application/json' \
--data '{
    "id":7, 
    "title":"ttttt",
    "content":"t"
}'
```

## Update a Post

```http
PUT /posts

curl --location --request PUT 'http://localhost:8092/posts' \
--header 'Content-Type: application/json' \
--data '{
    "id":7, 
    "title":"ttttt",
    "content":"tet"
}'
```

## Delete a Post

```http
DELETE /posts/{id}

curl --location --globoff --request DELETE 'http://localhost:8092/posts/{id}' 
```