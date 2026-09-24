# Task 10 — Environment Variables

| Variable | Default |
|----------|---------|
| `APP_NAME` | `EmployeeApp` |
| `APP_PORT` | `8080` |

```bash
go run .                                   # Starting EmployeeApp on port 8080
APP_NAME="Employee API" APP_PORT=9090 go run .   # Starting Employee API on port 9090
APP_PORT=abc go run .                      # invalid APP_PORT "abc": must be a number 1-65535

# Windows PowerShell
$env:APP_PORT="9090"; go run .
```
`os.LookupEnv` distinguishes "unset" from "set to empty"; `os.Getenv` does not.
