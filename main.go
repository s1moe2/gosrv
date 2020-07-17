package main

import (
	"fmt"
	"github.com/s1moe2/gosrv/server"
	"os"
)

func main() {
	if err := server.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
