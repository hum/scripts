package main

import (
	"fmt"
	"github.com/hum/scripts-go/portScanner"
)

func main() {
	goroutines := 100
	results := portScanner.Scan(1024, "scanme.nmap.org", goroutines)

	fmt.Printf("Open ports:\n%+v\n", results)
}
