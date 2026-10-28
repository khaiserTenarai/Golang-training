package config

import "os"

type Config struct {
    GatewayPort string
    EmployeePort string
    DepartmentPort string
    SalaryPort string
    EmployeeDBURL string
    DepartmentDBURL string
    SalaryDBURL string
}

func get(key, fallback string) string {
    if value := os.Getenv(key); value != "" { return value }
    return fallback
}

func Load() Config {
    return Config{
        GatewayPort: get("GATEWAY_PORT", "8080"),
        EmployeePort: get("EMPLOYEE_PORT", "8081"),
        DepartmentPort: get("DEPARTMENT_PORT", "8082"),
        SalaryPort: get("SALARY_PORT", "8083"),
        EmployeeDBURL: get("EMPLOYEE_DB_URL", "postgres://postgres:postgres@localhost:5432/employee_db"),
        DepartmentDBURL: get("DEPARTMENT_DB_URL", "postgres://postgres:postgres@localhost:5432/department_db"),
        SalaryDBURL: get("SALARY_DB_URL", "postgres://postgres:postgres@localhost:5432/salary_db"),
    }
}
