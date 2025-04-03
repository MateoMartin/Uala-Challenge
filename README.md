# How to Run the Uala Challenge App

## Prerequisites
1. **Go**: Ensure that Go is installed on your machine. You can download it from [here](https://golang.org/dl/).


## Running the Service
1. **Navigate to the project directory**: Open your terminal and navigate to the root directory of the project.

2. **Run the service**: Use the following command to start the service:
    ```sh
    go run cmd/main.go
    ```
   This command compiles and runs the Go program starting from the `main.go` file located in the `cmd` directory.

   You can also use docker-compose to run the project with:
     ```sh
    docker-compose up --build
    ```

3. **See Swagger**: You can see the Swagger documentation using the Swagger UI.

Swagger: http://localhost:8080/swagger/index.html

## Running Tests
To run all the tests in the project, use the following command:

1. **Navigate to the root directory** of the project.

2. **Run the tests**:
    ```sh
    go test ./...
    ```
   This command runs all tests in the project and reports the results.

## Endpoint Documentation

You can see the endpoints and parameters [here](./Endpoints.md)

## Architecture Decisions

You can see the Architecture decisions [here](./Architecture.md)

## Summary
- **Start the service**: `go run cmd/main.go`
- **Run all tests**: `go test ./...`
- **Endpoint Documentation**: [Endpoints](./Endpoints.md)
- **Architecture Documentation**: [Architecture](./Architecture.md)




