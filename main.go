package main

import (
	"fmt"
	"os"
)

func main() {
	workingdir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Println(workingdir)

}
