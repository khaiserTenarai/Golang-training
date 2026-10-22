# Architecture Decisions — Day 19 Go Microservices

This document records the major design decisions and reasoning behind them.

---

## Decision 1: Two Separate Services Instead of One Monolith

**What:** Split Employee and Department into independent microservices running on different ports.

**Why:** The assignment requires demonstrating service-to-service communication, independent deployability, and resilience patterns. Two services is the minimum needed to show these patterns meaningfully.

**Trade-off:** More complex than a monolith, but demonstrates real microservice patterns — circuit breaker, retry, correlation ID propagation — which only make sense with multiple services.

---

## Decision 2: Single Go Module with Multiple Main Packages

**What:** One `go.mod` at the root, with `employee-service/main.go` and `department-service/main.go` as separate executables sharing common packages.

**Why:** Avoids code duplication for shared types (DTOs, entities, middleware, resilience). Both services import from the same module. In production, each service would have its own repository, but for a training assignment, a mono-repo approach keeps things simple and self-contained.

**Alternative considered:** Separate go.mod per service — rejected because it would require duplicating DTO and resilience code.

---

## Decision 3: In-Memory Storage Instead of PostgreSQL

**What:** Both services use `map[int]Entity` with `sync.RWMutex` for thread-safe data storage.

**Why:** The assignment focus is microservice communication patterns, not database operations. In-memory storage eliminates the PostgreSQL dependency, making the project immediately runnable with zero setup.

**Trade-off:** Data is lost when services restart. Acceptable for training purposes.

---

## Decision 4: Entity/DTO Separation

**What:** Internal domain entities (`entity/entity.go`) are kept separate from API-facing DTOs (`dto/dto.go`).

**Why:**
- **Entities** represent internal data structure — may contain fields not meant for API consumers
- **DTOs** control exactly what the API exposes and accepts
- Changes to the internal model don't break the API contract
- DTOs include JSON tags; entities don't need them

**Example:** `entity.Employee` has a raw `time.Time` for `CreatedAt`, while `dto.EmployeeResponse` exposes it as a formatted string. The employee DTO also has an optional `Department` field that doesn't exist on the entity.

---

## Decision 5: Exponential Backoff for Retries

**What:** Failed HTTP calls are retried up to 3 times with delays of 200ms, 400ms, 800ms.

**Why:** Simple fixed-delay retries can overwhelm a struggling service. Exponential backoff gives the downstream service progressively more time to recover between attempts. 3 attempts is enough to handle transient network blips without excessive delay.

**Configuration:** `resilience.RetryConfig{MaxAttempts: 3, BaseDelay: 200ms}`

---

## Decision 6: Circuit Breaker with Three States

**What:** The circuit breaker transitions through CLOSED → OPEN → HALF-OPEN.

**Why:**
- **CLOSED** (normal): requests pass through. Failure count tracked.
- **OPEN** (after 3 failures): requests rejected immediately for 10 seconds. Prevents wasting resources calling a dead service and avoids cascading failures.
- **HALF-OPEN** (after cooldown): one test request allowed. If it succeeds twice, back to CLOSED. If it fails, back to OPEN.

**Why 3 failures:** Low enough to react quickly, high enough to avoid false positives from a single network glitch.

**Why 10-second cooldown:** Gives the downstream service time to recover without hammering it with probes.

---

## Decision 7: Correlation ID as Middleware

**What:** Every request gets a unique `X-Correlation-ID` header. When employee-service calls department-service, the same ID is forwarded.

**Why:** In a microservice architecture, a single user request touches multiple services. Without correlation IDs, debugging is extremely difficult — you'd have disconnected log entries across services with no way to link them. The correlation ID is the single thread that ties all logs together.

**Implementation:** Middleware generates the ID early in the request pipeline, stores it in the Go context, and all downstream HTTP calls extract it from the context and add it to outgoing request headers.

---

## Decision 8: Standard Library Only (No External Dependencies)

**What:** The project uses only Go's standard library (`net/http`, `encoding/json`, `sync`, `time`, `testing`).

**Why:** Keeps the project lightweight and immediately runnable. No `go mod tidy` needed to download packages. Every pattern (routing, middleware, circuit breaker, retry) is implemented from scratch, which is more educational than importing a library.

**Trade-off:** More code than using gorilla/mux, go-kit, or gobreaker — but the student sees exactly how each pattern works internally.

---

## Decision 9: Contract Tests with httptest

**What:** Tests create mock HTTP servers using `httptest.NewServer` and verify that API responses match the expected DTO shapes.

**Why:** Contract tests verify that the API contract between services hasn't broken. Unlike integration tests, they don't require running services — they're fast, deterministic, and isolated.

**What's tested:**
- Response status codes (200, 201, 400)
- Response body shape (matches DTO structs)
- Header propagation (correlation ID)
- Error responses for invalid input

---

## Decision 10: Graceful Degradation for Department Enrichment

**What:** If department-service is unavailable, employee-service still returns employee data — just with `"department": null` instead of failing the entire request.

**Why:** In microservices, partial failure is normal. The employee data is still valuable without department details. Failing the entire request because one dependency is down would be a poor user experience.

**How:** The `GetDepartment()` method returns `nil` on any error. The DTO's `Department` field is a pointer (`*DepartmentResponse`) with `omitempty` — so it's cleanly absent from the response when unavailable.

---

## Summary of Key Patterns

| Pattern | Purpose | Location |
|---------|---------|----------|
| DTO/Entity split | Decouple API from domain | `dto/` vs `entity/` |
| Externalized config | Change behavior without code changes | `config/config.go` |
| Correlation ID | Trace requests across services | `middleware/correlation.go` |
| Retry + backoff | Handle transient failures | `resilience/retry.go` |
| Circuit breaker | Prevent cascading failures | `resilience/circuit_breaker.go` |
| Health checks | Monitor service + dependency health | `/health` endpoints |
| Contract tests | Verify API contracts without running services | `tests/contract_test.go` |
| Graceful degradation | Return partial data when deps fail | `httpclient/client.go` |
