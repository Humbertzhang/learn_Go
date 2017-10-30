package main

import "fmt"

func main() {
	s := ""
	fmt.Print("Input a string:")
	fmt.Scanf("%s", &s)
	s2 := basename(s)
	fmt.Printf("Basename: %s\n", s2)
}

func basename(s string) string {
	start := 0
	end := len(s)
	for i := range s {
		if s[i] == '/' {
			start = i + 1
		} else if s[i] == '.' {
			end = i
		}
	}
	s2 := s[start:end]
	return s2
}
