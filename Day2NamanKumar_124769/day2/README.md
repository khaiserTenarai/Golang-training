# Day 2 — Go Data Types, Control Flow & Collections

**Module:** `day2` · **Go:** 1.22+ · **No external dependencies**

Each task is a separate runnable program in `cmd/`. The logic for the bigger tasks lives in `internal/` packages with unit tests (94–100% coverage).

## How to run

```bash
cd day2
go run ./cmd/01-variables          # change the folder name for each task
go test -cover ./...               # run all tests
```

With `make` (Mac/Linux): `make run T=05-salary`, `make mini`, `make all`.

| # | Task | Command |
|---|------|---------|
| 1 | Variables & constants | `go run ./cmd/01-variables` |
| 2 | Primitive data types | `go run ./cmd/02-datatypes` |
| 3 | Zero values | `go run ./cmd/03-zerovalues` |
| 4 | Type conversion | `go run ./cmd/04-conversion` (or add values: `... 123 4.5 abc`) |
| 5 | Salary calculator | `go run ./cmd/05-salary 50000 10 5` (basic, OT hours, bonus %) |
| 6 | Comparison & logical operators | `go run ./cmd/06-operators` |
| 7 | String processing | `go run ./cmd/07-strings "your text here"` |
| 8 | Bytes vs runes | `go run ./cmd/08-runes` |
| 9 | 12-month sales array | `go run ./cmd/09-sales` |
| 10 | Employee CRUD with a slice | `go run ./cmd/10-slice-crud` |
| 11 | Employee lookup with a map | `go run ./cmd/11-map-lookup 103 999` |
| 12 | Menu using switch | `go run ./cmd/12-switch-menu` (interactive) |
| 13 | Five for-loop problems | `go run ./cmd/13-loops 12321` |
| 14 | Scope & shadowing | `go run ./cmd/14-scope` |
| 15 | **Mini project** | `go run ./cmd/15-employee-cli` (interactive) |

## Project structure

```text
day2/
├── cmd/                  # one runnable program per task (01 … 15)
├── internal/
│   ├── salary/           # Task 5 pay-slip maths + tests
│   ├── textstats/        # Task 7 text counting + tests
│   ├── loops/            # Task 13 factorial, fibonacci, prime… + tests
│   └── employee/         # Task 15 slice + map store + tests
├── Makefile
└── README.md
```

---

## Task-by-task explanation

### 1. Variables and constants

**What it shows:** the six ways to declare a variable (`var name string = …`, `var age = 29`, `salary := …`, several at once, a grouped `var ( … )` block, and declare-then-assign), multiple assignment (swap in one line), the blank identifier `_`, typed and untyped constants, and `iota` for enum-style constants (Intern, Junior, Senior, Lead, Manager).

**Say:** "`:=` is the short form and only works inside functions. `var` works everywhere, including at package level. Constants are fixed at compile time, so assigning to one is a compile error. `iota` auto-numbers constants starting from 0, which is how Go does enums."

**Possible question:** *Typed vs untyped constant?* An untyped constant like `24` adapts to where it's used (int or float). A typed constant like `TaxRate float64` is always that type.

### 2. Primitive data types

**What it shows:** every integer size (int8 to int64, uint8 to uint64) with its max value and memory size, float32 vs float64 precision, complex numbers, bool, string, byte, and rune. It also shows two classic traps: **overflow** (`uint8 255 + 1 = 0`) and **float precision** (`0.1 + 0.2 = 0.30000000000000004`).

**Say:** "Go makes you pick exact sizes. Use `int` and `float64` by default. `byte` is just another name for uint8 and `rune` is another name for int32. Integers silently wrap around when they overflow, and floats can't store 0.1 exactly, so for money we use integer paise or a decimal library."

### 3. Zero values

**What it shows:** the default value of every type: `0`, `""`, `false`, and `nil` for pointers, slices, maps, channels, functions, interfaces and errors. A struct's zero value is all its fields at zero. It also shows what's safe (appending to a nil slice, reading a nil map) and what panics (writing to a nil map, dereferencing a nil pointer), using `recover` to catch the panic.

