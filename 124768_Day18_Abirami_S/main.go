package main

import (
    "context"
    "log"
    "os"
    "time"

    "github.com/joho/godotenv"

    "employee-microservices/config"
    "employee-microservices/discovery"
    "employee-microservices/employee"
    "employee-microservices/department"
    "employee-microservices/salary"
    "employee-microservices/eventbus"
    "employee-microservices/gateway"
    "employee-microservices/observability"
    "employee-microservices/saga"
)

func main() {
    _ = godotenv.Load()
    cfg := config.Load()

    logger := observability.NewLogger()
    bus := eventbus.New()
    bus.Start()

    // The demo runs services in one process so the in-memory event bus can
    // deliver events between them. Each service still has its own database.
    empRepo, err := employee.NewRepository(cfg.EmployeeDBURL)
    if err != nil {
        log.Fatalf("employee database: %v", err)
    }
    defer empRepo.Close()

    deptRepo, err := department.NewRepository(cfg.DepartmentDBURL)
    if err != nil {
        log.Fatalf("department database: %v", err)
    }
    defer deptRepo.Close()

    salRepo, err := salary.NewRepository(cfg.SalaryDBURL)
    if err != nil {
        log.Fatalf("salary database: %v", err)
    }
    defer salRepo.Close()

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    if err := empRepo.Ping(ctx); err != nil { log.Fatalf("employee database ping: %v", err) }
    if err := deptRepo.Ping(ctx); err != nil { log.Fatalf("department database ping: %v", err) }
    if err := salRepo.Ping(ctx); err != nil { log.Fatalf("salary database ping: %v", err) }

    deptService := department.NewService(deptRepo)
    salaryService := salary.NewService(salRepo)
    employeeService := employee.NewService(empRepo, deptService, bus)

    // Example async consumers. Events are delivered in the same process.
    bus.Subscribe("EmployeeCreated", func(e eventbus.Event) {
        logger.Printf("event=%s payload=%s", e.Name, e.Data)
        // This sample logs the event; salary initialization is an example hook.
        salaryService.HandleEmployeeCreated(e.Data)
    })
    bus.Subscribe("EmployeeCreated", func(e eventbus.Event) {
        logger.Printf("department consumer received event=%s payload=%s", e.Name, e.Data)
    })

    registry := discovery.NewRegistry()
    registry.Register("employee", "http://localhost:"+cfg.EmployeePort)
    registry.Register("department", "http://localhost:"+cfg.DepartmentPort)
    registry.Register("salary", "http://localhost:"+cfg.SalaryPort)

    empHandler := employee.NewController(employeeService)
    deptHandler := department.NewController(deptService)
    salaryHandler := salary.NewController(salaryService)

    // Each service has its own router/server, all started by this demo process.
    go runServer("employee", cfg.EmployeePort, empHandler.Routes(), logger)
    go runServer("department", cfg.DepartmentPort, deptHandler.Routes(), logger)
    go runServer("salary", cfg.SalaryPort, salaryHandler.Routes(), logger)

    sagaHandler := saga.NewController(employeeService, deptService)
    gatewayRouter := gateway.NewRouter(registry, sagaHandler, logger)
    logger.Printf("API Gateway listening on :%s", cfg.GatewayPort)
    if err := gatewayRouter.Run(":" + cfg.GatewayPort); err != nil {
        log.Fatal(err)
    }

    _ = os.Getenv("PORT")
}

func runServer(name, port string, router interface{ Run(...string) error }, logger *log.Logger) {
    logger.Printf("%s service listening on :%s", name, port)
    if err := router.Run(":" + port); err != nil {
        log.Printf("%s service stopped: %v", name, err)
    }
}
