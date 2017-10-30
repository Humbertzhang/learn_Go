package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	arg := os.Args[1]
	if len(os.Args) >= 3 {
		fmt.Fprintf(os.Stderr, "Too many urls:%s", arg)
		os.Exit(1)
	}
	if !strings.HasPrefix(arg, "http://") {
		arg = "http://" + arg
	}

	start := time.Now()
	http.Get(arg)
	secs1 := time.Since(start).Seconds()
	fmt.Println("time1", secs1)

	start2 := time.Now()
	http.Get(arg)
	secs2 := time.Since(start2).Seconds()
	fmt.Println("time2", secs2)
}
