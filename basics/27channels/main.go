package main

import (
	"fmt"
	"sync"
)

// channels are a way/pipeline for go routines to talk to each other
func main() {
	fmt.Println("Channels in GOlang")

	//use make() usually for creating channels, has a particular pipeline
	// myCh := make(chan int)//below is buffered channel of size 5
	myCh := make(chan int, 5)
	wg := &sync.WaitGroup{}

	//pushing values in channel, arrow always points left, not like fat arrow in JS
	// myCh <- 5 //like item going into a box
	// fmt.Println(<-myCh) //getting values out of box/channel
	wg.Add(2)
	go func(chnl chan int, waitGrp *sync.WaitGroup) {
		val, isChnlOpn := <-myCh
		fmt.Println(isChnlOpn) //based on channel open you can consume
		fmt.Println(val)       //1st consumption of channel val = 5

		//2nd consumption of channel val = 6
		fmt.Println(<-myCh) //reading the channel, some way is there to read the channel even if it closes
		//listening to closed channel is not an issue
		waitGrp.Done()
	}(myCh, wg)

	go func(chnl chan int, waitGrp *sync.WaitGroup) {
		myCh <- 5 //only one consumer to read this 1st item from channel
		myCh <- 6
		myCh <- 7   //only one is consuming this so will throw error, use a buffered channel or loop of consumer
		close(myCh) //closing the channel once putting the values
		waitGrp.Done()
	}(myCh, wg)

	wg.Wait()
}

/*
Fantastic question! Let’s **solve it by building a real mental model** around channels, deadlocks, and how they relate to `WaitGroup`. 🧠⚙️

---

## 😵 Why do **channels give deadlock errors**?

### 🔥 TL;DR: Deadlock = a goroutine is **waiting forever** on a channel op that can’t happen.

Go's runtime **detects** when all goroutines are blocked and **panics with:**

```
fatal error: all goroutines are asleep - deadlock!
```

---

## 🧩 Example 1: UNBUFFERED channel causes deadlock

```go
func main() {
	ch := make(chan int)
	ch <- 5 // ⛔ BLOCKS here forever, no one to receive
}
```

### ❌ Why this deadlocks:
- `ch <- 5` is **sending**
- But no one is **receiving**
- So the main goroutine blocks **forever**

> Go sees: “nothing's making progress” → panic 🚨

---

## ✅ Fix: Launch a receiver

```go
func main() {
	ch := make(chan int)

	go func() {
		val := <-ch // receiving goroutine
		fmt.Println("Received:", val)
	}()

	ch <- 5 // now it works!
}
```

---

## 🔄 Rule of Thumb:

| Channel Type     | Send blocks until | Receive blocks until |
|------------------|-------------------|-----------------------|
| `make(chan T)`   | receiver ready    | sender sends          |
| `make(chan T, N)`| buffer is full    | buffer has values     |

---

## 🧠 Use WaitGroup WITH Channels (real scenario)

Let’s say you want to **scrape multiple websites**, and **collect the results**.

---

### 🛠️ Problem: Launch multiple goroutines, return results

Here’s how you use **WaitGroup + channel** safely:

```go
func main() {
	sites := []string{
		"https://example.com",
		"https://google.com",
		"https://golang.org",
	}

	var wg sync.WaitGroup
	results := make(chan string, len(sites)) // buffer to avoid blocking

	for _, site := range sites {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			// fake work
			time.Sleep(1 * time.Second)
			results <- "Scraped: " + url
		}(site)
	}

	// 🔒 Wait in a separate goroutine to close the channel
	go func() {
		wg.Wait()
		close(results)
	}()

	// 🔁 Safely receive until channel closes
	for res := range results {
		fmt.Println(res)
	}
}
```

---

## 💡 Why use WaitGroup *with* channel?

| Problem                             | Solution              |
|-------------------------------------|------------------------|
| Don’t exit main before goroutines finish | `WaitGroup.Wait()`   |
| Don’t write to closed channel        | Close after `Wait()`  |
| Safely iterate over results          | Use `for range ch`    |

---

## ⚠️ What causes channel deadlocks? (Summary)

| Code Pattern                         | What’s wrong?                     |
|--------------------------------------|-----------------------------------|
| Sending to channel, no receiver      | Deadlock                          |
| Receiving from empty channel         | Deadlock                          |
| Writing after `close(channel)`       | Panic                             |
| Reading after all writers done but channel not closed | Infinite wait       |

---

## 🔚 Golden Rules

✅ Always pair sends with receives
✅ If you range over a channel, you **must** close it
✅ Use WaitGroup to ensure goroutines complete before main exits
✅ Use `buffered channels` to prevent unnecessary blocking if receiver comes later

---
Great follow-up! Let’s hit that 🔥 directly:

---

## 🧠 So, what's the **use of a channel** in Go?

### ➕ Channels let goroutines **communicate** and **synchronize** — safely and easily.

Think of a channel as:
> A **pipe** where one goroutine sends data 💌 and another one receives it 📬.

It’s **not just about async**, it’s about **coordinating** between goroutines.

---

## ✅ Use Cases for Channels

| Use Case                            | How Channels Help                        |
|-------------------------------------|------------------------------------------|
| Fan out → Fan in                    | Multiple goroutines send results to one  |
| Async file/network operations       | Stream results as soon as they're ready  |
| Sync points between goroutines      | Block until a message arrives            |
| Data pipelines                      | Pass data through multiple stages        |
| Worker pool                         | Distribute tasks to workers via channel  |

---

## 📂 Use Channels to Read Files Async?

Yes, absolutely! Let’s walk through it:

### ⚙️ Scenario: Read 5 files concurrently, and collect their line counts

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	files := []string{"file1.txt", "file2.txt", "file3.txt"}
	var wg sync.WaitGroup
	results := make(chan string, len(files))

	for _, fname := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()

			f, err := os.Open(file)
			if err != nil {
				results <- fmt.Sprintf("Error opening %s", file)
				return
			}
			defer f.Close()

			scanner := bufio.NewScanner(f)
			count := 0
			for scanner.Scan() {
				count++
			}
			results <- fmt.Sprintf("File: %s, Lines: %d", file, count)
		}(fname)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println(res)
	}
}
```

---

## 🧪 What’s happening here?

- 🔄 Spawn goroutines to read each file
- 📦 Each one sends result to `results` channel
- 🧯 `WaitGroup` ensures we wait before exiting
- 🧼 `close(results)` so `range` doesn’t hang

---

## 💬 Without channels?

- You’d have to store results in a `[]string` with locks/mutex
- It’s error-prone and verbose
- Channels make your code *elegant* and *safe* ✨

---

## TL;DR — Channels in real life:
> Use channels whenever **goroutines need to send results, signals, or events** to each other. They’re Go’s elegant built-in alternative to shared memory, mutexes, or queues.

---
*/
