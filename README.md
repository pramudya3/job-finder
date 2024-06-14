# Job and Company Management API

This is a Golang-based backend API for managing companies and jobs, designed with the Gin framework and SQLx for database interactions. The API includes features for registering companies, finding companies by ID, creating jobs, and searching for jobs.

## Features

- Register new companies
- Find companies by ID
- Delete companies by ID
- Create new jobs
- Search for jobs with pagination and filtering

## API Endpoints

### Company Endpoints

- **POST /api/v1/companies**: Register a new company
- **GET /api/v1/companies/:id**: Find a company by ID
- **DELETE /api/v1/companies/:id**: Delete a company by ID

### Job Endpoints

- **POST /api/v1/jobs**: Create a new job
- **GET /api/v1/jobs**: Find jobs with optional keyword and company name filters

### Swagger

- **GET /swagger/index.html**: Access the Swagger UI for API documentation

### Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/pramudya3/job-finder
    cd job-finder
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```
3. Ensure to connect with postgres database

4. Run the application:
    ```bash
    go run main.go
    ```