**Say:** "Go never leaves memory uninitialised, so there are no garbage values. But `nil` maps and pointers still need care: always create maps with `make`, and check pointers for nil before using them."

### 4. Type conversion (string → int → float)

**What it shows:** `strconv.Atoi` (string → int), `float64(n)` (int → float), and `strconv.ParseFloat`, with proper error handling for bad input (`"abc"`), decimals (`"3.99"`), spaces, and out-of-range numbers. Extras: float → int truncates, `7/2 = 3` but `float64(7)/2 = 3.5`, the `string(rune(65)) = "A"` trap, and int8 overflow.

**Say:** "Go never converts types automatically; every conversion is explicit. Converting from a string can fail, so these functions return an error that we always check. I also added a range check, because converting a huge float to int gives a garbage value instead of an error."

### 5. Employee salary calculator

**What it shows:** a full monthly pay slip using every arithmetic operator. HRA (40%) and DA (10%) use `*`; per-day and per-hour pay use `/`; overtime at 1.5×; bonus %; gross uses `+`; PF, tax and professional tax are deducted with `-`; annual pay is `* 12`; and `%` splits the net pay into ₹500 notes plus a remainder. It also demonstrates the assignment operators `+= -= *= /= %= ++`. Input is validated.

**Say:** "The calculation is in its own package, `internal/salary`, with unit tests that check the numbers. The rates are simple example values for learning, not real tax rules."

### 6. Comparison and logical operators

**What it shows:** `== != < <= > >=` (including string comparison), `&& || !`, a truth table, and real HR rules for bonus, promotion and review eligibility, such as `(experience >= 3 && rating >= 4.0 || isManager) && !onNotice`. It also demonstrates **short-circuit evaluation**, including the common nil-pointer guard `if e != nil && e.Salary > 50000`.

**Say:** "`&&` stops as soon as it finds false, and `||` stops as soon as it finds true. That's why checking `e != nil` first prevents a crash: the second part never runs."

### 7. String processing

**What it shows:** counts characters (with and without spaces), words, vowels and digits, plus consonants, spaces, punctuation, lines and bytes. It accepts text as an argument or typed input. Words are counted with `strings.Fields`, which handles multiple spaces and tabs correctly. Characters are counted as runes, so Kannada or Hindi text is counted correctly.

**Say:** "I loop over the string once with `for range`, which gives runes, not bytes, and a `switch` puts each rune in a category. The logic is in `internal/textstats` with tests, including a Kannada test."

### 8. Bytes vs runes

**What it shows:** `len("ಕನ್ನಡ")` is 15 bytes but only 5 runes. English letters take 1 byte, while Indian scripts take 3 bytes per character (Kannada, Hindi, Tamil, Telugu, Bengali, plus accented text and emoji). It loops byte-by-byte vs rune-by-rune, shows the **slicing bug** (`hindi[:2]` cuts a character in half) and the fix (`[]rune`), and reverses a string by bytes (broken) vs runes (correct).

**Say:** "A Go string is UTF-8 bytes. A rune is one Unicode code point. So `len()` gives bytes, and `utf8.RuneCountInString` gives characters. Always use `for range` or `[]rune` for non-English text. One more detail: ಕನ್ನಡ looks like 3 letters but has 5 runes, because the virama joins consonants together. Counting visible letters needs a special Unicode library."

### 9. Array of 12 months of sales

**What it shows:** a fixed `[12]float64` array with total, average, best and worst month, a text bar chart with above-average markers, quarterly totals (by slicing the array), and Jan → Dec growth %. It also shows that arrays are **values**: assigning copies all 12 elements.

**Say:** "An array's size is part of its type, so `[12]float64` can never grow. That makes it perfect for fixed data like 12 months. Unlike slices, assigning an array makes a full copy."

### 10. Employee CRUD using a slice

**What it shows:** Create (`append`, with a duplicate-ID check), Read (linear search), Update (by index), and Delete (`append(s[:i], s[i+1:]...)`, keeping the order). It prints `len` and `cap` so you can watch capacity grow 1 → 2 → 4, and shows `copy()` creating an independent copy.

