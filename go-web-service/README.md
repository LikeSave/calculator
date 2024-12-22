# Go Web Service

This project is a simple web service that provides an API endpoint for evaluating arithmetic expressions. The service listens for POST requests at the `/api/v1/calculate` endpoint and returns the result of the expression or appropriate error messages based on input validation and processing.

## Project Structure

```
go-web-service
├── src
│   ├── main.go              # Entry point of the application
│   ├── handlers
│   │   └── calculate.go     # Handler for the /api/v1/calculate endpoint
│   ├── models
│   │   └── expression.go     # Data structure for the expression input
│   └── utils
│       └── calculator.go     # Logic for evaluating arithmetic expressions
├── go.mod                   # Module definition for the Go project
└── README.md                # Documentation for the project
```

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

1. Clone the repository:
   ```
   git clone <repository-url>
   cd go-web-service
   ```

2. Navigate to the `src` directory:
   ```
   cd src
   ```

3. Install dependencies:
   ```
   go mod tidy
   ```

### Running the Service

To run the web service, execute the following command in the `src` directory:

```
go run main.go
```

The service will start and listen on `localhost:8080`.

### API Endpoint

#### POST /api/v1/calculate

- **Request Body:**
  ```json
  {
    "expression": "2 + 2"
  }
  ```

- **Response:**
  - On success:
    ```json
    {
      "result": 4
    }
    ```
  - On error:
    ```json
    {
      "error": "error message"
    }
    ```

### Example Request

Using `curl`, you can test the API as follows:

```bash
curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d '{"expression": "3 * (4 + 5)"}'
```

### License

This project is licensed under the MIT License. See the LICENSE file for details.