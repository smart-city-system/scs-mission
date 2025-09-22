# SCS Mission Service

A comprehensive security incident management system built with Go, providing mission assignment, guidance procedures, and incident documentation capabilities.

## 🏗️ Architecture

The SCS Mission Service is built using clean architecture principles with the following structure:

- **Echo Framework**: High-performance HTTP web framework
- **PostgreSQL**: Primary database for data persistence
- **MinIO**: Object storage for media files
- **JWT Authentication**: Secure API access
- **Swagger Documentation**: Interactive API documentation
- **Docker**: Containerized deployment

## 📋 Features

- **Mission Management**: Assign and track security incident missions
- **Guidance Procedures**: Step-by-step incident response workflows
- **Media Upload**: Support for images and videos (up to 10MB)
- **User Authentication**: JWT-based secure access
- **Real-time Tracking**: Monitor mission progress and completion
- **Comprehensive Logging**: Structured logging with Zap
- **API Documentation**: Interactive Swagger UI

## 🚀 Quick Start

### Prerequisites

- Go 1.23.3 or higher
- PostgreSQL 12+
- MinIO server
- Docker (optional)

### Environment Variables

Create a `.env` file in the root directory:

```env
# Server Configuration
PORT=8080
MODE=development
READ_TIMEOUT=30
WRITE_TIMEOUT=30

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=scs_mission

# MinIO Configuration
MINIO_ENDPOINT=localhost:9000
MINIO_ACCESS_KEY=your_access_key
MINIO_SECRET_KEY=your_secret_key
MINIO_BUCKET_NAME=scs-mission-files

# Logging Configuration
LOG_DEVELOPMENT=true
LOG_DISABLE_CALLER=false
LOG_DISABLE_STACKTRACE=false
LOG_ENCODING=console
LOG_LEVEL=debug
```

### Installation & Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd scs-mission
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up PostgreSQL database**
   ```sql
   CREATE DATABASE scs_mission;
   CREATE USER your_db_user WITH PASSWORD 'your_db_password';
   GRANT ALL PRIVILEGES ON DATABASE scs_mission TO your_db_user;
   ```

4. **Set up MinIO**
   ```bash
   # Using Docker
   docker run -p 9000:9000 -p 9001:9001 \
     -e "MINIO_ROOT_USER=your_access_key" \
     -e "MINIO_ROOT_PASSWORD=your_secret_key" \
     minio/minio server /data --console-address ":9001"
   ```

5. **Run the application**
   ```bash
   go run cmd/server/main.go
   ```

The server will start on `http://localhost:8080`

## 📚 API Documentation

### Swagger UI

Access the interactive API documentation at:
```
http://localhost:8080/swagger/index.html
```

### Authentication

The API uses Bearer token authentication. Include the JWT token in the Authorization header:

```bash
Authorization: Bearer <your-jwt-token>
```

### Main Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/api/v1/health` | Health check | No |
| GET | `/api/v1/missions/me` | Get user assignments | Yes |
| PATCH | `/api/v1/missions/complete` | Complete mission step | Yes |
| PUT | `/api/v1/missions/update` | Upload incident media | Yes |

### Example API Calls

**Get User Assignments:**
```bash
curl -X GET "http://localhost:8080/api/v1/missions/me" \
  -H "Authorization: Bearer <your-token>"
```

**Complete Mission Step:**
```bash
curl -X PATCH "http://localhost:8080/api/v1/missions/complete" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-token>" \
  -d '{
    "mission_id": "550e8400-e29b-41d4-a716-446655440000",
    "step_id": "550e8400-e29b-41d4-a716-446655440001"
  }'
```

**Upload Media Files:**
```bash
curl -X PUT "http://localhost:8080/api/v1/missions/update" \
  -H "Authorization: Bearer <your-token>" \
  -F "incident_id=550e8400-e29b-41d4-a716-446655440000" \
  -F "files=@image.jpg"
```

