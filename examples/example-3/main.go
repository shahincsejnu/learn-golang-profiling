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
