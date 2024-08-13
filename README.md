A simple Open API project written in Go.

![screenshot](https://github.com/thanhhaudev/openapi-go/blob/master/docs/screenshot.png?raw=true)

## Features
- Simple API endpoints
- Basic CRUD operations

## Prerequisites
Before you begin, ensure you have met the following requirements:

| Requirement        | Description                                                                                                                  |
|--------------------|------------------------------------------------------------------------------------------------------------------------------|
| **Go**             | Install Go from the [official website](https://golang.org/dl/).                                                              |
| **Docker**         | Install Docker from the [official website](https://www.docker.com/get-started).                                              |
| **Docker Compose** | Install Docker Compose from the [official website](https://docs.docker.com/compose/install/).                                |
| **golang-migrate** | Install golang-migrate from the [official documentation](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate). |
| **air**            | Install `air` command line tool by running `go install github.com/air-verse/air@latest`.                                     |
| **swag**           | Install `swag` command line tool by running `go install github.com/swaggo/swag/cmd/swag@latest`.                             |

## Usage

1. **Build the application:**
    ```sh
    make build
    ```
2. **Run the application:**
    ```sh
    make up
    ```
   Seed the database:
    ```sh
    make migrate/up
    ```
3. **Access logs:**
    ```sh
    make logs
    ```
4. **Restart the application:**
    ```sh
    make restart
    ```
5. **Stop the application:**
    ```sh
    make down
    ```

## Access the API documentation:
   - Open your browser and go to `http://localhost:8080/swagger/index.html`

## Authentication sample data

Here are some sample tenants with their respective API keys and secrets for authentication:

| Tenant Name    | API Key              | API Secret           | Scopes                      |
|----------------|----------------------|----------------------|-----------------------------|
| Default Tenant | KRY2oikKQ4DEgG5VOC57 | CJxNmBP07PfH1GYZqu1O | MANAGE_USER, MANAGE_MESSAGE |
| Tenant 1       | 6yDd4PnFH9MMIdGgOOkf | NWHZUUiqTqbIBGMfcLyS | MANAGE_USER                 |
| Tenant 2       | 4b7Ph2hsJP4ohC0tlw5J | 2UF9c2jvKsUfamAeISli | NONE                        |

## Contributing
1. Fork the repository
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request
