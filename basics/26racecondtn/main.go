package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition tutorials")

	//wait group is usually created as a pointer
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	//If you're appending to score from multiple goroutines, you'll get a race condition.
	// Fix: Use sync.Mutex to lock access.
	var score = []int{10, 78}

	//example of IIFEs imm invoked funcs or lambdas or anon funcs & pass wait group
	wg.Add(5)                                    //either declare total at start of wg.Add(1) before each go-routine invocation
	go func(wg *sync.WaitGroup, m *sync.Mutex) { //since this wg is a reference, we need to pass a pointer
		fmt.Println("IIFE 1")
		//any write operation on shared resource, lock & unlock
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		wg.Done() //programmer's responsiblity to signal wg is done, reduce the wg counter
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("IIFE 2")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("IIFE 3")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("IIFE 4")
		m.Lock()
		score = append(score, 4)
		m.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("IIFE 5")
		m.Lock()
		score = append(score, 5)
		m.Unlock()
		wg.Done()
	}(wg, mut) /* <- this is the
	function call, happening immediately after the definition.
	You are passing wg as an argument to the anonymous function.
	This is what makes it an Immediately Invoked Function Expression (IIFE) —
	a pattern borrowed from other languages like JavaScript, but also used in Go.
	*/

	wg.Wait() //check race happening or not using 'go run --race'
	// exit status 66? : This helps CI pipelines or scripts detect failure due to races.
	fmt.Printf("score: %v\n", score) //NOTE: these are independent go routines, so order cannot be guaranteed !!
}

/*
Absolutely! That snippet is using an **anonymous goroutine** with a **closure**, and it's passing a pointer to a `WaitGroup`. Let me break it down for you step by step — especially that last part that looks confusing:

---

### 🔍 Code:
```go
go func(wg *sync.WaitGroup) {
    fmt.Println("IIFE 1")
    score = append(score, 1)
}(wg)
```

---

### 🔍 Breakdown:

#### 1. **`go func(wg *sync.WaitGroup) { ... }`**

- This defines **an anonymous function** (i.e., no name, just `func(...) { ... }`).
- It takes one argument: a pointer to a `sync.WaitGroup`.

#### 2. **`go` keyword**

- This runs the anonymous function as a **goroutine** — in a **new thread** of execution.
- It does **not block** the main thread.

#### 3. **Function body `{ ... }`**

- Inside this function, you can do anything — here it:
  - Prints `IIFE 1`
  - Appends `1` to a `score` slice (assuming it's declared outside)

#### 4. **The `(...)` at the end → `}(wg)`**

- That’s the **function call**, happening **immediately after the definition**.
- You are **passing `wg` as an argument** to the anonymous function.
- This is what makes it an **Immediately Invoked Function Expression (IIFE)** — a pattern borrowed from other languages like JavaScript, but also used in Go.

#### 5. **Why `wg *sync.WaitGroup`?**

- The `WaitGroup` is a struct — if you want to modify its state (like calling `.Done()`), you need to **pass it by reference** (i.e., pointer `*`), or else you'd be modifying a copy.

---

### ✅ Full Pattern With WaitGroup:

```go
var wg sync.WaitGroup
var score []int

wg.Add(1)
go func(wg *sync.WaitGroup) {
    defer wg.Done() // tell WaitGroup this goroutine is finished
    fmt.Println("IIFE 1")
    score = append(score, 1)
}(&wg) // ⬅️ Notice we passed pointer here
wg.Wait()
```

---

### TL;DR

| Code Part       | Meaning                                      |
|------------------|-----------------------------------------------|
| `go`             | Start a goroutine                            |
| `func(wg *...)`  | Anonymous function with one argument         |
| `{ ... }`        | Function body                                |
| `(wg)`           | Calling the function and passing `wg`        |

---
*/
