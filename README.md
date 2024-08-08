# RosyPost Backend

## Project Overview
The RosyPost backend provides the server-side functionality for the RosyPost newsletter application, handling email creation, sending, and management.

## Features
- Handle email creation and sending.
- Manage newsletter data.
- Provide a robust API for frontend integration.

## Technologies Used
- **Backend**: Go (Gin framework), GORM (for ORM)
- **Database**: PostgreSQL (or specify your database)
- **Email**: SMTP (e.g., Gmail, SendGrid)

## Installation

### Prerequisites
- Go (for backend development)
- PostgreSQL (or your chosen database)
- A configured SMTP server (e.g., Gmail, SendGrid)

### Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/AliAngelinaaa/witu-newsletter-backend
   cd rwitu-newsletter-backend
   ```

2. Install Go dependencies:
   ```bash
   go mod tidy
   ```

3. Set up environment variables for SMTP and database:
   ```bash
   export SMTP_SERVER=smtp.example.com
   export SMTP_PORT=587
   export SMTP_USER=your-email@example.com
   export SMTP_PASS=your-email-password
   export DATABASE_URL=your-database-url
   ```

4. Run the backend application:
   ```bash
   go run main.go
   ```

## Configuration
- **SMTP Configuration**: Ensure SMTP settings are correctly set in your environment variables.
- **Database Configuration**: Update `DATABASE_URL` with your database connection string.

## Usage
1. Start the backend application as described in the Setup section.
2. The backend API will be available at the configured URL (default is `http://localhost:8080`).
3. Use API endpoints to manage newsletters and send emails.

## API Documentation

### Endpoints
- **Create Email**
  - **Method**: `POST`
  - **Endpoint**: `/send-email`
  - **Request Body**:
    ```json
    {
      "to": "recipient@example.com",
      "subject": "Newsletter Subject",
      "body": "Newsletter Content"
    }
    ```

- **Manage Newsletters**
  - **Fetch Newsletter**: `GET /posts/:id`
  - **Create Newsletter**: `POST /posts`
  - **Update Newsletter**: `PATCH /posts/:id`
  - **Delete Newsletter**: `DELETE /posts/:id`

## Contributing
- Fork the repository and clone your fork.
- Create a new branch for your feature or bug fix.
- Commit your changes and push them to your fork.
- Open a pull request against the main repository.

## License
- Distributed under the [MIT License](LICENSE).

---

This README is tailored specifically for the backend, providing relevant details and instructions.