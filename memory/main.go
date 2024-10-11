package main

import (
	"fmt"
	"os"
)

func main() {
	files := make([]*os.File, 0)
	for i := 0; ; i++ {
		file, err := os.OpenFile("test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Error at file %d: %v\n", i, err)
			break
		} else {
			_, _ = file.Write([]byte("Hello, World!"))
			files = append(files, file)
		}
	}
}