## 🐳 Docker Deployment

### Build Docker Image

```bash
docker build -f docker/service.Dockerfile -t scs-mission:latest .
```

### Run with Docker Compose

Create a `docker-compose.yml`:

```yaml
version: '3.8'
services:
  app:
    build:
      context: .
      dockerfile: docker/service.Dockerfile
    ports:
      - "8080:8080"
    environment:
      - PORT=8080
      - DB_HOST=postgres
      - DB_PORT=5432
      - DB_USER=scs_user
      - DB_PASSWORD=scs_password
      - DB_NAME=scs_mission
      - MINIO_ENDPOINT=minio:9000
      - MINIO_ACCESS_KEY=minioadmin
      - MINIO_SECRET_KEY=minioadmin
      - MINIO_BUCKET_NAME=scs-mission-files
    depends_on:
      - postgres
      - minio

  postgres:
    image: postgres:15
    environment:
      - POSTGRES_DB=scs_mission
      - POSTGRES_USER=scs_user
      - POSTGRES_PASSWORD=scs_password
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

  minio:
    image: minio/minio
    ports:
      - "9000:9000"
      - "9001:9001"
    environment:
      - MINIO_ROOT_USER=minioadmin
      - MINIO_ROOT_PASSWORD=minioadmin
    command: server /data --console-address ":9001"
    volumes:
      - minio_data:/data

volumes:
  postgres_data:
  minio_data:
```

Run with:
```bash
docker-compose up -d
```

## 🔧 Development

### Project Structure

```
scs-mission/
├── cmd/server/          # Application entry point
├── config/              # Configuration management
├── docs/                # Generated Swagger documentation
├── docker/              # Docker configuration
├── internal/
│   ├── container/       # Dependency injection
│   ├── controllers/     # HTTP handlers
│   ├── dto/             # Data transfer objects
│   ├── middlewares/     # HTTP middlewares
│   ├── models/          # Database models
│   ├── repositories/    # Data access layer
│   ├── server/          # Server setup
│   └── services/        # Business logic
└── pkg/                 # Shared packages
    ├── db/              # Database connection
    ├── errors/          # Error handling
    ├── logger/          # Logging utilities
    ├── minio/           # MinIO client
    ├── utils/           # Utility functions
    └── validation/      # Input validation
```

### Regenerate Swagger Documentation

After making changes to API annotations:

```bash
swag init -g cmd/server/main.go -o docs
```

### Running Tests

```bash
go test ./...
```

### Code Quality

```bash
# Format code
go fmt ./...

# Lint code
golangci-lint run

# Vet code
go vet ./...
```

## 📊 Data Models

### Core Entities

- **User**: System users with roles and authentication
- **Premise**: Physical locations and buildings
- **Alarm**: Security alarms and triggers
- **Incident**: Security incidents requiring response
- **GuidanceTemplate**: Reusable response procedures
- **IncidentGuidance**: Assigned guidance for specific incidents
- **IncidentMedia**: Media files attached to incidents

### Response Format

**Success Response:**
```json
{
  "status": 200,
  "code": "0000",
  "data": { ... }
}
```

**Error Response:**
```json
{
  "error": {
    "type": "VALIDATION_ERROR",
    "message": "Validation failed",
    "details": { ... }
  },
  "request_id": "req-123456",
  "timestamp": "2023-01-01T00:00:00Z"
}
```

## 🔒 Security

- JWT-based authentication
- Input validation on all endpoints
- SQL injection protection via GORM
- File type and size validation for uploads
- Structured error handling without sensitive data exposure

## 📝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Update documentation
6. Submit a pull request

## 📄 License

This project is licensed under the Apache 2.0 License - see the LICENSE file for details.

## 🆘 Support

For support and questions:
- Email: support@swagger.io
- Documentation: http://localhost:8080/swagger/index.html
- Issues: Create an issue in the repository
