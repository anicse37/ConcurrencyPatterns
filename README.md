
---

## Patterns Explained 🔍

### 1. **Worker Pool** (`worker_pool/main.go`)
A classic concurrency pattern where multiple worker goroutines process jobs from a shared job queue.

- **Goal:** Efficiently handle many tasks using a limited number of goroutines.
- **Concepts Used:** Goroutines, buffered channels, `sync.WaitGroup`.

### 2. **Fan-In / Fan-Out** (`fan_in_out/main.go`)
Distributes work across multiple goroutines (Fan-Out) and combines their output into a single channel (Fan-In).

- **Goal:** Parallelize independent tasks and merge results.
- **Concepts Used:** Channels, goroutines, synchronization.

### 3. **Pipeline** (`pipeline/V1/main.go` & `pipeline/V2/main.go`)
Creates a series of stages connected by channels, where each stage processes and passes data to the next.

- **Goal:** Break computation into stages, each stage doing part of the job.
- **Concepts Used:** Goroutines, channels, chaining.
- **V1:** Basic pipeline.
- **V2:** More advanced or optimized version.

### 4. **Atomic Operations** (`atomic/main.go`)
Demonstrates safe concurrent counter operations using `sync/atomic`.

- **Goal:** Update shared variables without using locks.
- **Concepts Used:** `sync/atomic` package.

---

## How to Run ▶️

Each example is self-contained. Use `bash script` to execute any pattern:

```bash
./run.sh