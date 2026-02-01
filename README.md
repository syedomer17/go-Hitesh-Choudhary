# 🔥 GO (GOLANG) — COMPLETE TOPICS & CONCEPTS LIST

---

## 1. Language Basics (Non-Negotiable)

* Go philosophy & design goals
* `package` system
* `main` package and `main()` function
* Imports (single, multiple, aliased, blank `_`)
* Comments (`//`, `/* */`)
* Keywords and identifiers
* Formatting (`gofmt`, `go vet`)
* Go compiler vs interpreter mindset

---

## 2. Variables, Constants & Types

* `var`, `:=`
* Explicit vs inferred types
* Zero values
* Constants (`const`)
* Untyped constants
* Type aliases vs defined types
* Type conversions (NO implicit casting)
* `iota`

---

## 3. Built-in Data Types

* Booleans
* Integers (`int`, `int8`… `uint64`)
* Floating points (`float32`, `float64`)
* Complex numbers
* Strings (immutable)
* Runes
* Bytes
* `uintptr`

---

## 4. Control Flow

* `if`, `else`
* Short variable declarations in `if`
* `switch` (expressionless, type switch)
* `for` (classic, while-style, infinite)
* `break`, `continue`
* `goto` (yes, exists — rarely used)
* `defer`
* `panic` & `recover`

---

## 5. Arrays, Slices & Memory

* Arrays (fixed length)
* Slices (dynamic)
* Slice internals (len, cap, backing array)
* Slice reallocation
* Copying slices
* Appending behavior
* Slice pitfalls & memory leaks
* Passing slices to functions

---

## 6. Maps

* Map declaration & initialization
* Zero value (`nil` maps)
* Reading/writing maps
* Checking key existence
* Map iteration order (randomized)
* Map as Set (`map[T]struct{}`)
* Maps & concurrency limitations

---

## 7. Strings & Text Processing

* UTF-8 encoding
* Bytes vs runes
* String iteration
* String concatenation
* `strings` package
* `strconv`
* Regular expressions (`regexp`)

---

## 8. Functions

* Function declaration
* Multiple return values
* Named return values
* Variadic functions
* Anonymous functions
* Closures
* Function types
* Passing functions as arguments
* Returning functions

---

## 9. Structs (Go’s “Objects”)

* Struct definition
* Field visibility (exported vs unexported)
* Anonymous structs
* Struct embedding (composition)
* JSON / DB tags
* Comparing structs
* Copy vs reference semantics

---

## 10. Methods

* Method receivers
* Pointer vs value receivers
* Method sets
* When methods are copied
* Receiver best practices

---

## 11. Interfaces (CRITICAL)

* Interface definition
* Implicit implementation
* Empty interface (`interface{}` / `any`)
* Type assertions
* Type switches
* Interface composition
* Interface pitfalls
* Designing small interfaces
* Dependency inversion in Go

---

## 12. Error Handling (Core Go Skill)

* `error` interface
* Creating errors
* Wrapping errors
* Sentinel errors
* Custom error types
* `errors.Is`, `errors.As`
* Error comparison
* When to panic
* Recover patterns
* Error handling anti-patterns

---

## 13. Pointers & Memory Model

* Pointers basics
* `&` and `*`
* Pointer vs value semantics
* Escape analysis
* Stack vs heap
* Garbage collection basics
* Common pointer mistakes

---

## 14. Packages & Modules

* Package naming conventions
* Export rules
* `init()` function
* Circular dependency rules
* Go modules (`go.mod`)
* Semantic versioning
* `replace` & `require`
* Private modules
* `internal` packages
* Project layout standards

---

## 15. Standard Library (Must-Know)

* `fmt`
* `os`
* `io`, `bufio`
* `path/filepath`
* `time`
* `math`
* `sort`
* `flag`
* `log`
* `context`
* `encoding/json`
* `encoding/xml`
* `net/http`
* `net/url`

---

## 16. File & OS Operations

* File read/write
* Permissions
* Directory traversal
* Environment variables
* Process execution
* Signals
* Cross-platform behavior

---

## 17. JSON & Serialization

* Marshal / Unmarshal
* Struct tags
* Omitempty
* Custom marshaling
* Streaming JSON
* Decoder vs Unmarshal
* Validation patterns

---

## 18. Concurrency (GO’S KILLER FEATURE)

* Goroutines
* Channels
* Buffered vs unbuffered channels
* Channel directions
* `select`
* Fan-in / fan-out
* Worker pools
* `sync.WaitGroup`
* `sync.Mutex`
* `sync.RWMutex`
* `sync.Once`
* `sync.Pool`
* Atomic operations
* Data races
* Deadlocks
* Livelocks

---

## 19. Context Management

* `context.Context`
* Cancellation
* Deadlines
* Timeouts
* Context propagation
* Context anti-patterns

---

## 20. Testing

* `testing` package
* Unit tests
* Table-driven tests
* Subtests
* Test coverage
* Benchmarks
* Fuzz testing
* Race detector
* Testable design patterns

---

## 21. Logging & Observability

* Standard logging
* Structured logging
* Log levels
* Context-aware logging
* Metrics basics
* Tracing concepts

---

## 22. HTTP & Networking

* HTTP server
* HTTP client
* Routing
* Middleware patterns
* Request lifecycle
* Headers, cookies
* CORS
* TLS basics
* Graceful shutdown
* Reverse proxies

---

## 23. CLI Development

* Argument parsing
* Flags
* Environment configs
* Exit codes
* Cross-platform binaries

---

## 24. Database Interaction

* `database/sql`
* Drivers
* Connection pooling
* Transactions
* Prepared statements
* Context-aware queries
* ORM vs raw SQL tradeoffs

---

## 25. Performance & Optimization

* Profiling (`pprof`)
* CPU profiling
* Memory profiling
* Benchmark analysis
* Reducing allocations
* Inlining
* Escape analysis tuning

---

## 26. Security Basics

* Input validation
* Secure random generation
* Hashing & encryption
* TLS usage
* Secrets management
* Common Go security pitfalls

---

## 27. Build & Tooling

* `go build`
* `go run`
* `go test`
* `go install`
* Cross compilation
* Static binaries
* Build tags

---

## 28. Deployment & Production

* Binary distribution
* Dockerizing Go apps
* Multi-stage builds
* Environment configs
* Graceful restarts
* Health checks

---

## 29. Advanced Topics (Senior Level)

* Go scheduler internals
* Garbage collector tuning
* Memory barriers
* Unsafe package (when NOT to use it)
* Plugin system
* CGO basics
* Reflection
* Generics (Go 1.18+)
* Type constraints
* Generic pitfalls

---

## 30. Go Design Philosophy (This Separates Seniors)

* Simplicity over cleverness
* Explicit > implicit
* Composition over inheritance
* Convention over configuration
* Readability > DRY
* Small interfaces
* Flat abstractions

---