# EHR Headless API

This is the backend API for a headless Electronic Health Record (EHR) system, built with Go, Fiber, and GORM.

## Local Development Setup

### Prerequisites

*   **Go:** Version 1.21 or later.
*   **PostgreSQL:** A running instance of PostgreSQL.
*   **Docker (Optional):** If you prefer to run PostgreSQL in a Docker container.

### 1. Clone the Repository

```bash
git clone <your-repository-url>
cd ehr-headless-api
```

### 2. Configure Environment Variables

Create a `.env` file in the root of the project by copying the example file:

```bash
cp .env.example .env
```

Now, open the `.env` file and update the variables to match your local PostgreSQL configuration:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_postgres_user
DB_PASSWORD=your_postgres_password
DB_NAME=ehr
JWT_SECRET=a-very-secret-key-for-jwt
ELASTIC_URL=http://localhost:9200
```

### 3. Install Dependencies

Download and install the necessary Go modules:

```bash
go mod tidy
```

### 4. Run the Application

Start the API server:

```bash
go run cmd/api/main.go
```

The server will start on `http://localhost:3000`.

### 5. Running Tests

To run the unit tests for the application, use the following command:

```bash
go test ./...
```
