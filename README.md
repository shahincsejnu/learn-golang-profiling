# learn-golang-profiling

This repository contains hand-notes with hands-on examples and resources related to `profiling` in Golang. I recently explored this profiling stuffs and during the exploration I've noted these points/examples/resources here, gonna update this further on the fly later.

## Introduction to Profiling

- Basically Profiling is a technique to understand where our program spends its time and which functions are the most expensive in terms of CPU usage, memory allocation, network activity etc.
- It measures and help us to analyze the runtime behavior of a program.
- Thus we can identify bottlenecks, inefficiencies etc of our code.
- Based on this we can optimize the performance of our program.

## Golang Profiling

- Golang profiling is the process of measuring and analyzing data about the performance of a Go program.

> We'll use [pprof](https://github.com/google/pprof) package for golang profiling, it's a part of Go standard library `runtime/pprof`. The Go runtime generates profiling data in a `pprof` compatible format, we will need to collect this data and then use `pprof` to filter it and visualize further.

## Types of Profiling

There are several types of profiling, such as:

- CPU profiling:
  - To measure the amount of time spent executing different parts of the code.
  - To find the functions or sections of the code that are taking too long to execute.
- Memory profiling:
  - To measure the amount of memory being used by our program
  - To find out memory leaks where memory is being used inefficiently.
- Concurrency profiling:
  - To measure the performance of our program when using multiple goroutines or channels.
  - To identify race conditions or deadlocks issues.

## Go Profiles and It's types

- `pprof` works using profiles.
- A Profile is a collection of stack traces of a particular event, profile file has `protobuf` format.
- There are several types of Go profiles, such as:
  - `Goroutine`: stack traces of all current Goroutines
  - `CPU`: stack traces of CPU returned by runtime
  - `Heap`: memory allocations of live objects
  - `Allocation`: a sampling of all past memory allocations
  - `Threadcreate`: stack traces that led to the creation of new OS threads
  - `Block`: stack traces that led to blocking on synchronization primitives
  - `Mutex`: stack traces of holders of contended mutexes
  - `Trace`: for low-level program execution tracing

We can only profile one option at any given time.

## What is pprof?

- Go has a built-in package `runtime/pprof`, it helps us to profile our Go programs.
- It's a tool for visualization and analysis of profiling data.
- pprof reads a collection of profiling samples in `profile.proto` format and generates reports which we can visualize and analyze.
- We need to enable profiling in our Go program to use this `pprof`, for example we can call the `StartCPUProfile` function. It'll collect data about the execution of the program. After using this we can normally run our program by `go run .` or `go run main.go` like this way in a terminal, we also gonna need to use `StopCPUProfile` function in this case to stop profiling and save the data to a given file or io writer.
- After collecting data (saving to a file or io writer), we can use the `pprof` command line to analyze the data.

## Path to get profiles in default web server

> There is also a standard HTTP interface to profiling data. Adding the following line will install handlers under the `/debug/pprof/` URL to download live profiles: `import _ "net/http/pprof"`  
> See the net/http/pprof package for more details.  
> See Example-4 also for ref.

- http://localhost:6060/debug/pprof/goroutine
- http://localhost:6060/debug/pprof/heap
- http://localhost:6060/debug/pprof/threadcreate
- http://localhost:6060/debug/pprof/block
- http://localhost:6060/debug/pprof/mutex
- http://localhost:6060/debug/pprof/profile
- http://localhost:6060/debug/pprof/trace?seconds=5

To analyze these profiles, the tool to use is `go tool pprof`, which is a bunch of tools for visualizing stack traces.

### Usage Examples

1. Use the `pprof` tool to look at the heap profile:

```
go tool pprof http://localhost:6060/debug/pprof/heap
```

2. to look at a 30-second CPU profile:

```
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
```

3. to look at the goroutine blocking profile, after calling runtime.SetBlockProfileRate in your program:

```
go tool pprof http://localhost:6060/debug/pprof/block
```

4. to look at the holders of contended mutexes, after calling runtime.SetMutexProfileFraction in your program:

```
go tool pprof http://localhost:6060/debug/pprof/mutex
```

The package also exports a handler that serves execution trace data for the "go tool trace" command. To collect a 5-second execution trace:

```
curl -o trace.out http://localhost:6060/debug/pprof/trace?seconds=5
go tool trace trace.out
```

To view all available profiles, open http://localhost:6060/debug/pprof/ in your browser.

