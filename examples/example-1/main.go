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

	defer pprof.StopCPUProfile()


	// our actual code

	sum := 0;
	for i := 0; i <= 787766777; i++ {
		sum += i
	}

	fmt.Println("Sum:", sum)
}
