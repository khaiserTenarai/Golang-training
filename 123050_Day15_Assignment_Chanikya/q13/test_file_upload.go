package main

/*
# Task 13 – Employee File Upload API

## Complete Steps to Run and Test

### 1. Project Location

Open PowerShell and navigate to the project folder:

```
cd "C:\Training\Assignments\Day_15\13-Implement File Upload\ems"
```

### 2. Check Project Files

Ensure your project contains the required Go files, including:

- `main.go`
- `config/`
- `controller/`
- `middleware/`
- `auth/`
- `model/`
- `repository/`
- `service/`
- `.env`
- `go.mod`

### 3. Verify the Configuration

Open `.env` and confirm that the database and server settings are correct.

Example:

```
DB_HOST=localhost
DB_PORT=5432
DB_NAME=employee_db
DB_USER=postgres
DB_PASSWORD=Root
DB_SSLMODE=disable

SERVER_PORT=10005

JWT_SECRET=employee-management-app-jwt-secret
JWT_EXPIRY_MINUTES=60

ADMIN_USERNAME=admin
ADMIN_PASSWORD=Admin@12345

USER_USERNAME=user
USER_PASSWORD=User@12345
```

Use the actual credentials configured on your computer.

### 4. Start PostgreSQL

Ensure PostgreSQL is running and the `employee_db` database exists.

If the database is unavailable, the Go application may fail to start.

### 5. Start the Go Application

Run:

```
go run .
```

Expected terminal output:

```
Server running at http://localhost:10005
```

Keep this terminal open while testing the API.

### 6. Test the Health Endpoint

Open a second PowerShell window and run:

```
curl.exe http://localhost:10005/health
```

Expected response:

```
{
  "status": "healthy"
}
```

This confirms that the application and database health check are working.

### 7. Obtain an Admin JWT Token

Send a POST request to:

```
http://localhost:10005/login
```

Example using PowerShell:

```
$body = @{
    username = "admin"
    password = "Admin@12345"
} | ConvertTo-Json

Invoke-RestMethod `
    -Uri "http://localhost:10005/login" `
    -Method POST `
    -ContentType "application/json" `
    -Body $body
```

Copy the JWT token returned by the login endpoint. Use the actual response field containing the token.

**Note:** The token expires according to `JWT_EXPIRY_MINUTES`. Generate a new token if it expires.

### 8. Create a Valid PDF Test File

Run the following command in PowerShell to create a small PDF file in your project directory:

```
@'
%PDF-1.4
1 0 obj
<< /Type /Catalog /Pages 2 0 R >>
endobj
2 0 obj
<< /Type /Pages /Kids [3 0 R] /Count 1 >>
endobj
3 0 obj
<< /Type /Page /Parent 2 0 R /MediaBox [0 0 300 144] /Resources << /Font << /F1 4 0 R >> >> /Contents 5 0 R >>
endobj
4 0 obj
<< /Type /Font /Subtype /Type1 /BaseFont /Helvetica >>
endobj
5 0 obj
<< /Length 44 >>
stream
BT /F1 18 Tf 30 80 Td (EMS File Upload Test) Tj ET
endstream
endobj
trailer
<< /Root 1 0 R /Size 6 >>
%%EOF
'@ | Set-Content -Path ".\sample.pdf" -Encoding ASCII
```

The file will be created at:

```
C:\Training\Assignments\Day_15\13-Implement File Upload\ems\sample.pdf
```

### 9. Upload the File

Use the admin JWT from Step 7.

Replace `PASTE_FRESH_ADMIN_TOKEN_HERE` with the token you obtained from login.

```
curl.exe -X POST "http://localhost:10005/employees/upload" -H "Authorization: Bearer PASTE_FRESH_ADMIN_TOKEN_HERE" -F "file=@C:\Training\Assignments\Day_15\13-Implement File Upload\ems\sample.pdf"
```

Expected response:

```
{
  "filename": "generated_filename.pdf",
  "message": "File uploaded successfully",
  "path": "uploads/generated_filename.pdf",
  "size": 452
}
```

The generated filename and size may differ.

### 10. Verify the Uploaded File

The API saves the uploaded file in the `uploads` directory relative to the application's working directory.

List the uploaded files:

```
Get-ChildItem .\uploads
```

Verify that the generated PDF exists in this directory.

### 11. Test Using Hoppscotch

Open:

[https://hoppscotch.io/](https://hoppscotch.io/)

Configure the request:

**Method**

```
POST
```

**URL**

```
http://localhost:10005/employees/upload
```

**Authorization**

- Select `Bearer Token`.
- Paste a fresh admin JWT.

**Body**

- Select `multipart/form-data`.
- Add a field named `file`.
- Attach the actual PDF file to that field.
- Allow Hoppscotch to generate the multipart `Content-Type` header automatically.

Click **Send**.

If you use Hoppscotch from the browser, your Go server's CORS middleware must allow the Hoppscotch origin and handle `OPTIONS` preflight requests.

### 12. Common Errors and Solutions

| Error                                              | Cause                                           | Solution                                |
| -------------------------------------------------- | ----------------------------------------------- | --------------------------------------- |
| `File is required`                                 | Missing file or incorrect multipart field name  | Use the exact field name `file`         |
| `File content does not match an allowed file type` | File content does not match an allowed type     | Use a genuine PDF, JPG, JPEG, or PNG    |
| `Failed to open/read local data`                   | Incorrect file path or inaccessible file        | Verify the path and file existence      |
| `401 Unauthorized`                                 | Missing, invalid, or expired JWT                | Log in again and use a fresh token      |
| `403 Forbidden`                                    | User lacks the required admin role              | Authenticate with an admin account      |
| `Network Error` in Hoppscotch                      | CORS, browser access, or interceptor issue      | Check CORS and the selected interceptor |
| Database connection failure                        | PostgreSQL unavailable or incorrect credentials | Start PostgreSQL and verify `.env`      |

### 13. Complete Execution Flow

1. Start PostgreSQL.
2. Navigate to the EMS project directory.
3. Run `go run .`.
4. Check `GET /health`.
5. Log in through `POST /login`.
6. Obtain a valid admin JWT.
7. Create or select a valid PDF/image.
8. Send `POST /employees/upload` with Bearer Token authentication.
9. Confirm the success response.
10. Verify the file in the `uploads` directory.

**Expected result:** The API validates the request, authenticates the admin, checks the file, generates a filename, saves the file, and returns a success response.
*/
