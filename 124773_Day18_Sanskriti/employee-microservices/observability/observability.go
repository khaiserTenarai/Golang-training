package observability

import (
    "log"
    "os"
)

func NewLogger() *log.Logger {
    return log.New(os.Stdout, "[employee-microservices] ", log.LstdFlags|log.LUTC|log.Lmicroseconds)
}
