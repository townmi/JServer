package utils

import (
	"fmt"
	"os"
)

func printOs() {

	environ := os.Environ()
	for i := range environ {
		fmt.Println(environ[i])
	}
	logname := os.Getenv("LOGNAME")
	fmt.Printf("logname is %s\n", logname)
}
