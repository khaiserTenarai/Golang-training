# Task 07 — Package Design

```
07-package-design/
├── main.go            package main     → entry point, wiring only
├── employee/          package employee → business/domain logic
└── utils/             package utils    → generic helpers (no business rules)
```
Dependency direction: `main → employee → utils`. `utils` imports nothing from the project, so it never causes import cycles.

Run: `go run .`
```
Before raise: [101] Asha Rao - ₹12,50,000.00
After 10% raise: [101] Asha Rao - ₹13,75,000.00
```
