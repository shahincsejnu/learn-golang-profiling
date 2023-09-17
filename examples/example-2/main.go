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
