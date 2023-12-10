

```markdown
# Go File Processing API with JWT Authentication

This Go API provides functionality to process files and extract information such as the number of words, counts, vowels, and punctuation. Additionally, it incorporates JWT token-based authentication and authorization.

## Packages

### 1. `filehandle`

The `filehandle` package is responsible for handling file processing tasks. It includes functionalities to create, read, and analyze text files. 

#### Usage

```go
// Import the package
import "github.com/yourusername/yourproject/filehandle"

// Example: Read content from a file
content, err := filehandle.ReadFile("example.txt")
if err != nil {
    // Handle error
}

// Example: Get file statistics
stats := filehandle.GetFileStatistics(content)
fmt.Printf("Word Count: %d\nCharacter Count: %d\n", stats.WordCount, stats.CharCount)
```

### 2. `login`

The `login` package provides JWT token-based authentication and authorization.

#### Usage



### 3. `utils`

The `utils` package contains utility structs and functions used across the project.



## Getting Started

1. Clone the repository: `git clone https://github.com/AbdulRafayZia/Gorilla_mux.git`
2. Install dependencies: `go mod tidy`
3. Run the application: `go run main.go`

## API Endpoints



### 2. Authentication

- **Endpoint**: `/login`
- **Method**: POST
- **Request Body**: JSON containing the username and password
- **Response**: JWT token if authentication is successful

### 3. Authorization

- **Endpoint**: `/protected`
- **Method**: GET
- **Authentication**: Requires a valid JWT token