**Say:** "A slice is a growable view over an array. When it's full, `append` allocates a bigger array, roughly doubling. For update I change `employees[idx]` directly, because the value in a `for range` loop is a copy and changing it wouldn't update the slice. Search here is O(n); Task 11 makes it O(1) with a map."

### 11. Employee lookup using a map

**What it shows:** `map[int]Employee` with create, add, update (copy, modify, store back), lookup with the **comma-ok idiom** (`emp, ok := byID[id]`), delete, iterating in sorted order (map order is random), a second index (name → ID), and grouping by department with `map[string][]Employee`.

**Say:** "A map gives instant lookup by key. A missing key silently returns the zero value, so I always use comma-ok to know whether it really exists. Maps have no fixed order, so to print them in order I sort the keys first."

### 12. Menu-driven program using switch

**What it shows:** a bank account menu (balance, deposit, withdraw, history, tier) that demonstrates every switch form: value switch, multiple values in one case (`"0", "q", "exit"`), tagless switch (like if/else-if), switch with an init statement, `fallthrough`, and a type switch. It enforces rules like a minimum balance of ₹1,000 and positive amounts.

**Say:** "Go's switch doesn't need `break`; it stops after the matching case automatically. If you actually want to run the next case too, you write `fallthrough`. A switch with no value works like a clean if/else-if chain."

### 13. Five for-loop problems

**What it shows:**
- **Factorial** with a classic loop, refusing values above 20! because they overflow uint64.
- **Fibonacci** using `a, b = b, a+b`.
- **Prime** checking divisors only up to √n.
- **Reverse number** with `% 10` to take the last digit and `/ 10` to drop it.
- **Palindrome** for numbers and strings, including Kannada text like ಕನಕ.

It also demonstrates every form of Go's `for`: classic, while-style, infinite with `break`/`continue`, `range` over an integer (Go 1.22), and a labeled break.

**Say:** "Go has only one loop keyword, `for`, but it covers every style. For primes I only test up to the square root, because any bigger factor would pair with a smaller one I've already checked."

### 14. Scope and shadowing

**What it shows:** package, function, block, and if-init scope. Shadowing hides an outer variable with an inner one of the same name. It demonstrates two **real bugs**: `count := 5` inside an `if` creates a new variable (the outer one stays 0), and the classic `err` shadowing bug where a function returns `nil` even though validation failed. It also shows closures capturing an outer variable.

**Say:** "`:=` always declares a *new* variable in the current block. Inside an `if` or loop, that can silently hide the outer variable. The fix is to use `=` when you mean to assign. The `shadow` vet tool can find these automatically."

### 15. Mini project: Employee Management CLI

**What it shows:** an interactive menu with 8 features:
1. Add, with validation and auto IDs starting at 101
2. List, sorted by ID, name, salary or experience
3. Search by ID or part of a name
4. Update department and/or salary
5. Delete, with a y/n confirmation
6. Filter by department
7. Give a percentage raise to a department or everyone
8. Reports: payroll total, average, highest and lowest salary, a department report, and experience levels

Salaries are shown in Indian format (12,34,567).

**How each concept is used:**
- **Slice** `[]Employee` keeps records in order.
- **Map** `map[int]int` (ID → position) gives O(1) lookup by ID. After a delete, the index is rebuilt because later records shift left. A second map groups departments for the report.
- **Loops** handle searching, raises, statistics and table printing.
- **Conditions** handle validation, `switch` menus and experience levels.
- **Functions** are small, with one job each. Business logic is in `internal/employee`, and input/output is in `main`.

**Demo order:** 2 (list) → 1 (add) → 3 (search `nair`) → 4 (update) → 5 (delete) → 7 (raise) → 8 (reports) → 9 (invalid input) → 0 (exit).

**Say:** "Using both a slice and a map gives the best of both: an ordered list and instant lookup. The tricky part was delete: after removing from the middle of the slice, every later position changes, so I rebuild the index. There's a unit test specifically for that."

**Possible question:** *Why not just a map?* A map has no order, so listing would come out random every time. The slice keeps insertion order.