> for more check this [doc](https://pkg.go.dev/net/http/pprof)

## Go Profiling With pprof : Hands-on examples

To profile a program, we can use the `runtime/pprof` package that exposes the necessary API to start and stop profiling.

### Example-1 : CPU profiling

In this section, we will profile a program that sums integers. The program consists only of a main package. To perform profiling follow below steps:

- Enable profiling in the Go program by calling the `pprof.StartCPUProfile` (for ex.) function at the beginning of the program. This will start collecting data about the execution of the program (CPU profiling).
- Now we can stop profiling and generate the profile data when we wanna do that in our code by calling the `pprof.StopCPUProfile` function. This will stop profiling and save the data to a file (that file or io writer we'll have to provide, we're gonna see in the example code below)
- Now can run the program as normal way by building a binary of the program (ex: `go build -o main .`), this will create/save the the result/data in the given file.
- Let's use `pprof` tool now for visualizing and analyzing the generated/created data, have to pass the name of the binary and profile data file as arguments. For example:
  - We can run by this command `go tool pprof <binary_name> <file_name_of_pb>`. This command will launch an interactive mode. We'll have to type respective `pprof` commands (`top`, `web` etc.) to display the statistics

Example Go program:

```
package main

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
)

func main() {
	// create the file where we wanna save the profiling result to be saved
	f, err := os.Create("./result.pb")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// start profiling
	err = pprof.StartCPUProfile(f)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

  // stop profiling just before returning from this function, that's why used `defer` keyword
	defer pprof.StopCPUProfile()


	// our actual code for summing some integers

	sum := 0;
	for i := 0; i <= 787766777; i++ {
		sum += i
	}

	fmt.Println("Sum:", sum)
}

```

This code will start profiling when the program starts, and it will stop profiling and save the data to a file when the program exits.

That’s it! we just ran a simple Go program and profiled it using `pprof`.

### Example-2: CPU Profiling 2

Let's see another example for CPU profiling. let's use the following sample code:

```
package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func main() {
	// create the file where we wanna save the data
	f, err := os.Create("./cpu.pprof")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Start the profiling
	pprof.StartCPUProfile(f)

	// stop the profiling at the end of this function
	defer pprof.StopCPUProfile()

	sum := sumNumbers(1e9)

	fmt.Println("Sum: ", sum)
}

func sumNumbers(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}

```

Now, let's run the program by `go run .`, a file named `cpu.pprof` will be generated containing the CPU profile data. after that we can analyze this data using the `go tool pprof` command:

```
go tool pprof cpu.pprof
```

This will open an interactive prompt where we can use `pprof` commands to analyze the profile data. For example, we can use the `top` command to see the top CPU consumers:

```
(pprof) top
```

We can also generate a graphical representation of the CPU usage using the `web` command:

```
(pprof) web
```

This will generate an `SVG` file that we can open in our browser to visualize the CPU usage of our program.

### Example-3: Profiling Memory Usage

> basically by memory we're referencing the heap memory

Profiling memory usage is similar to profiling CPU usage. We will need to:

- Take a memory snapshot using `pprof.WriteHeapProfile`
- Analyze the generated profile data

```
package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func main() {
	// create the file where we wanna save the data
	f, err := os.Create("./heap.prof")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer f.Close()

	if err := pprof.WriteHeapProfile(f); err != nil {
		fmt.Println("could not write memory profile: ", err)
		return
	}
}

```

Now, let's run the code by `go run .`, a file named `heap.prof` will be generated containing the memory profile data. We can analyze this data using the `go tool pprof` command:

```
go tool pprof heap.prof
```

Again, we can use commands like `top` and `web` to analyze the memory usage of your program.

### Example-4: Heap/others profile with `net/http/pprof`

let's use this sample code:

```
package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"

	"github.com/gorilla/mux"
)

func main() {
    // we need a webserver to get the pprof webserver
	router := mux.NewRouter()
	router.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	router.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	router.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	router.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	router.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
	router.Handle("/debug/pprof/{cmd}", http.HandlerFunc(pprof.Index)) // special handling for Gorilla mux

	err := http.ListenAndServe("localhost:6060", router)
	fmt.Println("pprof server listen failed: %v", err)
}
```

Let's run this code by `go tool pprof http://localhost:6060/debug/pprof/heap`. It'll open an interactive mode where we run the `pprof` commands, like `top` etc.

We can also do the same thing outside interactive mode with `go tool pprof -top http://localhost:6060/debug/pprof/heap`

And, to generate a PNG profile use `go tool pprof -png http://localhost:6060/debug/pprof/heap > out.png`

## Easily Profiling with `profile` package

- A simple packege for easily profiling in Go
- This is the package [repo](https://github.com/pkg/profile)
- Enabling profiling in your application is as simple as one line at the top of your main function:

  ```
  import "github.com/pkg/profile"

  func main() {
      defer profile.Start().Stop()
      ...
  }
  ```

- Options it has:

  ```
  // CPUProfile enables cpu profiling. Note: Default is CPU
  defer profile.Start(profile.CPUProfile).Stop()

  // GoroutineProfile enables goroutine profiling.
  // It returns all Goroutines alive when defer occurs.
  defer profile.Start(profile.GoroutineProfile).Stop()

  // BlockProfile enables block (contention) profiling.
  defer profile.Start(profile.BlockProfile).Stop()

  // ThreadcreationProfile enables thread creation profiling.
  defer profile.Start(profile.ThreadcreationProfile).Stop()

  // MemProfileHeap changes which type of memory profiling to
  // profile the heap.
  defer profile.Start(profile.MemProfileHeap).Stop()

  // MemProfileAllocs changes which type of memory to profile
  // allocations.
  defer profile.Start(profile.MemProfileAllocs).Stop()

  // MutexProfile enables mutex profiling.
  defer profile.Start(profile.MutexProfile).Stop()
  ```

  for more check the repo's doc

## Resources

- https://pkg.go.dev/runtime/pprof
- https://github.com/google/pprof
- https://www.youtube.com/watch?v=nok0aYiGiYA
- https://www.practical-go-lessons.com/chap-36-program-profiling
- https://www.freecodecamp.org/news/how-i-investigated-memory-leaks-in-go-using-pprof-on-a-large-codebase-4bec4325e192/
- https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/
- https://pkg.go.dev/net/http/pprof
