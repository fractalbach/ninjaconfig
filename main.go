package main

import (
	"fmt"
	"github.com/fractalbach/ninjaconfig/somepkg"
	"log"
	"os"
)

var output = `
==================================================
                    NINJACONFIG                 
__________________________________________________

%s
`

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(output, wd)

	somepkg.F1()
}
