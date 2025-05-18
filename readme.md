# Go Hexagonal Architecture Template

[⬇️ Download README.md](https://github.com/yourusername/go-hex-template/raw/main/README.md)

[![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat&logo=go)](https://golang.org/doc/go1.20)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://github.com/yourusername/go-hex-template/graphs/commit-activity)

A robust and scalable starter template for Go applications using Hexagonal Architecture (Ports and Adapters), featuring built-in AWS Lambda integration and containerization support.

> **Note**: This project is under active development. Currently, the focus is on implementing comprehensive unit tests and Docker-based integration tests.

## 🏗️ Architecture

This project follows the Hexagonal Architecture (also known as Ports and Adapters) pattern, which allows for:

- Clear separation of concerns
- Domain-driven design
- Improved testability
- Technology independence
- Easier maintenance and evolution

```
┌────────────────────────────────────────────────────────────────┐
│                         Application                            │
│                                                                │
│    ┌─────────────────────────────────────────────────────┐    │
│    │                      Domain                         │    │
│    │                                                     │    │
│    │    ┌─────────────────────────────────────────┐     │    │
│    │    │              Core Domain                │     │    │
│    │    │                                         │     │    │
│    │    │    ┌─────────────┐    ┌─────────────┐   │     │    │
│    │    │    │   Entity    │    │   Service   │   │     │    │
│    │    │    └─────────────┘    └─────────────┘   │     │    │
│    │    └─────────────────────────────────────────┘     │    │
│    │                                                     │    │
│    │    ┌─────────────────┐      ┌─────────────────┐    │    │
│    │    │   Input Ports   │      │  Output Ports   │    │    │
│    │    └─────────────────┘      └─────────────────┘    │    │
│    └─────────────────────────────────────────────────────┘    │
│                                                                │
│    ┌─────────────────────┐      ┌─────────────────────┐       │
│    │    Input Adapters   │      │   Output Adapters   │       │
│    │    (Controllers)    │      │   (Repositories)    │       │
│    └─────────────────────┘      └─────────────────────┘       │
└────────────────────────────────────────────────────────────────┘
```

## 📋 Features

- ✅ Clean Hexagonal Architecture implementation
- ✅ AWS Lambda integration
- ✅ GORM for database interactions
- ✅ REST API with Gin-Gonic
- ✅ Logging with Zap
- ✅ Environment-based configuration
- ✅ Database migrations
- ✅ Docker support
- ✅ Multi-environment setup (local, dev, prod)
- 🔄 Unit testing (in progress)
- 🔄 Integration testing with Docker (in progress)
- 🔄 CICD with Github actions

## 🛠️ Tech Stack

- [Go](https://golang.org/) - Core programming language
- [Gin-Gonic](https://github.com/gin-gonic/gin) - HTTP web framework
- [GORM](https://gorm.io/) - ORM library for database operations
- [Zap](https://github.com/uber-go/zap) - Structured, leveled logging
- [Resty](https://github.com/go-resty/resty) - HTTP and REST client library
- [AWS Lambda](https://aws.amazon.com/lambda/) - Serverless compute service
- [Docker](https://www.docker.com/) - Containerization
- [PostgreSQL](https://www.postgresql.org/) - Primary database

## 📂 Project Structure

```
.
├── cmd                           # Application entry points
│   ├── api                       # Regular API server
│   └── lambda_api                # AWS Lambda entry point
├── deployments                   # Deployment configurations
│   ├── docker                    # Docker configuration files
│   ├── events                    # Lambda event templates
│   ├── lambda                    # AWS Lambda configuration
│   └── sql_scripts               # Database migrations
│       ├── dev                   # Development environment scripts
│       ├── local                 # Local environment scripts
│       ├── prd                   # Production environment scripts
│       └── test                  # Testing environment scripts
├── internal                      # Private application code
│   ├── adapter                   # Adapters implementation
│   │   ├── input                 # Input adapters (controllers, API)
│   │   └── output                # Output adapters (repositories)
│   ├── application               # Application core
│   │   ├── domain                # Domain models and logic
│   │   ├── port                  # Interface definitions (ports)
│   │   └── service               # Service implementation
│   ├── config                    # Configuration
│   └── constants                 # Application constants
└── tool                          # Project tools
```

## 🚀 Getting Started

### Prerequisites

- Go 1.20+
- Docker and Docker Compose
- AWS SAM CLI (for Lambda testing)
- Make

### Installation & Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/renatofagalde/golang-starter-template.git
   cd golang-starter-template
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

### Running the Application

```bash
# Run the application in local environment (uses PostgreSQL in container)
make local

# Run in development environment (connects to AWS database)
make dev

# Run in production environment (connects to AWS database)
make prod

# Start the full local environment (database + migrations + app)
make up-local

# View available commands
make help
```

### AWS Lambda Deployment

Build the Lambda function:
```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -o bootstrap ./cmd/lambda_api
zip -j deployment.zip bootstrap
```

Test locally with SAM:
```bash
sam local invoke --event deployments/events/list_notes.json --template deployments/lambda/template.yaml
```

### Database Connection

To connect to remote database through SSH tunnel:
```bash
ssh -i .ssh/key.pem -f -N -L 5432:xxxxx-1-instance-1.cvwyig4oy864.us-east-1.rds.amazonaws.com:5432 ec2-user@ec2-34-xxx-xxx-xxx.compute-1.amazonaws.com -v
```

Use the `-f` option to run in background.

## 🧪 Testing

Currently implementing:
- Unit tests for all service layers
- Integration tests with Docker containers
- API endpoint tests

Run tests:
```bash
# Run all tests
go test -v ./...

# Run tests with coverage
go test -v -cover ./...
```