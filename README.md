# Go Web Server Tutorial

An educational web application that teaches users how to build web servers in Go through interactive tutorials and examples.

## Features

- **Comprehensive Tutorials**: Learn from basic to advanced Go web server concepts
- **Interactive Examples**: Explore working code examples with syntax highlighting
- **RESTful API Development**: Understand how to design and implement RESTful APIs
- **Downloadable Code**: Get ready-to-use code examples for your own projects
- **Progressive Learning Path**: Follow a structured learning path from fundamentals to advanced topics

## Tutorial Topics

1. **Basic Concepts**
   - Hello World Web Server
   - Serving HTML Pages
   - Handling Different URL Routes

2. **Intermediate Concepts**
   - Using HTML Templates
   - Handling HTML Forms
   - Using Middleware

3. **Advanced Concepts**
   - Building JSON APIs
   - Using Context
   - Graceful Shutdown

4. **RESTful API Development**
   - RESTful API Basics
   - API Versioning Strategies
   - API Documentation

## Getting Started

### Prerequisites

- Go 1.19 or higher

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/go-web-server-tutorial.git
   cd go-web-server-tutorial
   ```

2. Initialize Go modules:
   ```bash
   go mod tidy
   ```

3. Run the server:
   ```bash
   go run main.go
   ```

4. Open your browser and navigate to:
   ```
   http://localhost:5000
   ```

## Project Structure

```
├── content/            # Tutorial and example content
├── handlers/           # HTTP handlers and request processing
├── static/             # Static assets (CSS, JS, images)
│   ├── css/
│   ├── js/
│   └── examples/       # Generated code examples
├── templates/          # HTML templates
├── .github/            # GitHub Actions workflows
├── main.go             # Application entry point
└── go.mod              # Go module definition
```

## Development

To add new tutorials or examples:

1. Add tutorial content to the appropriate function in `content/tutorials.go`
2. Add example code to `content/examples.go`
3. The server will automatically generate the example files in the `static/examples` directory

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Go standard library documentation
- The Go community for their excellent resources